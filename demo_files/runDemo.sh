#!/bin/sh
go get -u github.com/vacoj/fasturtle
../fasturtle --buffer-left="{{" -buffer-right="}}" --input=tokenizedDemo.xml --output=myoutput_multitokenjson.xml --tokens=tokensDemo.json,tokensDemo2.json
../fasturtle --buffer-left="{{" -buffer-right="}}" --input=tokenizedDemo.xml --output=myoutput_databag.xml --databag=_default
../fasturtle --buffer-left="{{" -buffer-right="}}" --input=tokenizedDemo.xml --databag=fasturtle --databag-secret=dbsecret --output=results.xml
