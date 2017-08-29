#!/bin/sh
go get -u github.com/vacoj/fasturtle

../fasturtle --buffer-left="{{" -buffer-right="}}" --input=tokenizedDemo.LR-buffer.xml --output=myoutput_multitokenjson_LR-buffer.xml --tokens=tokensDemo.json,tokensDemo2.json
../fasturtle --buffer="__" --input=tokenizedDemo.xml --output=myoutput_multitokenjson.xml --tokens=tokensDemo.json,tokensDemo2.json
# ../fasturtle --buffer-left="{{" -buffer-right="}}" --input=tokenizedDemo.xml --output=myoutput_databag.xml --databag=_default
# ../fasturtle --buffer-left="{{" -buffer-right="}}" --input=tokenizedDemo.xml --databag=fasturtle --databag-secret=dbsecret --output=results.xml

# extract test/demo
../fasturtle --extract --buffer="__" --input=tokenizedDemo.xml --output=myoutput_multitokenjson.json
../fasturtle --extract --buffer-left="{{" --buffer-right="}}" --input=tokenizedDemo.LR-buffer.xml --output=myoutput_multitokenjson.LR.json

../fasturtle --buffer="__" --input=tokenizedDemo.xml --output=myoutput_multitokenjson_missing.xml --tokens=tokensDemo_missing.json

../fasturtle --buffer="__" --input=tokenizedDemo.xml --assembly-bindings-source=tokenizedDemoAssemblies.xml --output=myoutput_multitokenjson_assemblies.xml --tokens=tokensDemo.json,tokensDemo2.json