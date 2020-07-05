package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
)

func Push() {
	branches := FetchBranch()

	base_prompt := promptui.Select{
		Label: "Select Base",
		Items: branches,
	}

	_, base, err := base_prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", base)

	branches = remove(branches, base)

	head_prompt := promptui.Select{
		Label: "Select Branch Containing target sha",
		Items: branches,
	}

	_, head, err := head_prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", head)

	// fmt.Printf("Commits: %+v\n", FetchCommits(head))
	res := FetchCommits(head)
	var commits []string
	for _, c := range res {
		commits = append(commits, fmt.Sprintf("%s: %s", c.Sha, strings.Trim(c.Title, "\n")))
	}
	index := 1
	var sha []string
	for index >= 0 {
		commitsPrompt := promptui.SelectWithAdd{
			Label:    "Select sha to CherryPick",
			Items:    commits,
			AddLabel: "Finish Select",
		}
		var input string
		index, input, _ = commitsPrompt.Run()
		if index >= 0 {
			i := strings.Index(input, ":")
			sha = append(sha, input[:i])
			commits = remove(commits, input)
		}
	}
	message, status := CherryPick(base, sha)
	if status != 200 {
		log.Printf("Fail CherryPick: %s", message)
	}
}

func remove(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}
