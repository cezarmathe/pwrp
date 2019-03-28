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

/*Trace logs a message on the Trace level*/
func (log *SmartLogger) Trace(args ...interface{}) {
	log.logger.Trace(args)
}

/*Debug logs a message on the Debug level is debug logging is enabled*/
func (log *SmartLogger) Debug(args ...interface{}) {
	if log.enableDebug {
		log.debugLogger.Debug(args)
	}
}

/*Info logs a message on the Info level*/
func (log *SmartLogger) Info(args ...interface{}) {
	log.logger.Info(args)
}

/*Warn logs a message on the Warn level*/
func (log *SmartLogger) Warn(args ...interface{}) {
	log.logger.Warn(args)
}

/*WarnErr logs a message on the Warn level with the specified error*/
func (log *SmartLogger) WarnErr(err error, args ...interface{}) {
	log.logger.WithError(err).Warn(args)
}

/*Fatallogs a message on the Fatal level*/
func (log *SmartLogger) Fatal(args ...interface{}) {
	log.logger.Fatal(args)
}

/*FatalErr logs a message on the Fatal level with the specified error*/
func (log *SmartLogger) FatalErr(err error, args ...interface{}) {
	log.logger.WithError(err).Fatal(args)
}
