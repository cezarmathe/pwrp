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
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

/*getEntry creates a log entry with the proper fields*/
func (log *SmartLogger) getEntry(debug bool) *logrus.Entry {
	var fields logrus.Fields

	/*if a subtag exists, add it to the fields then set it as empty in the log object*/
	if log.currentSubTag != "" {
		fields = logrus.Fields{
			"tag":    log.tag,
			"subtag": log.currentSubTag,
		}
		log.currentSubTag = ""
	} else {
		fields = logrus.Fields{
			"tag": log.tag,
		}
	}

	/*if using the debug logger, return an entry based on the debug logger*/
	if debug {
		return log.debugLogger.WithFields(fields)
	}
	return log.logger.WithFields(fields)
}

/*Trace logs a message on the Trace level*/
func (log *SmartLogger) Trace(args ...string) {
	log.getEntry(false).Trace(strings.Join(args, " "))
}

/*Debug logs a message on the Debug level is debug logging is enabled*/
func (log *SmartLogger) Debug(args ...string) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	log.getEntry(true).WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Debug(name, "(): ", strings.Join(args, " "))
}

/*Info logs a message on the Info level*/
func (log *SmartLogger) Info(args ...string) {
	log.getEntry(false).Info(strings.Join(args, " "))
}

/*Warn logs a message on the Warn level*/
func (log *SmartLogger) Warn(args ...string) {
	log.getEntry(false).Warn(strings.Join(args, " "))
}

/*WarnErr logs a message on the Warn level with the specified error*/
func (log *SmartLogger) WarnErr(err error, args ...string) {
	log.getEntry(false).WithError(err).Warn(strings.Join(args, " "))
}

/*Fatallogs a message on the Fatal level*/
func (log *SmartLogger) Fatal(args ...string) {
	log.getEntry(false).Fatal(strings.Join(args, " "))
}

/*FatalErr logs a message on the Fatal level with the specified error*/
func (log *SmartLogger) FatalErr(err error, args ...string) {
	log.getEntry(false).WithError(err).Fatal(strings.Join(args, " "))
}

/*
WithSubTag writes a subtag when the following log method is called
It is intended to be used as log.WithSubTag("my_subtag").Trace("message")
*/
func (log *SmartLogger) WithSubTag(subtag string) *SmartLogger {
	log.currentSubTag = subtag
	return log
}
