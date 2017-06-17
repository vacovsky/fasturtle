# Fasturtle

A detokenization tool for configuration files and Madlibs.

## Installation

```bash
go get github.com/vacoj/fasturtle
```
## Usage

### Detokenize syntax
```bash
fasturtle --input="input.xml" --output="output.xml" --tokens="token.json" --buffer="__" 
```

### Extract tokens syntax

```bash
fasturtle --extract --input="input.xml"
```

## Examples

### Detokenize
Token File:

```json
{
    "mykey": "some important value",
    "myotherkey": "this other thing",
    "myintkey": 444
}
```

Tokenized Input:

```xml
<someXML name="something" value=__mykey__>
<someXML name="theotherthing" value=__myotherkey__>
<someXML name="thisshouldbeanint" value=__myintkey__>
```

Detokenized Output:

```xml
<someXML name="something" value="some important value">
<someXML name="theotherthing" value="this other thing">
<someXML name="thisshouldbeanint" value=444>
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