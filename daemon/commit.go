package daemon

import (
	"github.com/libgit2/git2go"
	"log"
)

func GetCommits(repo *git.Repository, remoteBranch string) []*git.Commit {
	branch, err := repo.LookupBranch(remoteBranch, git.BranchRemote)
	if err != nil {
		log.Fatalf("ERR: %v\n", err)
	}
	head := branch.Reference
	c, _ := repo.LookupCommit(head.Target())
	var commits []*git.Commit
	commit := c
	commits = append(commits, c)
	for commit.ParentCount() > 0 {
		commit = commit.Parent(0)
		commits = append(commits, commit)
	}
	return commits
}
