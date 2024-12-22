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

func NewRepositoryInfo(userName string, repositoryName string, branchName string) RepositoryInfo {
	return RepositoryInfo{
		UserName:       userName,
		RepositoryName: repositoryName,
		BranchName:     branchName,
	}
}

func ParseRepository(repository string) (userName string, repositoryName string, err error) {
	f := func(c rune) bool {
		return c == '/'
	}
	s := strings.FieldsFunc(repository, f)
	if len(s) != 2 {
		return "", "", fmt.Errorf("Invalid repository name")
	}
	return s[0], s[1], nil
}
