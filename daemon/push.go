package daemon

import (
	"github.com/libgit2/git2go"
	"log"
)

func Push(repo *git.Repository, branch string) error {
	remote, _ := repo.Remotes.Lookup("origin")
	if err := remote.Push([]string{"refs/heads/" + branch}, &git.PushOptions{}); err != nil {
		log.Println("Fail to push branch")
		return err
	}
	return nil
}
