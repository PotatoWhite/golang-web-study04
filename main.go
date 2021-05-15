package main

import (
	"fmt"
	"github.com/potatowhite/web/study04/cipher"
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

var recvData string

type RecvComponent struct {
}

func (self *RecvComponent) Operator(data string) {
	recvData = data
}

// Zip decorator
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	read, err := lzw.Read([]byte( data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(read))
}

type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	cipherText, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(cipherText))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decrypted, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decrypted))
}

func main() {

	senderChain := &EncryptComponent{key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	senderChain.Operator("Hello World")
	fmt.Println(sentData)

	recvChain := &UnzipComponent{
		com: &DecryptComponent{key: "abcde",
			com: &RecvComponent{},
		},
	}

	recvChain.Operator(sentData)
	fmt.Println(recvData)
}
