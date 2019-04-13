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

import "net/http"

func httpRequest(generator *Generator) ([]byte, error) {

	url, err := generator.getArg("url")
	if err != nil {
		return nil, err
	}

	response, err := http.Get(url.(string))
	if err != nil {
		return nil, err
	}

	byteResponse := make([]byte, response.ContentLength)

	_, err = response.Body.Read(byteResponse)

	if err != nil {
		return nil, err
	}

	return byteResponse, nil
}
