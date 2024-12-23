package clocexplorer

import (
	"log"
	"net/http"
	"strings"
)

var httpClient = &http.Client{}

func NewFileData(filePath []string, languages DefinedLanguages) DefinedLanguages {
	for _, fp := range filePath {
		fileType, ok := GetFileType(fp)
		if !ok {
			continue
		}
		//FileNames[ext] = append(FileNames[ext], fp)
		languages.Langs[fileType].Files = append(languages.Langs[fileType].Files, fp)
	}
	return languages
}

type ClockFile struct {
	Code     uint64
	Comments uint64
	Blanks   uint64
}

func AnalyzeFile(languages DefinedLanguages, ri RepositoryInfo) DefinedLanguages {
	for lang, langData := range languages.Langs {
		cf := analyzeCode(langData.Files, langData.lineComments, ri)
		languages.Langs[lang].Code = cf.Code
		languages.Langs[lang].Comments = cf.Comments
		languages.Langs[lang].Blanks = cf.Blanks
	}

	return languages
}

func analyzeCode(filepaths []string, lineComment []string, ri RepositoryInfo) ClockFile {
	cf := ClockFile{}
	for _, filepath := range filepaths {
		code, err := fetchCodeFromGitHub(ri.UserName, ri.RepositoryName, filepath, ri.BranchName)
		if err != nil {
			log.Println(err)
			return cf
		}
		analyzeCodeData := analyzeCodeContent(code, lineComment)
		cf.Blanks += analyzeCodeData.Blanks
		cf.Comments += analyzeCodeData.Comments
		cf.Code += analyzeCodeData.Code
	}

	return cf
}

func analyzeCodeContent(code string, lineComment []string) ClockFile {
	cf := ClockFile{}
	for _, line := range strings.Split(code, "\n") {
		line = strings.TrimSpace(line)

		//空行
		if len(strings.TrimSpace(line)) == 0 {
			cf.Blanks++
			continue
		}

		//コメント
		for _, lc := range lineComment {
			if strings.HasPrefix(line, lc) {
				cf.Comments++
				break
			}
		}

		cf.Code++
	}
	return cf
}
