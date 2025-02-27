/*
Command example-grpc-server is an example grpc server
to be called by example-gateway-server.
*/
package main

import (
	"context"
	"flag"

	"github.com/3128px/grpc-gateway/v2/examples/internal/server"
	"github.com/golang/glog"
)

var (
	addr    = flag.String("addr", ":9090", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	if err := server.Run(ctx, *network, *addr); err != nil {
		glog.Fatal(err)
	}
}
