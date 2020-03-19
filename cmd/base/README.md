# 基础服务


```
micro new github.com/yametech/fuxi/api/base --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.base
- Type: srv
- Alias: base

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

## 当前使用使用的是ECTD,需要安装ECTD。

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./base-srv
```

Build a docker image
```
make docker
```
