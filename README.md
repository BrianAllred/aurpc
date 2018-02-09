# aurpc

A Go library for accessing the Arch User Repository via the AUR RPC

[![CircleCI](https://circleci.com/gh/BrianAllred/aurpc.svg?style=svg)](https://circleci.com/gh/BrianAllred/aurpc) [![Go Report Card](https://goreportcard.com/badge/github.com/brianallred/aurpc)](https://goreportcard.com/report/github.com/brianallred/aurpc)

## Get the package

`go get github.com/BrianAllred/aurpc`

## Use the code

* Search by name, name and description, or maintiner:

        search, err := aurpc.SearchByName("linux-ck")
        search, err := aurpc.SearchByNameDesc("ck1 patchset")
        search, err := aurpc.SearchByMaintainer("graysky")

* Get package info:

        info, err := aurpc.PkgInfo("linux-ck")
        info, err := aurpc.PkgInfo("linux-ck", "linux-ck-headers")