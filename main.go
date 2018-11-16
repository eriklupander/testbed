package main

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

func init() {
	viper.AutomaticEnv()
}

func main() {
	fmt.Println("Starting testbed application")
	err := Connect()
	if err != nil {
		//panic(err.Error())
		fmt.Println(err.Error())
	}
	SetupGin()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
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
