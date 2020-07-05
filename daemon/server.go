package daemon

import (
	pb "commit_career/pb/proto"
	"context"
	"fmt"
	// "github.com/libgit2/git2go"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strings"
)

type CommitCareerService struct{}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func (s *CommitCareerService) FetchBranches(ctx context.Context, message *pb.FetchBranchRequest) (*pb.FetchBranchResponse, error) {
	repoPath := "/home/oshiro/works/bares.git"
	repo := Clone(repoPath)
	defer Clean(repo.Path())
	br := BranchList(repo)
	var branches []string
	for _, v := range br {
		n, _ := v.Name()
		branches = append(branches, n)
	}
	return &pb.FetchBranchResponse{
		Name: branches,
	}, nil
}

func (s *CommitCareerService) FetchCommits(ctx context.Context, message *pb.FetchCommitsRequest) (*pb.FetchCommitsResponse, error) {
	repoPath := "/home/oshiro/works/bares.git"
	repo := Clone(repoPath)
	defer Clean(repo.Path())
	remoteBranch := message.Branch
	refs := GetCommits(repo, remoteBranch)
	var commits []*pb.FetchCommitsResponse_Commit
	for _, v := range refs {
		commit := &pb.FetchCommitsResponse_Commit{
			Sha:   v.Id().String(),
			Title: v.Message(),
		}
		commits = append(commits, commit)
	}

	return &pb.FetchCommitsResponse{
		Commits: commits,
	}, nil
}

func (s *CommitCareerService) CherryPick(ctx context.Context, message *pb.CherryPickRequest) (*pb.CherryPickResponse, error) {
	repoPath := "/home/oshiro/works/bares.git"
	repo := Clone(repoPath)
	defer Clean(repo.Path())
	var localBranch string
	if strings.Contains(message.Base, "origin/") {
		localBranch = message.Base[7:]
	} else {
		localBranch = message.Base
	}
	var msg string = "ok"
	var status int32 = 200
	for _, sha := range message.Sha {
		_, err := CherryPick(repo, sha, localBranch)
		if err != nil {
			log.Println("Fail to CherryPick")
			msg = err.Error()
			status = 500
		}
	}

	if status == 200 {
		err := Push(repo, localBranch)
		if err != nil {
			log.Println("Fail to CherryPick")
			msg = err.Error()
			status = 500
		}
	}

	return &pb.CherryPickResponse{
		Message: msg,
		Status:  status,
	}, nil
}

func Run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Println("http server start")
		http.HandleFunc("/hello", handler)
		http.ListenAndServe(":50052", nil)
	}()

	s := grpc.NewServer()
	pb.RegisterCommitCareerServer(s, &CommitCareerService{})

	log.Println("gRPC server start")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
