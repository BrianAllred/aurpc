// Copyright (c) 2018 Brian Allred
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package aurpc

type Error struct {
	Version     int           `json:"version"`
	Type        string        `json:"type"`
	Resultcount int           `json:"resultcount"`
	Results     []interface{} `json:"results"`
	Error       string        `json:"error"`
}
