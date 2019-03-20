/*
	PWRP - Personal Work Recorder Processor
	Copyright (C) 2019  Cezar Mathe <cezarmathe@gmail.com> [https://cezarmathe.com]

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published
	by the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package config

/*ErrCreateStorageDir describes the creation of a certain storage directory*/
type ErrCreateStorageDir struct {
	Dirname string
}

/*NewErrCreateStorageDir returns a new ErrCreateStorageDir*/
func NewErrCreateStorageDir(dirname string) *ErrCreateStorageDir {
	return &ErrCreateStorageDir{dirname}
}

func (err *ErrCreateStorageDir) Error() string {
	return "encountered an error when creating the storage directory at \"" + err.Dirname + "\""
}

/*ErrNoPermissions describes a missing branch*/
type ErrNoPermissions struct {
	Path string
}

/*NewErrNoPermissions returns a new ErrNoPermissions*/
func NewErrNoPermissions(path string) *ErrNoPermissions {
	return &ErrNoPermissions{path}
}

func (err *ErrNoPermissions) Error() string {
	return "no permissions for writing data at " + err.Path
}
