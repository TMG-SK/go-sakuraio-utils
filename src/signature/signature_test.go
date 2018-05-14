package sakuraio

import (
	"testing"
)

const secret string = "secret"

func TestCheckSignature(t *testing.T) {
	result := CheckSignature(secret, []byte("hello"), "5112055c05f944f85755efc5cd8970e194e9f45b")

	if result == false {
		t.Fatal("Signature not match")
	}
}
