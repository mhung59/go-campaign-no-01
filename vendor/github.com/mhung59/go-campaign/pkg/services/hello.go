package services

import (
	f "fmt"
)

func Hello(name string) string {
	msg := f.Sprintf("Hello, %v", name)

	return msg
}
