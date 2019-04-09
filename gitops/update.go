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
	"reflect"

	"gopkg.in/src-d/go-git.v4"
)

func Update(gitRepo *git.Repository) bool {
	log.DebugFunctionCalled(*gitRepo)

	workTree, err := gitRepo.Worktree()
	if err != nil {
		log.ErrorErr(err, "encountered an error when extracting the repository work tree")

		log.DebugFunctionReturned(false)
		return false
	}

	log.Trace("pulling remote changes")
	err = workTree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(git.NoErrAlreadyUpToDate) {
			log.Trace("repository is already up to date")
		} else {
			log.ErrorErr(err, "encountered an error when pulling the remote changes")

			log.DebugFunctionReturned(false)
			return false
		}
	}

	log.DebugFunctionReturned(true)
	return true
}
