package blind

// NotCodec is a codec that does nothing, just leaves the data as is
// It implements encoding.Codec interface (google.golang.org/grpc/encoding)
type NotCodec struct{}

// Marshal returns its argument as []byte
func (c *NotCodec) Marshal(v interface{}) ([]byte, error) {
	return *v.(*[]byte), nil
}

// Unmarshal assigns the given data to the value
func (c *NotCodec) Unmarshal(data []byte, v interface{}) error {
	*v.(*[]byte) = data
	return nil
}

// Name returns the name of the Codec implementation. The returned string
// is not expected to be used as part of content type in transmission.
func (c *NotCodec) Name() string {
	return "bytes"
}
