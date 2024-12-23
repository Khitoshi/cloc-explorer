package clocexplorer

import (
	"fmt"
	"log"
	"strings"
)

type RepositoryInfo struct {
	UserName       string
	RepositoryName string
	BranchName     string
}

func NewRepositoryInfo(repository string, branchName string) (RepositoryInfo, error) {
	userName, repositoryName, err := ParseRepository(repository)
	if err != nil {
		log.Println(err)
		return RepositoryInfo{}, err
	}

	return RepositoryInfo{
		UserName:       userName,
		RepositoryName: repositoryName,
		BranchName:     branchName,
	}, nil
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
