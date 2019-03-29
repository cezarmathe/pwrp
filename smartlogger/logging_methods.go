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
func (log *SmartLogger) Trace(args ...interface{}) {
	log.getEntry(false).Trace(args...)
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

/*DebugFunctionCalled logs a message on Debug level stating that the calling function was called.*/
func (log *SmartLogger) DebugFunctionCalled(params ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()

	entry := log.getEntry(true).WithFields(logrus.Fields{
		"file": file,
		"line": line,
	})

	if len(params) > 0 {
		entry.Debug(name, "(): ", "called with ", params)
	} else {
		entry.Debug(name, "(): ", "called")
	}
}

/*DebugFunctionCalled logs a message on Debug level stating that the calling function returned.*/
func (log *SmartLogger) DebugFunctionReturned(returnValues ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()

	entry := log.getEntry(true).WithFields(logrus.Fields{
		"file": file,
		"line": line,
	})

	if len(returnValues) > 0 {
		entry.Debug(name, "(): ", "returned ", returnValues)
	} else {
		entry.Debug(name, "(): ", "returned")
	}
}

/*Info logs a message on the Info level*/
func (log *SmartLogger) Info(args ...interface{}) {
	log.getEntry(false).Info(args...)
}

/*Warn logs a message on the Warn level*/
func (log *SmartLogger) Warn(args ...interface{}) {
	log.getEntry(false).Warn(args...)
}

/*WarnErr logs a message on the Warn level with the specified error*/
func (log *SmartLogger) WarnErr(err error, args ...interface{}) {
	if len(args) > 0 {
		log.getEntry(false).WithError(err).Warn(args...)
	} else {
		log.getEntry(false).Warn(err)
	}
}

/*Error logs a message on the Error level*/
func (log *SmartLogger) Error(args ...interface{}) {
	log.getEntry(false).Error(args...)
}

/*ErrorErr logs a message on the Error level with the specified error*/
func (log *SmartLogger) ErrorErr(err error, args ...interface{}) {
	if len(args) > 0 {
		log.getEntry(false).WithError(err).Error(args...)
	} else {
		log.getEntry(false).Error(err)
	}
}

/*Fatal logs a message on the Fatal level*/
func (log *SmartLogger) Fatal(args ...interface{}) {
	log.getEntry(false).Fatal(args...)
}

/*FatalErr logs a message on the Fatal level with the specified error*/
func (log *SmartLogger) FatalErr(err error, args ...interface{}) {
	if len(args) > 0 {
		log.getEntry(false).WithError(err).Fatal(args...)
	} else {
		log.getEntry(false).Fatal(err)
	}
}

/*
WithSubTag writes a subtag when the following log method is called
It is intended to be used as log.WithSubTag("my_subtag").Trace("message")
*/
func (log *SmartLogger) WithSubTag(subtag string) *SmartLogger {
	log.currentSubTag = subtag
	return log
}
