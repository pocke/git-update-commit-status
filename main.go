package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/github/hub/github"
	api "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// git update-commit-status STATUS [revision]
func main() {
	if err := Main(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main(args []string) error {
	status, rev, err := ParseArgs(args)
	if err != nil {
		return err
	}

	rev, err = ParseRev(rev)
	if err != nil {
		return err
	}

	prj, err := Project()
	if err != nil {
		return err
	}

	c, err := APIClient()
	if err != nil {
		return err
	}
	st := &api.RepoStatus{State: &status}
	_, _, err = c.Repositories.CreateStatus(context.Background(), prj.Owner, prj.Name, rev, st)
	return err
}

func ParseArgs(args []string) (string, string, error) {
	if len(args) < 2 {
		return "", "", errors.New("Too few arguments. Usage: git update-commit-status STATUS [revision]")
	}
	if len(args) == 2 {
		return args[1], "@", nil
	}
	return args[1], args[2], nil
}

func ParseRev(rev string) (string, error) {
	rev2, err := exec.Command("git", "rev-parse", rev).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(rev2), "\n"), nil
}

// RepoName returns owner/repo.
func Project() (*github.Project, error) {
	repo, err := github.LocalRepo()
	if err != nil {
		return nil, err
	}
	prj, err := repo.MainProject()
	if err != nil {
		return nil, err
	}

	return prj, nil
}

func APIClient() (*api.Client, error) {
	c := github.CurrentConfig()
	host, err := c.DefaultHost()
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: host.AccessToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	return api.NewClient(tc), nil
}
