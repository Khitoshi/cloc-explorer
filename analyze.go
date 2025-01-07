package clocexplorer

import (
	"log"
	"strings"
	"sync"
)

func PopulateFilePaths(filePath []string, languages *DefinedLanguages) *DefinedLanguages {
	for _, fp := range filePath {
		fileType, ok := GetFileType(fp)
		if !ok {
			continue
		}
		languages.Langs[fileType].Files = append(languages.Langs[fileType].Files, fp)
	}
	return languages
}

type ClockFile struct {
	Lang     string
	Code     uint64
	Comments uint64
	Blanks   uint64
}

func AnalyzeLanguages(languages *DefinedLanguages, ri RepositoryInfo) *DefinedLanguages {
	cfChan := make(chan ClockFile, len(languages.Langs))
	var wg sync.WaitGroup
	wg.Add(len(languages.Langs))
	for lang, langData := range languages.Langs {
		go func(lang string, langData *Language) {
			defer wg.Done()
			analyzeFilePaths(cfChan, lang, langData.Files, langData.lineComments, ri)
		}(lang, langData)
	}

	wg.Wait()
	close(cfChan)
	for cf := range cfChan {
		languages.Langs[cf.Lang].Comments = cf.Comments
		languages.Langs[cf.Lang].Blanks = cf.Blanks
		languages.Langs[cf.Lang].Code = cf.Code
	}

	return languages
}

func analyzeFilePaths(cfChan chan ClockFile, lang string, filepaths []string, lineComment []string, ri RepositoryInfo) {
	cf := ClockFile{Lang: lang}
	for _, filepath := range filepaths {
		code, err := fetchCodeFromGitHub(ri.UserName, ri.RepositoryName, filepath, ri.BranchName)
		if err != nil {
			log.Println(err)
		}
		analyzeCodeData := analyzeCodeContent(code, lineComment)
		cf.Blanks += analyzeCodeData.Blanks
		cf.Comments += analyzeCodeData.Comments
		cf.Code += analyzeCodeData.Code
	}

	cfChan <- cf
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
