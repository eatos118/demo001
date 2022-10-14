package main

import (
	"fmt"

	pb "github.com/eatos118/demo001/helloworld/helloworldpb"
)

func main() {
	p := pb.SimpleRequest{
		PageNumber: 1,
	}
	fmt.Print(p.String())
}
