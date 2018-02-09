// Copyright (c) 2018 Brian Allred
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package aurpc

// Error object returned from a bad RPC request
type Error struct {
	Version     int           `json:"version"`
	Type        string        `json:"type"`
	Resultcount int           `json:"resultcount"`
	Results     []interface{} `json:"results"`
	Error       string        `json:"error"`
}
