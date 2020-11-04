package passwords

import (
	"learning-golang/api.example.com/src/system/passwords"
	"log"
	"testing"
)

// Init function
func Init() {
	log.Println("Testing Password")
}

// TestBasicLog function
func TestBasicLog(t *testing.T) {
	pass := "TEST"
	hash, err := passwords.Encrypt(pass)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println(hash)
	ok := passwords.IsValid(hash, pass)
	if !ok {
		t.Error("Password not matching ... hashing it not working")
	}
	log.Println("Successfully tested hashing!")
}
