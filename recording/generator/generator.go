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

/*ArgUse tells the generator what is the application of a certain argument.*/
type ArgUsage string

const (
	FlagArgUsage            ArgUsage = "flag"
	EnvironmentArgUsage     ArgUsage = "environment"
	CommandArgumentArgUsage ArgUsage = "command_arg"
)

/*Args is a container for generator arguments.*/
type Args struct {
	Usage ArgUsage    `toml:"usage"`
	Key   string      `toml:"key"`
	Value interface{} `toml:"value"`
}

/*Generator is a container for a task that runs on a project at the moment of recording.*/
type Generator struct {
	Key  string `toml:"report_key"`
	Task string `toml:"task"`
	Path string `toml:"path"`
	Args []Args `toml:"args"`
}

func (generator *Generator) Run() error {

}
