package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	target = "192.168.6.211:50050"
)

func main() {
	// opening client connection

	cc, err := grpc.Dial(target,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(&DoNotCodec{})),
	)
	if err != nil {
		log.Fatalf("failed to dial client connection: %v", err)
	}

	// request definition

	msg := []byte{10, 4, 89, 117, 114, 105} // HelloRequest with name "Yuri" encoded here
	mdForInvoke := metadata.Pairs("the-way-name", "Invoke")
	mdForStream := metadata.Pairs("the-way-name", "ClientStream")

	// the way #1 Invoke

	invokeResult, headerInv, trailerInv, err := invoke(cc, mdForInvoke, "/helloworld.Greeter/SayHello", msg)
	log.Printf("'Invoke' result: %v", invokeResult)
	log.Printf("'Invoke' metadata\n\theader:  %v\n\ttrailer: %v", headerInv, trailerInv)
	if err != nil {
		log.Printf("'Invoke' error: %v", err)
	}

	// the way #2 ClientStream

	streamResult, headerStr, trailerStr, err := streamCall(cc, mdForStream, "/helloworld.Greeter/SayHello", msg)
	log.Printf("'Stream' result: %v", streamResult)
	log.Printf("'Stream' metadata\n\theader:  %v\n\ttrailer: %v", headerStr, trailerStr)
	if err != nil {
		log.Printf("'Stream' error: %v", err)
	}

	// tear down
	cc.Close()
}

func invoke(cc *grpc.ClientConn, md metadata.MD, method string, input []byte) (result []byte, header metadata.MD, trailer metadata.MD, err error) {
	err = cc.Invoke(metadata.NewOutgoingContext(context.Background(), md), method, &input, &result,
		grpc.Header(&header), grpc.Trailer(&trailer))
	return
}

func streamCall(cc *grpc.ClientConn, md metadata.MD, method string, input []byte) (result []byte, header metadata.MD, trailer metadata.MD, err error) {
	clientStream, err := cc.NewStream(metadata.NewOutgoingContext(context.Background(), md),
		&grpc.StreamDesc{ ClientStreams: true },
		"/helloworld.Greeter/SayHello")
	if err != nil {
		return
	}

	clientStream.SendMsg(&input)
	clientStream.RecvMsg(&result)

	header, err = clientStream.Header()
	trailer = clientStream.Trailer()
	return
}
