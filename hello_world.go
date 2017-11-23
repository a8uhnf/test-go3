package test_go3

import (
	"fmt"
	a"github.com/a8uhnf/test-go2"
)

func HelloWorld()  {
	fmt.Println("Hello World from test-go3")
	a.HelloWorldMaster()
}

func HelloWorldTestDep()  {
	fmt.Println("Hello World from test-go2. Branch: test-dep")
	a.HelloWorldMaster()
}
