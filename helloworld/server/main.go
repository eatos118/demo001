package main

import (
	"fmt"

	pb "github.com/eatos118/demo001/helloworld/helloworldpb"
)

func main() {
	p := pb.Person{
		Name: "test",
	}
	fmt.Print(p)
}
