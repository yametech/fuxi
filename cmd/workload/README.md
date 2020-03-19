# Workload Service

This is the Workload service

Generated with

```
micro new github.com/yametech/fuxi/cmd/workload --namespace=go.micro --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.workload
- Type: api
- Alias: workload

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
./workload-api
```

Build a docker image
```
make docker
```