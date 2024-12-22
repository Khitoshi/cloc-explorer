package clocexplorer

type RepositoryInfo struct {
	UserName       string
	RepositoryName string
	BranchName     string
}

func NewRepositoryInfo(userName string, repositoryName string, branchName string) RepositoryInfo {
	return RepositoryInfo{
		UserName:       userName,
		RepositoryName: repositoryName,
		BranchName:     branchName,
	}
}
