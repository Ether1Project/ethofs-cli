# ethofs-cli
Command-line Interface/Access to The ethoFS Upload System/Network

## Building the source

```shell
go build -o build/bin/ethofs-cli
```

## Running pre-built binary (standard upload)

```shell
./build/bin/ethofs-cli -upload -path=test.txt
```

## Running pre-built binary (recursive directory upload)

```shell
./build/bin/ethofs-cli -r -upload -path=testDir
```
