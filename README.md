# Fasturtle

A detokenization tool for configuration files and Madlibs.

## Usage

```bash
fasturtle --buffer="__" --input="/path/to/tokenized/file" --output="/path/to/save/output" --tokens="/path/to/json/keyvals"
```

## Example

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