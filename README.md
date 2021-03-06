## Server for web preview using Headless Chrome written in Go.

[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/jimmy-go/cropper.svg?branch=master)](https://travis-ci.org/jimmy-go/cropper)
[![Go Report Card](https://goreportcard.com/badge/github.com/jimmy-go/cropper)](https://goreportcard.com/report/github.com/jimmy-go/cropper)
[![GoDoc](http://godoc.org/github.com/jimmy-go/cropper?status.png)](http://godoc.org/github.com/jimmy-go/cropper)
[![Coverage Status](https://coveralls.io/repos/github/jimmy-go/cropper/badge.svg?branch=master)](https://coveralls.io/github/jimmy-go/cropper?branch=master)

### Install:
```
go get github.com/jimmy-go/cropper
```

### Usage:

```
$ make run
```

* Exposed endpoint in http://localhost:9900/v1/preview?url=string&width=int&height=int

See `scripts/run.sh` for customization.

### License:

MIT License

Copyright (c) 2017 Angel del Castillo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
