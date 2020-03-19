# Ns Service

This is the Ns service

Generated with

```
micro new github.com/yametech/fuxi/cmd/ns --namespace=go.micro --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.ns
- Type: api
- Alias: ns

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./ns-api
```

Build a docker image
```
make docker
```