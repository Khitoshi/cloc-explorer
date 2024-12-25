package clocexplorer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GitHubFile struct {
	Path string `json:"path"`
}

type GitHubTreeResponse struct {
	Tree []GitHubFile `json:"tree"`
}

var httpClient = &http.Client{}

func FetchFilesFromGitHub(ri RepositoryInfo) (paths []string, err error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1", ri.UserName, ri.RepositoryName, ri.BranchName)

	res, err := httpClient.Get(url)
	if err != nil {
		log.Println(url, "\n", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var treeResponse GitHubTreeResponse
	err = json.Unmarshal(body, &treeResponse)
	if err != nil {
		log.Println("Failed to parse JSON:", err)
		return nil, err
	}

	for _, file := range treeResponse.Tree {
		paths = append(paths, file.Path)
	}

	return paths, nil
}

func fetchCodeFromGitHub(userName string, repositoryName string, fileName string, branchName string) (bodystr string, err error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", userName, repositoryName, branchName, fileName)

	res, err := http.Get(url)
	if err != nil {
		log.Println(url, "\n", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(body), nil
}
