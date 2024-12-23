package main

import (
	"log"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
)

func main() {
	userName, repositoryName, err := clocexplorer.ParseRepository("Khitoshi/gocloc")
	if err != nil {
		log.Println(err)
		return
	}
	ri := clocexplorer.NewRepositoryInfo(userName, repositoryName, "master")

	paths, err := clocexplorer.FetchFilesFromGitHub(ri)
	if err != nil {
		log.Println(err)
		return
	}

	//for _, line := range strings.Split(prettyJSON.String(), "\n") {
	//	log.Println(line)
	//}

	fb := clocexplorer.NewFileData(paths)
	log.Println(fb)
	od := clocexplorer.AnalyzeFile(fb, ri)

	log.Printf("言語:%s", od.Language["Go"].FileType.Lang)
	log.Printf("Comments:%d", od.Language["Go"].Comments)
	log.Printf("Code:%d", od.Language["Go"].Code)
	log.Printf("Blanks:%d", od.Language["Go"].Blanks)

}
