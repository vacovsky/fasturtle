#!/bin/sh
go get -u github.com/vacoj/fasturtle
fasturtle --buffer="__" --input=tokenizedDemo.xml --output=myoutput_multitokenjson.xml --tokens=tokensDemo.json,tokensDemo2.json
fasturtle --buffer="__" --input=tokenizedDemo.xml --output=myoutput_databag.xml --databag=_default
fasturtle --buffer=__ --input=tokenizedDemo.xml --databag=fasturtle --databag-secret=dbsecret --output=results.xml
