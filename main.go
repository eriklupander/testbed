package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
}

func main() {
	fmt.Println("Starting testbed application")
	Connect()
}

type T1 struct {
	val  string
	size int
}

func (t T1) Equal(t2 T1) bool {
	return t.size == t2.size
}

func (t T1) ToString() string {
	return t.val
}

type T2 struct {
	val  string
	size int
}

func (t T2) ToString() string {
	return t.val
}
