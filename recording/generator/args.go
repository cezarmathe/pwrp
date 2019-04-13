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

package generator

import "github.com/pkg/errors"

/*ArgUse tells the generator what is the application of a certain argument.*/
type ArgUsage string

const (
	FlagArgUsage            ArgUsage = "flag"
	EnvironmentArgUsage     ArgUsage = "environment"
	CommandArgumentArgUsage ArgUsage = "command_arg"
	SpecialArgUsage         ArgUsage = "special"
)

var (
	ErrArgNotFound = errors.New("arg not found")
)

/*Args is a container for generator arguments.*/
type Args struct {
	Usage ArgUsage    `toml:"usage"`
	Key   string      `toml:"key"`
	Value interface{} `toml:"value"`
}

func (generator *Generator) getArg(key string) (interface{}, error) {
	for _, arg := range generator.Args {
		if arg.Key == key {
			return arg.Value, nil
		}
	}
	return nil, ErrArgNotFound
}
