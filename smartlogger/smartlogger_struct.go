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

import (
	"github.com/sirupsen/logrus"
)

/*SmartLogger is a struct that handles regular and debug logging*/
type SmartLogger struct {
	logger      *logrus.Logger
	debugLogger *logrus.Logger
	enableDebug bool
	logLevel    *logrus.Level
}

/*NewSmartLogger creates a new SmartLogger*/
func NewSmartLogger(enableDebug bool, level logrus.Level) *SmartLogger {
	smartLogger := new(SmartLogger)

	smartLogger.logger = logrus.New()
	smartLogger.logger.SetLevel(level)

	if enableDebug {
		smartLogger.debugLogger = logrus.New()
		smartLogger.debugLogger.SetReportCaller(true)
	}

	return smartLogger
}
