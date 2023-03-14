package users

import (
	"testing"
)

func TestHashing(t *testing.T) {
	username, password := "username", "46231s"
	generated_hash, err := EncryptPassword(username, password)

	if err != nil {
		t.Fatal(err)
	}
	res, err := VerifyPassword(username, password, string(generated_hash))
	if err != nil || res == false {
		t.Fatal("not working : ", err)
	}
}
