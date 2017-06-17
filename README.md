# Fasturtle

A detokenization tool for configuration files and Madlibs.

## Usage

### Detokenize syntax
```bash
fasturtle --buffer="__" --input="/path/to/tokenized/file" --output="/path/to/save/output" --tokens="/path/to/json/keyvals"
```


### Extract tokens syntax

```bash
./fasturtle --input="/path/to/tokenized/file" --buffer="__" --extract=true
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