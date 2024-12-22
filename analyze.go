package clocexplorer

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var httpClient = &http.Client{}

type fileType struct {
	Lang    string
	Format  string
	Comment string
}

func NewFileType() []fileType {
	return []fileType{
		{"Go", ".go", "//"},
	}
}

type FileData struct {
	FileName   []string
	FormatType fileType
}

type ClockFile struct {
	Code     uint64
	Comments uint64
	Blanks   uint64
	FileType fileType
}

type OutputData struct {
	Total    ClockFile
	Language map[string]ClockFile
}

func AnalyzeFile(fd []FileData, ri RepositoryInfo) OutputData {

	od := OutputData{}
	for _, f := range fd {
		cf := analyzeCode(f, ri)
		od.Language[cf.FileType.Lang] = cf
	}

	return od
}

func analyzeCode(fd FileData, ri RepositoryInfo) ClockFile {
	cf := ClockFile{}

	for _, f := range fd.FileName {
		code, err := fetchCodeFromGitHub(ri.UserName, ri.RepositoryName, f, ri.BranchName)
		if err != nil {
			log.Println(err)
			return cf
		}
		log.Println(code)
		code = strings.TrimSpace(code)
		log.Println(code)

		/*
			if strings.HasPrefix(code, fd.FormatType.Comment) {
				cf.Comments += 0
				continue
			}

			cf.Code += 0

			cf.Blanks += 0
		*/
	}

	return cf
}

func FetchFilesFromGitHub(ri RepositoryInfo) (bodystr []byte, err error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents?ref=%s", ri.UserName, ri.RepositoryName, ri.BranchName)
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

	//return string(body), nil
	return body, nil
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
