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
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

/*Checkout switches the work tree of the given repository to the given branch name*/
func Checkout(repository *git.Repository, branchName string) bool {
	log.DebugFunctionCalled(*repository, branchName)

	success := new(bool)
	*success = false

	defer log.DebugFunctionReturned(*success)

	log.Trace("retrieving the work tree")
	workTree, err := repository.Worktree()
	if err != nil {
		log.ErrorErr(err, "could not extract the work tree")
		return *success
	}

	head, err := repository.Head()
	if head.Name() == plumbing.NewBranchReferenceName(branchName) {
		log.Trace("already on metadata branch")
		*success = true
		return *success
	}

	log.Trace("checking if the branch exists")
	_, err = repository.Branch(branchName)
	if err != nil {
		/*if error is branch does not exist, create it*/
		if err == git.ErrBranchNotFound {
			log.WarnErr(err, "branch ", branchName, " does not exist, attempting to create it")

			err = workTree.Checkout(&git.CheckoutOptions{
				Branch: plumbing.NewBranchReferenceName(branchName),
				Hash:   plumbing.ZeroHash,
				Create: true,
			})

			if err != nil {
				log.ErrorErr(err, "could not create the branch ", branchName)
				return *success
			}
			log.Info("created the metadata branch ", branchName)

			*success = true
			return *success
		} else {
			log.ErrorErr(err, "encountered an error when switching to the branch ", branchName)
		}
	} else {
		log.Trace("branch exists, checking out")

		err = workTree.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewBranchReferenceName(branchName),
		})
		if err != nil {
			log.ErrorErr(err, "could not check out the metadata branch")
			return *success
		}
	}

	if err != nil {
		log.ErrorErr(err, "could not checkout the branch ", branchName)
		return *success
	}
	*success = true
	return *success
}
