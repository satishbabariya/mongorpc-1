package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mongorpc/mongorpc"
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewMongoRPCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	r, err := c.ListCollections(ctx, &proto.ListCollectionsRequest{
		Database: "sample_mflix",
	})
	if err != nil {
		logrus.Fatalf("could not get collection: %v", err)
	}
	json, err := json.Marshal(mongorpc.ValueToArray(r.Collections))
	if err != nil {
		logrus.Fatalf("could not Marshal response: %v", err)
	}
	logrus.Printf("Collection: %s", json)
}
