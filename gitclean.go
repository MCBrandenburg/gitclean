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

	reader := bufio.NewReader(os.Stdin)
	iter.ForEach(func(pr *plumbing.Reference) error {
		_, ok := ignoreBranches[pr.Name().Short()]
		if pr.Name() != h.Name() && !ok {
			fmt.Printf("Remove branch %s? ('yes' to remove): ", pr.Name().Short())
			s, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("error reading input:", err)
			}
			s = strings.Replace(s, lineEnding, "", -1)

			if s == "yes" {
				fmt.Printf("removing %s\n", pr.Name().Short())
				if err := repo.Storer.RemoveReference(pr.Name()); err != nil {
					fmt.Println(err)
				}
			}
		}
		return nil
	})
}
