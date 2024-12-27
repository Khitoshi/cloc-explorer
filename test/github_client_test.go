package test

import (
	"testing"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
	"github.com/stretchr/testify/assert"
)

func TestFetchFilesFromGitHub(t *testing.T) {

	ri, err := clocexplorer.NewRepositoryInfo("octocat/Hello-World", "master")
	assert.NoError(t, err)

	paths, err := clocexplorer.FetchFilesFromGitHub(ri)
	assert.NoError(t, err)
	assert.NotEmpty(t, paths)
	assert.Equal(t, paths[0], "README")
}
