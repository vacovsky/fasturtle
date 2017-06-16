# Fasturtle

A detokenization tool for configuration files and Madlibs.

## Usage

```bash
fasturtle --buffer="$$" --input="/path/to/tokenized/file" --output="/path/to/save/output" --tokens="/path/to/json/keyvals"
```

## Example

You can run this example via ```runDemo.sh``` in this repository.

Token File:

```json
{
    "mykey": "some important value",
    "myotherkey": "this other thing"
}
```

Tokenized Input:

```xml
<someXML name="something" value="$$mykey$$">
<someXML name="theotherthing" value="$$myotherkey$$">
```

Detokenized Output:

```xml
<someXML name="something" value="some important value">
<someXML name="theotherthing" value="this other thing">
```