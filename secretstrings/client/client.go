package main

import (
	"flag"
	"net/rpc"

	"bufio"
	"fmt"
	"os"
	"secretstrings/stubs"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	stdin := bufio.NewReader(os.Stdin)
	flag.Parse()
	fmt.Println("Server: ", *server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()
	text, _ := stdin.ReadString('\n')
	request := stubs.Request{Message: text}
	response := new(stubs.Response)
	client.Call(stubs.PremiumReverseHandler, request, response)
	fmt.Println("Responded: " + response.Message)

	//TODO: connect to the RPC server and send the request(s)
}
