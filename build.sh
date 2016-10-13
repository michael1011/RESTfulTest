#!/bin/sh

deleteFile() {
    if [ -f "$1" ];
    then
       rm -f $1
    fi

}

deleteFile bindata_assetfs.go
deleteFile RESTfulTest

go-bindata-assetfs build.txt public/ public/css public/js

go build