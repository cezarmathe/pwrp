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

package keys

const (
	recordingBaseKey = "recording"

	/*RecordingRepositoryListKey is a key for the recording repository list*/
	RecordingRepositoryListKey = recordingBaseKey + separator + "repositories"
	/*RecordingProtocolKey is a key for the recording repository cloning protocol*/
	RecordingProtocolKey = recordingBaseKey + separator + "protocol"

	recordingSkipsBaseKey = recordingBaseKey + separator + "skips"

	/*RecordingSkipsMissingBranchKey is a key for the recording skip for a missing branch*/
	RecordingSkipsMissingBranchKey = recordingSkipsBaseKey + separator + "missing_branch"
	/*RecordingSkipsBadURLKey is a key for the recording skip for a bad URL*/
	RecordingSkipsBadURLKey = recordingSkipsBaseKey + separator + "bad_url"
	/*RecordingSkipsBadProtocolKey is a key for the recording skip for a bad protocol*/
	RecordingSkipsBadProtocolKey = recordingSkipsBaseKey + separator + "bad_protocol"
	/*RecordingSkipsAllKey is a key for the recording skip for all errors*/
	RecordingSkipsAllKey = recordingSkipsBaseKey + separator + "all"
)
