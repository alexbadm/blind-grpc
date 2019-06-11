package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/alexbadm/blind-grpc"
)

const (
	target = "192.168.6.211:50050"
)

func main() {
	cc, _ := blind.Dial(target, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("failed to dial client connection: %v", err)
	// }

	msg := []byte{10, 4, 89, 117, 114, 105} // HelloRequest with name "Yuri" encoded here
	md := metadata.Pairs("md-key", "some value")

	result, header, err := cc.Invoke(md, "/helloworld.Greeter/SayHello", msg)
	log.Printf("\n\tresult: %v\n\theader: %v", result, header)
	if err != nil {
		log.Printf("error: %v", err)
	}

	cc.Close()
}
