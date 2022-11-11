package main

import (
	"fmt"

	"github.com/eatos118/demo001/helloworld/proto/pb"
)

func main() {
	p := pb.SimpleRequest{
		PageNumber: 1,
	}
	fmt.Print(p.String())
}
