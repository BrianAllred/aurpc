// Copyright (c) 2018 Brian Allred
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package aurpc

// Search object returned from a package search request
type Search struct {
	Version     int    `json:"version"`
	Type        string `json:"type"`
	Resultcount int    `json:"resultcount"`
	Results     []struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"results"`
}
