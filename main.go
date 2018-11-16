package main

import "fmt"
import "github.com/twinj/uuid"

func main() {
	doStuff()
}
func doStuff() {
	fmt.Println("UUID: " + uuid.NewV4().String())
}

func Dump(obj Stringer) {
	fmt.Printf("%v\n", obj.ToString())
}

type Stringer interface {
	ToString() string
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
