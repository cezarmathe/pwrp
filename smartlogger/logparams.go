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

package smartlogger

import "github.com/sirupsen/logrus"

/*LogParams is a container for easily transferring logging parameters*/
type LogParams struct {
	Debug bool
	Level logrus.Level
}

/*GetParams returns the parameters of a SmartLogger*/
func (log *SmartLogger) GetParams() LogParams {
	return LogParams{
		log.enableDebug,
		log.logLevel,
	}
}
