/*
 * PWRP - Personal Work Recorder Processor
 * Copyright (C) 2019  Cezar Mathe <cezarmathe@gmail.com> [https://cezarmathe.com]
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package gitops

import (
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

/*Clone clones a repository in a given path with specifically-selected options.*/
func Clone(repositoryURL, storagePath string) (*git.Repository, error) {
	log.DebugFunctionCalled(repositoryURL, storagePath)

	log.Trace("extract repository name from ", repositoryURL)
	urlEndpoints := strings.Split(repositoryURL, "/")
	repositoryName := urlEndpoints[len(urlEndpoints)-2] + "/" + urlEndpoints[len(urlEndpoints)-1]
	if strings.HasSuffix(repositoryName, ".git") {
		repositoryName = strings.TrimSuffix(repositoryName, ".git")
	}
	storagePath += "/" + repositoryName

	log.Trace("repository storage path: ", storagePath)

	/*try to open the repository first*/
	/*if it exists, load it and pull any remote changes*/
	/*otherwise, clone the repository*/
	log.Trace("trying to open the repository located at ", storagePath)
	gitRepo, err := git.PlainOpen(storagePath)

	if err != nil {
		log.Trace("repository does not exist, cloning...")
		gitRepo, err = git.PlainClone(storagePath, false, &git.CloneOptions{
			URL:   repositoryURL,
			Depth: 1,
		})
	} else {
		log.Trace("repository exists, pulling remote changes")
		workTree, err := gitRepo.Worktree()
		if err != nil {
			log.WarnErr(err, "encountered an error when extracting the repository work tree")
			return nil, err
		}

		err = workTree.Pull(&git.PullOptions{RemoteName: "origin"})
	}
	if gitRepo != nil {
		log.DebugFunctionReturned(*gitRepo, err)
	} else {
		log.DebugFunctionReturned(nil, err)
	}
	return gitRepo, err
}
