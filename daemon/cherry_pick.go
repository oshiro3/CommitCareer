package daemon

import (
	"github.com/libgit2/git2go"
	"log"
)

func CherryPick(repo *git.Repository, commitId string, baseBranch string) (*git.Oid, error) {
	oid, err := git.NewOid(commitId)
	if err != nil {
		log.Println("Fail to get object Id")
		return nil, err
	}
	commit, _ := repo.LookupCommit(oid)
	if err != nil {
		log.Println("Fail to get Commit")
		return nil, err
	}
	CheckoutBranch(repo, baseBranch)

	head, err := repo.Head()
	if err != nil {
		log.Println("Fail to get Head")
		return nil, err
	}
	base, err := repo.LookupCommit(head.Target())
	if err != nil {
		log.Println("Fail to get Head Commit")
		return nil, err
	}

	cherrypickOption, _ := git.DefaultCherrypickOptions()
	err = repo.Cherrypick(commit, cherrypickOption)
	if err != nil {
		log.Printf("Fail to Cherry Pick")
		return nil, err
	}

	index, _ := repo.Index()
	treeId, err := index.WriteTree()
	if err != nil {
		log.Printf("Fail to Cherry Pick")
		return nil, err
	}

	tree, err := repo.LookupTree(treeId)
	if err != nil {
		log.Printf("Fail to lookup Tree")
		return nil, err
	}

	hc, _ := repo.LookupCommit(head.Target())
	nid, err := repo.CreateCommit("HEAD", commit.Author(), commit.Committer(), commit.Message(), tree, hc, base)
	if err != nil {
		log.Printf("Fail to create Commit")
		return nil, err
	}

	repo.StateCleanup()
	head, _ = repo.Head()
	hc, _ = repo.LookupCommit(head.Target())
	return nid, nil
}
