package blind

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ClientConn represents a client connection to an RPC server.
type ClientConn struct {
	// C is the underlying grpc-client-connection.
	C *grpc.ClientConn
}

// Dial creates a client connection to the given target.
func Dial (target string, opts ...grpc.DialOption) (*ClientConn, error) {
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.ForceCodec(&NotCodec{})))
	cc, err := grpc.Dial(target, opts...)
	return &ClientConn{cc}, err
}

// DialContext creates a client connection to the given target.
// For more details see https://godoc.org/google.golang.org/grpc#DialContext
func DialContext (ctx context.Context, target string, opts ...grpc.DialOption) (*ClientConn, error) {
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.ForceCodec(&NotCodec{})))
	cc, err := grpc.DialContext(ctx, target, opts...)
	return &ClientConn{cc}, err
}

// Invoke sends the RPC request on the wire and returns after response is
// received.  This is typically called by generated code.
//
// All errors returned by Invoke are compatible with the status package.
func (cc *ClientConn) Invoke(md metadata.MD, method string, input []byte) (result []byte, header metadata.MD, err error) {
	err = cc.C.Invoke(metadata.NewOutgoingContext(context.Background(), md),
		method, &input, &result, grpc.Header(&header))
	return
}

// InvokeWithTrailer does all the Invoke method does but additionally returns response trailers
func (cc *ClientConn) InvokeWithTrailer(md metadata.MD, method string, input []byte) (result []byte, header metadata.MD, trailer metadata.MD, err error) {
	err = cc.C.Invoke(metadata.NewOutgoingContext(context.Background(), md), method, &input, &result,
		grpc.Header(&header), grpc.Trailer(&trailer))
	return
}

// Close tears down the ClientConn and all underlying connections.
func (cc *ClientConn) Close() error {
	return cc.C.Close()
}
