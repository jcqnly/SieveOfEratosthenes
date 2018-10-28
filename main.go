package main

import "flag"

func main() {

	var start uint
	var end uint
	flag.UintVar(&start, "start", 0, "usage")
}
