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
	logger        *logrus.Logger
	debugLogger   *logrus.Logger
	enableDebug   bool
	logLevel      *logrus.Level
	tag           string
	currentSubTag string
}

/*NewSmartLogger creates a new SmartLogger*/
func NewSmartLogger(enableDebug bool, level logrus.Level, tag string) *SmartLogger {
	smartLogger := new(SmartLogger)

	smartLogger.logger = logrus.New()
	smartLogger.logger.SetLevel(level)

	smartLogger.enableDebug = enableDebug
	smartLogger.debugLogger = logrus.New()
	smartLogger.debugLogger.SetLevel(logrus.DebugLevel)
	smartLogger.debugLogger.SetReportCaller(false)

	smartLogger.tag = tag

	return smartLogger
}

/*FromLogParams creates a new SmartLogger with the given parameters*/
func FromLogParams(params LogParams, tag string) *SmartLogger {
	smartLogger := new(SmartLogger)

	smartLogger.logger = logrus.New()
	smartLogger.logger.SetLevel(params.Level)

	smartLogger.enableDebug = params.Debug
	smartLogger.debugLogger = logrus.New()
	smartLogger.debugLogger.SetLevel(logrus.DebugLevel)
	smartLogger.debugLogger.SetReportCaller(false)

	smartLogger.tag = tag

	return smartLogger
}

/*EnableDebug enables logging on the debug level*/
func (log *SmartLogger) EnableDebug(enableDebug bool) {
	log.enableDebug = enableDebug
}

/*SetLevel sets the logging level for the standard logger*/
func (log *SmartLogger) SetLevel(level logrus.Level) {
	log.logger.SetLevel(level)
}

/*SetTag sets the standard tag for the loggers*/
func (log *SmartLogger) SetTag(tag string) {
	log.tag = tag
}
