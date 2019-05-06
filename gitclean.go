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
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
)

func main() {
	var (
		lineEnding string
	)

	ignoreBranches := make(map[string]struct{})

	ignoreEnv := os.Getenv("GITCLEAN_IGNORE")

	switch {
	case ignoreEnv != "":
		fmt.Printf("GITCLEAN_IGNORE set to '%s'\n", ignoreEnv)
		ib := strings.Split(ignoreEnv, ",")
		for _, i := range ib {
			ignoreBranches[i] = struct{}{}
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

	reader := bufio.NewReader(os.Stdin)
	branchesToDelete := []plumbing.ReferenceName{}
	iter.ForEach(func(pr *plumbing.Reference) error {
		_, ok := ignoreBranches[pr.Name().Short()]
		if pr.Name() != h.Name() && !ok {
			fmt.Printf("Remove branch %s[y,N,q]? ", pr.Name().Short())
			s, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("error reading input:", err)
			}
			s = strings.Replace(s, lineEnding, "", -1)

			switch s {
			case "y":
				branchesToDelete = append(branchesToDelete, pr.Name())
			case "q":
				branchesToDelete = nil
				return storer.ErrStop
			}
		}
		return nil
	})

	if len(branchesToDelete) > 0 {
		fmt.Println("The following branches will be deleted:")
		for _, r := range branchesToDelete {
			fmt.Printf("\t%v\n", r.Short())
		}
		fmt.Print("Continue [y,N]? ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
		}
		s = strings.Replace(s, lineEnding, "", -1)
		if s == "y" {
			for _, r := range branchesToDelete {
				fmt.Printf("\tremoving %v\n", r.Short())
				if err := repo.Storer.RemoveReference(r); err != nil {
					fmt.Println("error deleting branch", err)
				}
			}
		}
	}
	fmt.Println("Exiting")
}
