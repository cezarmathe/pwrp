package gitops

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
)

func Clone(repositoryURL, storagePath string) (*git.Repository, error) {
	log.Debug("clone(): ", "called")

	urlEndpoints := strings.Split(repositoryURL, "/")
	repositoryName := urlEndpoints[len(urlEndpoints)-1]
	if strings.HasSuffix(repositoryName, ".git") {
		strings.TrimSuffix(repositoryName, ".git")
	}
	storagePath += "/" + repositoryName
	return git.PlainClone(storagePath, false, &git.CloneOptions{
		URL: repositoryURL,
	})
}
