package authn

import (
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	b, _ := Encrypt("admin!#999")
	fmt.Println(string(b))

	s := "$2a$10$eXgOsRqZq8YPYLKt5.YFOuMpjvs4Y0pF7d83/U3r6RmNcAoz65732"

	err := Compare(s, "admin!#999")
	fmt.Println(err)
}
