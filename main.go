package main

import (
	"fmt"
	"github.com/potatowhite/web/study04/lzw"
)

// Component
type Component interface {
	Operator(string)
}

var sentData string

type SendComponent struct {
}

func (self *SendComponent) Operator(data string) {
	// Send Data
	sentData = data
}

// Zip decorator
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipped, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(zipped))
}

func main() {
	fmt.Print("Hello World")
}
