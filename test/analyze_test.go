package test

import (
	"testing"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
	"github.com/stretchr/testify/assert"
)

func Test_PopulateFilePaths(test *testing.T) {
	languages := clocexplorer.NewDefinedLanguages()
	filePaths := []string{"main.go", "index.js"} // テスト用ファイルパス
	updatedLang := clocexplorer.PopulateFilePaths(filePaths, languages)

	assert.NotNil(test, updatedLang)
	assert.Equal(test, 1, len(updatedLang.Langs["Go"].Files))
	assert.Equal(test, "main.go", updatedLang.Langs["Go"].Files[0])
}
