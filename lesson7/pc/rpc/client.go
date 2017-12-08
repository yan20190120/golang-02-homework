package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
	"fmt"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply string
		err = client.Call("Listener.GetLine", line, &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	}
}
