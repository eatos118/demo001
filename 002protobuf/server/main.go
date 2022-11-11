package main

import (
	"fmt"

	"github.com/eatos118/demo001/002protobuf/proto/pb"
)

func main() {
	p := pb.SimpleRequest{
		PageNumber: 1,
	}
	fmt.Print(p.String())
}
