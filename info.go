// Copyright (c) 2018 Brian Allred
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package aurpc

// Info object returned from a package info request
type Info struct {
	Version     int    `json:"version"`
	Type        string `json:"type"`
	Resultcount int    `json:"resultcount"`
	Results     []struct {
		ID             int           `json:"ID"`
		Name           string        `json:"Name"`
		PackageBaseID  int           `json:"PackageBaseID"`
		PackageBase    string        `json:"PackageBase"`
		Version        string        `json:"Version"`
		Description    string        `json:"Description"`
		URL            string        `json:"URL"`
		NumVotes       int           `json:"NumVotes"`
		Popularity     float64       `json:"Popularity"`
		OutOfDate      interface{}   `json:"OutOfDate"`
		Maintainer     string        `json:"Maintainer"`
		FirstSubmitted int           `json:"FirstSubmitted"`
		LastModified   int           `json:"LastModified"`
		URLPath        string        `json:"URLPath"`
		Depends        []string      `json:"Depends"`
		MakeDepends    []string      `json:"MakeDepends"`
		License        []string      `json:"License"`
		Keywords       []interface{} `json:"Keywords"`
	} `json:"results"`
}
