#!/bin/bash

filename=".env"
appname="frozen-sierra-65437"

while read line; do
    echo "line: $line"
    heroku config:add $line --app $appname
done < $filename