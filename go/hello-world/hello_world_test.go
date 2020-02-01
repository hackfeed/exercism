package greeting

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if observed := HelloWorld(); observed != expected {
		t.Fatalf("HelloWorld() = %v, want %v", observed, expected)
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld()
	}
}
