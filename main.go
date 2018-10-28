package main

import (
	"flag"
	"log"
)

func main() {

	var start uint
	var end uint
	flag.UintVar(&start, "start", 0, "usage")
	flag.UintVar(&end, "end", 10, "usage")

	flag.Parse()
	//log.Println(start, end)
	log.Println(start)
	log.Println(end)
}
