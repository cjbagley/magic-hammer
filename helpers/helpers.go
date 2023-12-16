package helpers

import (
	"fmt"
	"os"
)

func Exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
