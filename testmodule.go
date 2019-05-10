package testmodule

import "fmt"

func TestHi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}