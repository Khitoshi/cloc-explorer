package clocexplorer

import (
	"fmt"
	"strings"
)

type RepositoryInfo struct {
	UserName       string
	RepositoryName string
	BranchName     string
}

func NewRepositoryInfo(repository string, branchName string) (RepositoryInfo, error) {
	userName, repositoryName, err := parseRepository(repository)
	if err != nil {
		return RepositoryInfo{}, err
	}

	return RepositoryInfo{
		UserName:       userName,
		RepositoryName: repositoryName,
		BranchName:     branchName,
	}, nil
}

func parseRepository(repository string) (userName string, repositoryName string, err error) {
	f := func(c rune) bool {
		return c == '/'
	}
	s := strings.FieldsFunc(repository, f)
	if len(s) != 2 {
		return "", "", fmt.Errorf("parseRepository: Invalid repository or branch name")
	}
	return s[0], s[1], nil
}
