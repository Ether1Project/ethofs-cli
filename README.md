# ethofs-cli
Command-line Interface/Access to The ethoFS Upload System/Network

## Building the source

```shell
go build -o build/bin/ethofs-cli
```
## Running pre-built binary (register new ethoFS account)

```shell
./build/bin/ethofs-cli -register -name={AccountName} -key={PrivateKey}
```

## Running pre-built binary (list currently hosted data)

```shell
./build/bin/ethofs-cli -list -key={PrivateKey}
```

## Running pre-built binary (standard upload)

```shell
./build/bin/ethofs-cli -upload -path=test.txt -key={PrivateKey}
```

## Running pre-built binary (recursive directory upload)

```shell
./build/bin/ethofs-cli -r -upload -path=testDir -key={PrivateKey}
```

## Running pre-built binary (remove upload)

```shell
./build/bin/ethofs-cli -remove -contractaddress={HostingContractAddress} -key={PrivateKey}
```

## Running pre-built binary (upload contract extension)

```shell
./build/bin/ethofs-cli -extend -blocks={ExtensionBlockCount} -contractaddress={HostingContractAddress} -key={PrivateKey}
```

![Upload Example](ethofs-cli.png)
