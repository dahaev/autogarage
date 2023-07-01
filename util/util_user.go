package util

import (
	"fmt"
)

func RandomEmailUser() string {
	return fmt.Sprintf("%s@gmail.com", randomString(7))
}

func RandomHashPassword() string {
	return fmt.Sprintf("secret")
}

func RandomUserString(n int) string {
	return fmt.Sprintf("%s", randomString(n))
}
