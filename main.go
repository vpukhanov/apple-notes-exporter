package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vpukhanov/notes-exporter/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	r := new(float32)
	*r = 0.5
	c := &pb.Color{
		Red:   r,
		Green: r,
		Blue:  r,
		Alpha: r,
	}

	fname := "test.pb"
	out, err := proto.Marshal(c)
	if err != nil {
		log.Fatalln("Failed to encode color:", err)
	}
	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write color:", err)
	}

	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	color := &pb.Color{}
	if err := proto.Unmarshal(in, color); err != nil {
		log.Fatalln("Failed to parse color:", err)
	}

	fmt.Printf("%+v\n", color)
}
