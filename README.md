# blind-grpc

An example of calling grpc-servers without knowing protobuf schema.

This repo demonstrates how anyone can invoke grpc-server with preliminary encoded messages as input.

## Usage

```go
// import "github.com/alexbadm/blind-grpc"

cc, _ := blind.Dial("some-server:port", grpc.WithInsecure())
result, header, err := cc.Invoke(nil, "/helloworld.Greeter/SayHello", []byte{10, 4, 89, 117, 114, 105})
// do smth with result, header, err
```
