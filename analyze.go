package clocexplorer

import (
	"log"
	"net/http"
	"path/filepath"
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

func NewFileData(filePath []string) FileData {
	fd := FileData{}
	fd.FormatType = NewFileType()[0]

	for _, ft := range filePath {
		if getFileType(ft) != fd.FormatType.Format {
			continue
		}
		fd.FileName = append(fd.FileName, ft)
	}
	return fd
}

func getFileType(path string) string {
	ext := filepath.Ext(path)
	return ext
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

func AnalyzeFile(fd FileData, ri RepositoryInfo) OutputData {
	od := OutputData{}
	//for _, f := range fd {
	//	cf := analyzeCode(f, ri)
	//	od.Language[cf.FileType.Lang] = cf
	//}
	cf := analyzeCode(fd, ri)
	od.Language = make(map[string]ClockFile)
	od.Language[cf.FileType.Lang] = cf

	return od
}

func analyzeCode(fd FileData, ri RepositoryInfo) ClockFile {
	cf := ClockFile{}
	cf.FileType = fd.FormatType
	for _, f := range fd.FileName {
		code, err := fetchCodeFromGitHub(ri.UserName, ri.RepositoryName, f, ri.BranchName)
		if err != nil {
			log.Println(err)
			return cf
		}
		a := analyzeCodeContent(code)
		cf.Blanks += a.Blanks
		cf.Comments += a.Comments
		cf.Code += a.Code
	}

	return cf
}

func analyzeCodeContent(code string) ClockFile {
	cf := ClockFile{}
	for _, line := range strings.Split(code, "\n") {
		line = strings.TrimSpace(line)

		//空行
		if len(strings.TrimSpace(line)) == 0 {
			cf.Blanks++
			continue
		}

		//コメント
		if strings.HasPrefix(line, "//") {
			cf.Comments++
			continue
		}

		cf.Code++
	}
	return cf
}
