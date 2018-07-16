// +build go1.9
// Need type aliases from 1.9

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	var name, addr string
	flag.StringVar(&name, "svcname", "", "The name of the gRPC service to healthcheck. Defaults to empty, which gRPC specifies as checking overall server status.")
	flag.Parse()

	args := flag.Args()
	switch len(args) {
	case 0:
		fmt.Fprintf(os.Stderr, "Must specify an address (host:port) to query\n")
		os.Exit(1)
	case 1:
	default:
		fmt.Fprintf(os.Stderr, "Too many args, expected one address (host:port) for querying\n")
		os.Exit(1)
	}
	addr = args[0]

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	cc, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to gRPC server: %s\n", err)
		os.Exit(1)
	}

	cl := grpc_health_v1.NewHealthClient(cc)
	resp, err := cl.Check(ctx, &grpc_health_v1.HealthCheckRequest{Service: name})

	if err != nil {
		fmt.Fprintf(os.Stderr, "error while attempting healthcheck: %s\n", err)
		os.Exit(1)
	}

	switch resp.Status {
	case grpc_health_v1.HealthCheckResponse_UNKNOWN:
		fmt.Fprintf(os.Stderr, "unknown service %q\n", name)
		os.Exit(1)
	case grpc_health_v1.HealthCheckResponse_NOT_SERVING:
		os.Exit(1)
	}
}
