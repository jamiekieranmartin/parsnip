# parsnip

Parsnip is a minimal text to JSON converter.

By default, parsnip converts input text to JSON based on a given expression

```bash
$ parsnip "(\S+) (\S+)" "Jamie Martin"
{"1":"Jamie","2":"Martin"}
```

Named groups can be used to map key-value pairs.

```bash
$ parsnip "(?P<first>\S+) (?P<last>\S+)" "Jamie Martin"
{"first":"Jamie","last":"Martin"}
```

Write to file

```bash
$ parsnip -out "./result.json" "(?P<first>\S+) (?P<last>\S+)" "Jamie Martin"
```
