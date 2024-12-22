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
	log.Println(paths)

	//for _, line := range strings.Split(prettyJSON.String(), "\n") {
	//	log.Println(line)
	//}

	//fb := []clocexplorer.FileData{}
	//clocexplorer.AnalyzeFile(fb, ri)
}
