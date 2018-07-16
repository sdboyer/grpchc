## grpchc

The simplest possible implementation of [standard gRPC healthchecking](https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

## Installation

```
go get -u github.com/sdboyer/grpchc
```

## Usage

Point it at the gRPC service you want to healthcheck:

```
grpchc localhost:3000
```

grpchc will attempt to connect and perform a single healthcheck query. If the service responds with anything other than SERVING (1), or does not respond within a timeout of 2 seconds, grpchc will exit 1.

You can also optionally specify a service name:

```
grpchc -svcname package_names.ServiceName localhost:3000
```

NOTE: TLS connections are not currently supported.
