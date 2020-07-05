package daemon

import (
	"github.com/libgit2/git2go"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Clean(path string) {
	gitFile, _ := filepath.Split(filepath.Dir(path))
	log.Printf("cleaning: %v", gitFile)
	os.RemoveAll(gitFile)
}

func Clone(repoPath string) *git.Repository {
	clonePath := "/tmp/commit_career_work"
	err := exec.Command("/usr/bin/git", "clone", repoPath, clonePath).Run()
	if err != nil {
		log.Fatalf("clone err: %+v\n", err)
	}
	repo, err := git.OpenRepository(clonePath)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	return repo
}
