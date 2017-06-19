#!/bin/sh

./fasturtle_x64 --buffer="__" --input="./tokenizedDemo.xml" --output="myoutput_multitokenjson.xml" --tokens="./tokensDemo.json,tokensDemo2.json"
./fasturtle_x64 --buffer="__" --input="./tokenizedDemo.xml" --output="myoutput_databag.xml" --databag="_default"
