package main

import (
	"fmt"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func main() {

}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " " + e.Err.Error()
}
