package daemon

import (
	"github.com/libgit2/git2go"
	"log"
)

func BranchList(repo *git.Repository) []*git.Branch {
	itr, err := repo.NewBranchIterator(git.BranchRemote)
	if err != nil {
		log.Println(err)
	}
	defer itr.Free()

	var arr []*git.Branch
	var f = func(b *git.Branch, _ git.BranchType) error {
		// c, _ := b.Name()
		// log.Printf("branches: %+v\n", c)
		arr = append(arr, b)
		return nil
	}
	itr.ForEach(f)
	return arr
}
