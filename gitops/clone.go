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

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
)

/*Clone clones a repository in a given path with specifically-selected options.*/
func Clone(repositoryURL, storagePath string) (*git.Repository, error) {
	log.Debug("clone(): ", "called")

	urlEndpoints := strings.Split(repositoryURL, "/")
	repositoryName := urlEndpoints[len(urlEndpoints)-1]
	if strings.HasSuffix(repositoryName, ".git") {
		strings.TrimSuffix(repositoryName, ".git")
	}
	storagePath += "/" + repositoryName
	return git.PlainClone(storagePath, false, &git.CloneOptions{
		URL:   repositoryURL,
		Depth: 1,
	})
}
