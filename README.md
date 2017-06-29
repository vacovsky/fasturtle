# Fasturtle - a detokenization tool

A detokenization tool for configuration files and Madlibs, with support for Chef data bags.

This tool's input files can be easily generated from existing web.config files by using <a href="https://github.com/vacoj/shinroller">ShinRoller</a>. (https://github.com/vacoj/shinroller)

## Installation

```bash
go get -u github.com/vacoj/fasturtle
go get -u github.com/vacoj/shinroller   # recommended
```

## Usage

### CLI Help

```text
Usage of fasturtle:

  -buffer string
        Characters used to buffer the keys within the input file.  The default value is an empty string.  Example: "__mykey__" (not used if --buffer-left or --buffer-right are provided).

  -buffer-left string
        Characters used to buffer the keys within the input file on the left side of a token key.  The default value is an empty string.  Example: "{{mykey" (usually used in conjunction with --buffer-left).

  -buffer-right string
        Characters used to buffer the keys within the input file on the right side of a token key.  The default value is an empty string.  Example: "mykey}}" (usually used in conjunction with --buffer-left).

  -databag string
        Name of the Chef data bag containing the tokenized values.  Under the hood, this relies on your environment having a properly configured knife.rb and necessary certs in place to connect to the Chef server.  Alternately, use --tokens to specify a json file.
        
  -databag-secret string
        Path to the data bag secret.  Only necessary if you use encrypted data bags.

  -extract
        If true, enters extract mode.  In extract mode, the output file or stdout becomes a list of the tokens found within the input file.

  -input string
        Path the tokenized input file.

  -output string
        Destination path and file name for the detokenized file.  If not set, detokenized file is printed to stdout.

  -tokens string
        Path to the JSON key-value pair set(s) to be used for detokenization of the input file.  For multiple files, separate file paths with a comma (,).  If the same key exists in one or more of the files, then the file furthest to the right takes precendece (overrides) the previous.
  -unsafe
        If true, will not throw error if all tokens are not replaced.  Default is false, and if a token still exists after detokenization, an error will be thro
wn.
```

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
fasturtle --extract --buffer-left="__" --buffer-right="__" --input="input.xml" --output="output.json"
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
