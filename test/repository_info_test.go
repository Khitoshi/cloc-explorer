package test

import (
	"testing"

	clocexplorer "github.com/Khitoshi/cloc-explorer"
	"github.com/stretchr/testify/assert"
)

func TestNewRepositoryInfo(t *testing.T) {
	tests := []struct {
		repository  string
		branchName  string
		expected    clocexplorer.RepositoryInfo
		expectError bool
	}{
		{
			repository: "user/repo",
			branchName: "main",
			expected: clocexplorer.RepositoryInfo{
				UserName:       "user",
				RepositoryName: "repo",
				BranchName:     "main",
			},
			expectError: false,
		},
		{
			repository:  "invalidrepo",
			branchName:  "main",
			expected:    clocexplorer.RepositoryInfo{},
			expectError: true,
		},
	}

	for _, test := range tests {
		result, err := clocexplorer.NewRepositoryInfo(test.repository, test.branchName)
		if test.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}
