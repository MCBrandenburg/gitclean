package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	var (
		lineEnding string
	)

	ignoreBranches := make(map[string]interface{})

	ignoreEnv := os.Getenv("GITCLEAN_IGNORE")

	switch {
	case ignoreEnv != "":
		fmt.Printf("GITCLEAN_IGNORE set to '%s'\n", ignoreEnv)
		ib := strings.Split(ignoreEnv, ",")
		for _, i := range ib {
			ignoreBranches[i] = nil
		}
	default:
		fmt.Println("GITCLEAN_IGNORE not set")
	}

	switch runtime.GOOS {
	case "windows":
		lineEnding = "\r\n"
	default:
		lineEnding = "\n"
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	iter, err := repo.Branches()
	if err != nil {
		log.Fatal(err)
	}

	h, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	branches := []*plumbing.Reference{}
	iter.ForEach(func(pr *plumbing.Reference) error {
		_, ok := ignoreBranches[pr.Name().Short()]
		fmt.Println("debug", pr.Name().Short(), ok)
		if pr.Name() != h.Name() && !ok {
			branches = append(branches, pr)
		}
		return nil
	})

	if len(branches) == 0 {
		fmt.Printf("No branches other than current branch(%s) exists\n", h.Name().Short())
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)

	for _, r := range branches {
		fmt.Printf("Remove branch %s? ('yes' to remove): ", r.Name().Short())
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
		}
		s = strings.Replace(s, lineEnding, "", -1)

		if s == "yes" {
			fmt.Printf("removing %s\n", r.Name().Short())
			if err := repo.Storer.RemoveReference(r.Name()); err != nil {
				fmt.Println(err)
			}
		}
	}
}
