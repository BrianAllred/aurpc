// Copyright (c) 2018 Brian Allred
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package aurpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var rpcUri = "https://aur.archlinux.org/rpc/?v=5"
var client = &http.Client{
	Timeout: time.Second * 10,
}

func aurSearch(keywords string, byField string) (*Search, error) {
	if len(keywords) < 1 {
		return nil, errors.New("invalid keywords")
	}

	var queries = "&type=search&by=" + byField + "&arg=" + keywords

	var searchUri = rpcUri + queries

	resp, err := client.Get(searchUri)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()

	var search Search
	var rpcError Error

	if err := json.Unmarshal([]byte(responseBody), &search); err != nil {
		return nil, err
	}

	if search.Type == "error" {
		if err := json.Unmarshal([]byte(responseBody), &rpcError); err != nil {
			return nil, err
		}

		return nil, errors.New(rpcError.Error)
	}

	return &search, nil
}

func SearchByName(keywords string) (*Search, error) {
	return aurSearch(keywords, "name")
}

func SearchByNameDesc(keywords string) (*Search, error) {
	return aurSearch(keywords, "name-desc")
}

func SearchByMaintainer(keywords string) (*Search, error) {
	return aurSearch(keywords, "maintainer")
}

func PkgInfo(packages ...string) (*Info, error) {
	if len(packages) < 1 {
		return nil, errors.New("invalid packages")
	}

	var queries = "&type=info"

	for _, pkg := range packages {
		queries += "&arg[]=" + pkg
	}

	var infoUri = rpcUri + queries

	resp, err := client.Get(infoUri)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()

	var info Info
	var rpcError Error

	if err := json.Unmarshal([]byte(responseBody), &info); err != nil {
		return nil, err
	}

	if info.Type == "error" {
		if err := json.Unmarshal([]byte(responseBody), &rpcError); err != nil {
			return nil, err
		}

		return nil, errors.New(rpcError.Error)
	}

	return &info, nil
}
