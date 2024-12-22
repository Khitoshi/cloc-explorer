package clocexplorer

import (
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
