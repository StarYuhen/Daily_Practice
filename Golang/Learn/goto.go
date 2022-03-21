package main

import "log"

func main() {
P:
	log.Println("sdfdsgdfg")

LOOP:
	log.Println("dgdfsgdfh")
	goto P

	goto LOOP
}
