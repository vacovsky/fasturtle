# Fasturtle

A detokenization tool for configuration files and Madlibs.

## Installation

```bash
go get github.com/vacoj/fasturtle
```
## Usage

Pro-tip: Don't use ```$``` or ```*``` in your buffer characters.  It's just a bad idea, and I'm not accounting for it.

### Detokenize syntax

#### Local token files in JSON format

```bash
fasturtle --input="input.xml" --output="output.xml" --tokens="token.json,token2.json" --buffer="__"
```

#### Using Chef Data Bags

```bash
fasturtle --input="input.xml" --output="output.xml" --databags="_default" --buffer="__"
```

### Extract tokens syntax

```bash
fasturtle --extract --input="input.xml" --output="output.json"
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
    "pirri": "pirri url"
}
file 2
{
    "secondfilekey": "second file, so crazy!"
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
<someXML name="thisshouldbeanint" value=444>
<someXML name="pirri" value="pirri url">
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
{"myintkey":"","mykey":"","myotherkey":""}
```
