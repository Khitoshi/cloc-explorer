package main

import (
	"fmt"
	"log"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
	flags "github.com/jessevdk/go-flags"
)

type CmdOptions struct {
	MatchRepository string `long:"match-repository" description:"Match GitHubRepository name"`
	MatchBranch     string `long:"match-branch" description:"Match Branch name"`
}

func NewCmdOptions() *CmdOptions {
	return &CmdOptions{
		MatchRepository: "",
		MatchBranch:     "main",
	}
}

func WriteResult(languages clocexplorer.DefinedLanguages) {
	const commonHeader string = "name                           files          blank        comment           code"
	const defaultOutputSeparator string = "-------------------------------------------------------------------------"
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

	fmt.Printf("\n")
	fmt.Printf("%-27v %6v %14v %14v %14v\n", "Total", totalFiles, totalBlanks, totalComments, totalCode)
}

func main() {
	var opts CmdOptions

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "cloc-explorer"
	parser.Usage = "[OPTIONS] PATH[...]"

	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	ri, err := clocexplorer.NewRepositoryInfo(opts.MatchRepository, opts.MatchBranch)
	if err != nil {
		log.Println(err)
		return
	}

	filePaths, err := clocexplorer.FetchFilesFromGitHub(ri)
	if err != nil {
		log.Println(err)
		return
	}

	languages := clocexplorer.NewDefinedLanguages()

	languages = clocexplorer.PopulateFilePaths(filePaths, languages)
	languages = clocexplorer.AnalyzeLanguages(languages, ri)

	WriteResult(*languages)
}
