# ethofs-cli
Command-line Interface/Access to The ethoFS Upload System/Network

## Pre-built Binaries

```shell
build/bin/ethofs-cli-windows
build/bin/ethofs-cli-macos
build/bin/ethofs-cli-linux
```
## Building the source

```shell
go build -o build/bin/ethofs-cli-linux
```
## Running pre-built binary (register new ethoFS account) (Work-in-Progress)

```shell
./build/bin/ethofs-cli-linux -register -name={AccountName} -key={PrivateKey}
```

## Running pre-built binary (list currently hosted data)

```shell
./build/bin/ethofs-cli-linux -list -key={PrivateKey}
```

## Running pre-built binary (standard upload)

```shell
./build/bin/ethofs-cli-linux -upload -path=test.txt -key={PrivateKey}
```

## Running pre-built binary (recursive directory upload)

```shell
./build/bin/ethofs-cli-linux -r -upload -path=testDir -key={PrivateKey}
```

## Running pre-built binary (remove upload) (Work-in-Progress)

```shell
./build/bin/ethofs-cli-linux -remove -contractaddress={HostingContractAddress} -key={PrivateKey}
```

## Running pre-built binary (upload contract extension) (Work-in-Progress)

```shell
./build/bin/ethofs-cli-linux -extend -blocks={ExtensionBlockCount} -contractaddress={HostingContractAddress} -key={PrivateKey}
```

![Upload Example](ethofs-cli.png)
