package bb

import (
	"fmt"
)

type Repo struct {
	Project  string
	RepoName string
	Env      string
	Url      string
}

func NewRepo(env string, prj string, name string, url string) (repo Repo) {
	repo = Repo{Project: prj, RepoName: name, Env: env, Url: url}
	return repo
}

func (r Repo) GetCloneUrl() string {
	gitclone := fmt.Sprintf("ssh://git@%s/%s/%s.git", r.Url, r.Project, r.RepoName)
	return gitclone
}
