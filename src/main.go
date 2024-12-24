package main

import (
	"fmt"
	"log"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
)

func WriteResult(languages clocexplorer.DefinedLanguages) {
	// write header
	//o.WriteHeader()
	const commonHeader string = "name                           files          blank        comment           code"
	const defaultOutputSeparator string = "-------------------------------------------------------------------------" +
		"-------------------------------------------------------------------------"
	//"-------------------------------------------------------------------------"
	fmt.Println(commonHeader)
	fmt.Println(defaultOutputSeparator)

	totalFiles := uint64(0)
	totalBlanks := uint64(0)
	totalComments := uint64(0)
	totalCode := uint64(0)

	for _, lang := range languages.Langs {
		fmt.Printf("%-27v %6v %14v %14v %14v\n", lang.Name, len(lang.Files), lang.Blanks, lang.Comments, lang.Code)

		totalFiles += uint64(len(lang.Files))
		totalBlanks += lang.Blanks
		totalComments += lang.Comments
		totalCode += lang.Code
	}

	// write footer
	//o.WriteFooter()
	//fmt.Println('\n')
	fmt.Printf("\n")
	fmt.Printf("%-27v %6v %14v %14v %14v\n", "Total", totalFiles, totalBlanks, totalComments, totalCode)
}

func main() {
	ri, err := clocexplorer.NewRepositoryInfo("Khitoshi/gocloc", "master")
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

	//for lang, langData := range language.Langs {
	//	log.Printf("言語:%s", lang)
	//	log.Printf("Comments:%d", langData.Comments)
	//	log.Printf("Code:%d", langData.Code)
	//	log.Printf("Blanks:%d", langData.Blanks)
	//	log.Printf("\n\n")
	//}
	WriteResult(language)
}
