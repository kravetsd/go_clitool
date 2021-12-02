package flgs

import (
	"flag"
	"fmt"
	"github.com/kravetsd/go-clitool/bb"
)

func GetClone() (link, path, b string) {

	reponame := flag.String("reponame", "", "specify a repository name")
	project := flag.String("project", "", "specify a project")
	branch := flag.String("env", "", "specify a branch a.k.a environment ")
	path = *flag.String("path", ".", "specify a path where to cone the repo ")
	url := flag.String("url", ".", "specify a bb url ")

	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			fmt.Printf("flag %v must be set!\n", f.Name)
		}
	})
	b, p, r, u := *branch, *project, *reponame, *url
	cloneRepo := bb.NewRepo(b, p, r, u)
	return cloneRepo.GetCloneUrl(), path, b
}
