package main

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (mr MyReader) Read(b *[]byte) (int, error) {
	return 0, nil
}

func main() {
	reader.Validate(MyReader{})
}
