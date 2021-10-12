# Cisco BCS (Cisco Business Critical Services)

[![Status](https://img.shields.io/badge/status-wip-yellow)](https://github.com/darrenparkinson/bcs) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/darrenparkinson/bcs) ![GitHub](https://img.shields.io/github/license/darrenparkinson/bcs?color=brightgreen) [![GoDoc](https://pkg.go.dev/badge/darrenparkinson/bcs)](https://pkg.go.dev/github.com/darrenparkinson/bcs) [![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/bcs)](https://goreportcard.com/report/github.com/darrenparkinson/bcs)

This repository consists of a library and CLI utility for extracting information from the Cisco BCS Insights API.

* Docs for this API here: https://api.csco-bcs.com/v2/
* Postman Collection here: https://www.getpostman.com/collections/c2abd62e08b6d27a7116 

## Using the CLI

```sh
go get github.com/darrenparkinson/bcs/cmd/bcs-cli
```

Currently there are only two commands implemented:

* `download` - will download bulk data given a customer ID and API key
* `parse` - will parse a downloaded file to provide basic stats

You can see detailed help as follows:

* `$ bcs-cli download --help`
* `$ bcs-cli parse --help`

## Using the library

```sh
go get github.com/darrenparkinson/bcs/pkg/ciscobcs
```

Currently only basic bulk download capability implemented.  Ideally we'll implement all the services.