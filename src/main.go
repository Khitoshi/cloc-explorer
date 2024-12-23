package main

import (
	"log"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
)

func main() {
	ri, err := clocexplorer.NewRepositoryInfo("Khitoshi/GameGraphicsLibrary", "main")
	if err != nil {
		log.Println(err)
		return
	}

	paths, err := clocexplorer.FetchFilesFromGitHub(ri)
	if err != nil {
		log.Println(err)
		return
	}

	languages := clocexplorer.NewDefinedLanguages()

	language := clocexplorer.NewFileData(paths, *languages)
	language = clocexplorer.AnalyzeFile(language, ri)

	for lang, langData := range language.Langs {
		log.Printf("言語:%s", lang)
		log.Printf("Comments:%d", langData.Comments)
		log.Printf("Code:%d", langData.Code)
		log.Printf("Blanks:%d", langData.Blanks)
		log.Printf("\n\n")
	}
}
