# Fasturtle

A detokenization tool for configuration files and Madlibs, with support for Chef data bags.

## Installation

```bash
go get -u github.com/vacoj/fasturtle
```

## Usage

Pro-tip: Don't use ```$``` or ```*``` in your buffer characters.  It's just a bad idea, and I'm not accounting for it.

### Detokenize syntax

#### Local token files in JSON format

Note:  this method supports a form of value overrides.  Values duplicated in JSON files later in the list will override values declared in earlier in the list.

```bash
fasturtle --input="input.xml" --output="output.xml" --tokens="token.json,token2.json" --buffer="__"
```

#### Using *Unencrypted* Chef Data Bags

```bash
fasturtle --input="input.xml" --output="output.xml" --databag="_default" --buffer="__"
```

#### Using *Encrypted* Chef Data Bags

```bash
fasturtle --input="input.xml" --output="output.xml" --databag="_default" --data-secret="my_secret_file" --buffer="__"
```

### Extract tokens syntax

```bash
fasturtle --extract --buffer="__" --input="input.xml" --output="output.json"
# --output is optional, but if provided, and if the file extension ends in ".json", the keys will be formatted as JSON.
# for example: {"myintkey":"","mykey":"","myotherkey":""}
# This makes it easy to create a tokens file template, since you only need to plug in the values.
```

## Examples

### Detokenize

Token File:

```json
file 1
{
    "mykey": "some important value",
    "myotherkey": "this other thing",
    "myintkey": 444,
    "pirri": "http://pirri.vacovsky.us"
}
file 2
{
    "secondfilekey": "second file, so crazy!",
    "myintkey": 15,
}
```

Tokenized Input:

```xml
<someXML name="something" value=__mykey__>
<someXML name="secondfilekey" value=__secondfilekey__>
<someXML name="theotherthing" value=__myotherkey__>
<someXML name="thisshouldbeanint" value=__myintkey__>
<someXML name="pirri" value=__pirri__>
```

Detokenized Output:

```xml
<someXML name="something" value="some important value">
<someXML name="secondfilekey" value="second file, so crazy!">
<someXML name="theotherthing" value="this other thing">
<someXML name="thisshouldbeanint" value=15>  <!-- this value is 15 because token2.json overrode the value of __myintkey__ -->
<someXML name="pirri" value="http://pirri.vacovsky.us">
```

### Extract tokens

Tokenized Input:

```xml
<someXML name="something" value=__mykey__>
<someXML name="theotherthing" value=__myotherkey__>
<someXML name="thisshouldbeanint" value=__myintkey__>
```

Outputs a list of used tokens:

```bash
__mykey__
__myotherkey__
__myintkey__
```

or

```javascript
{"myintkey":"","mykey":"","myotherkey":""}  // notice we remove the buffer chars
```
