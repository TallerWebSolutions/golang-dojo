package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	result := helloWorld()
	if result != "Hello, World!" {
		t.Error("Hello World failed")
	}

}
