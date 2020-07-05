package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"log"

	pb "commit_career/pb/proto"
)

var host string = "127.0.0.1"

func FetchBranch() []string {
	conn, err := grpc.Dial(fmt.Sprintf("%s:50051", host), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	c := pb.NewCommitCareerClient(conn)
	message := &pb.FetchBranchRequest{Limit: 10}
	res, err := c.FetchBranches(context.TODO(), message)
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return res.Name
}

func FetchCommits(branch string) []*pb.FetchCommitsResponse_Commit {
	conn, err := grpc.Dial(fmt.Sprintf("%s:50051", host), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	c := pb.NewCommitCareerClient(conn)
	message := &pb.FetchCommitsRequest{Branch: branch}
	res, err := c.FetchCommits(context.TODO(), message)
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return res.Commits
}

func CherryPick(base string, sha []string) (string, int32) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:50051", host), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	c := pb.NewCommitCareerClient(conn)
	message := &pb.CherryPickRequest{Base: base, Sha: sha}
	res, err := c.CherryPick(context.TODO(), message)
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return res.Message, res.Status
}
