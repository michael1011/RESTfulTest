# RESTfulTest [![Jenkins](https://michael1011.at/jenkins/buildStatus/icon?job=RESTfulTest)](https://michael1011.at/jenkins/job/RESTfulTest/) [![Travis](https://travis-ci.org/michael1011/RESTfulTest.svg?branch=master)](https://travis-ci.org/michael1011/RESTfulTest)
## An application to test RESTful services

Find out how to use RESTfulTest: `./RESTfulTest help`

### How to install

You can find the latest releases [here](https://github.com/michael1011/RESTfulTest/releases). If there isn't a compiled version that fits your system or architecture you can compile it yourself. Make sure [Go](https://golang.org/) is properly installed. Download the source code under the releases and execute:

```
go get github.com/yosssi/gohtml
go get github.com/fatih/color
go get github.com/jteeuwen/go-bindata
go get github.com/elazarl/go-bindata-assetfs

go-bindata-assetfs public/ public/css public/js

go build
````
