package test

import (
	"testing"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
	"github.com/stretchr/testify/assert"
)

func Test_GetFileType(t *testing.T) {
	//拡張子が存在する場合
	ext, ok := clocexplorer.GetFileType("hoge.go")
	assert.True(t, ok)
	assert.Equal(t, ext, clocexplorer.Exts["go"])

	//拡張子が存在しない場合
	ext, ok = clocexplorer.GetFileType("hoge")
	assert.False(t, ok)
	assert.Equal(t, ext, clocexplorer.Exts[""])

	//拡張子が複数存在する場合
	ext, ok = clocexplorer.GetFileType("hoge.hoge.go")
	assert.True(t, ok)
	assert.Equal(t, ext, clocexplorer.Exts["go"])

	//登録していない拡張子の場合
	ext, ok = clocexplorer.GetFileType("hoge.txt")
	assert.False(t, ok)
	assert.Equal(t, ext, clocexplorer.Exts["txt"])

	ext, ok = clocexplorer.GetFileType("")
	assert.False(t, ok)
	assert.Equal(t, ext, clocexplorer.Exts[""])
}
