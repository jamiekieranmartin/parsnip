# parsnip

Parsnip is a minimal text to key-value converter.

## Install

```bash
go get github.com/jamiekieranmartin/parsnip
```

Otherwise you can download the binary from [Releases](https://github.com/jamiekieranmartin/parsnip/releases)

## Usage

### CLI

```bash
parsnip "(\S+) (\S+)" "Jamie Martin"
```

### Golang SDK

```go
// parse input given expression
parsed, err := parsnip.Parse("(\S+) (\S+)", "Jamie Martin")
if err != nil {
  panic(err)
}

fmt.Println(parsed)
```

## CLI flags

### `-out`

Output to JSON file. Defaults to none.

```bash
parsnip -out "./result.json" "(?P<first>\S+) (?P<last>\S+)" "Jamie Martin"
```

## Configuration

By default, parsnip converts input text to JSON based on a given regular expression.

```bash
$ parsnip "(\S+) (\S+)" "Jamie Martin"
{"1":"Jamie","2":"Martin"}
```

Named groups can be used to map key-value pairs.

```bash
$ parsnip "(?P<first>\S+) (?P<last>\S+)" "Jamie Martin"
{"first":"Jamie","last":"Martin"}
```

See [parsnip_test.go](./parsnip_test.go) for more examples.
