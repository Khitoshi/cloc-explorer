package main

import (
	"log"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
)

func main() {
	userName, repositoryName, err := clocexplorer.ParseRepository("Khitoshi/rust-directx12")
	if err != nil {
		log.Println(err)
		return
	}
	ri := clocexplorer.NewRepositoryInfo(userName, repositoryName, "main")

	body, err := clocexplorer.FetchFilesFromGitHub(ri)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))

	//fb := []clocexplorer.FileData{}
	//clocexplorer.AnalyzeFile(fb, ri)
}
