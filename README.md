# gohan

[![Build Status](https://travis-ci.org/oirik/gohan.svg?branch=master)](https://travis-ci.org/oirik/gohan)
[![GitHub release](https://img.shields.io/github/release/oirik/gohan.svg)](RELEASE)
[![apache license](https://img.shields.io/badge/license-Apache-blue.svg)](LICENSE)

http (https) server for both static files and proxy, written by golang.

## Install

```
$ go get github.com/oirik/gohan
```

Or download binaries from [github releases](https://github.com/oirik/gohan/releases)

## Usage

Execute simply. gohan listens at 8080 and response from current directory.

```
$ gohan
```

Or check options for details.

```
$ gohan -h

Usage of gohan:
  -certFile string
        SSL cert file path (default "cert.pem")
  -keyFile string
        SSL key file path (default "key.pem")
  -path string
        static files directory (default ".")
  -port int
        http server port number (default 8080)
  -proxy string
        reverse proxy destination host. ex) localhost:8080
  -ssl
        use SSL

```
