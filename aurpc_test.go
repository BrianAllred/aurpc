// Copyright (c) 2018 Brian Allred
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package aurpc_test

import (
	"testing"

	"github.com/BrianAllred/aurpc"
)

func TestSearchByName(t *testing.T) {
	search, err := aurpc.SearchByName("linux-ck")
	if err != nil || search.Resultcount < 1 {
		t.Fail()
	}
}

func TestSearchByNameDesc(t *testing.T) {
	search, err := aurpc.SearchByNameDesc("ck1 patchset")
	if err != nil || search.Resultcount < 1 {
		t.Fail()
	}
}

func TestSearchByMaintainer(t *testing.T) {
	search, err := aurpc.SearchByMaintainer("graysky")
	if err != nil || search.Resultcount < 1 {
		t.Fail()
	}
}

func TestPkgInfo(t *testing.T) {
	info, err := aurpc.PkgInfo("linux-ck")
	if err != nil || info.Resultcount < 1 {
		t.Fail()
	}
}

func TestMultiplePkgInfo(t *testing.T) {
	info, err := aurpc.PkgInfo("linux-ck", "linux-ck-headers")
	if err != nil || info.Resultcount < 2 {
		t.Fail()
	}
}
