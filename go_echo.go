package main

import (
	"fmt"
	"time"
	"os"
	"flag"
	"net"
	"strconv"
	"bufio"
)

const defaultPort int = 8000
var listenPort int

// Parse command line arguments to get listening port
func parseArgs() {
	flag.IntVar(&listenPort, "p", defaultPort, "TCP port to listen on.")
	flag.Parse()
}

func main() {
	// Parse command line arguments to get listening port
	parseArgs()
	
	fmt.Printf("Listening on port %d\n", listenPort)
	
	// Listen on all interfaces on the specified port
	listener, err := net.Listen("tcp", ":" + strconv.Itoa(listenPort))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	
	// Accept connection
	connection, err := listener.Accept()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	
	// Loop forever reading from connection
	for {
		// Read string from connection
		message, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		
		// Output message to console
		fmt.Println(message)
		
		// Echo message back to client
		curTime := time.Now().String()
		newMessage := "Received on " + curTime + ":\n" + message + "\n"
		connection.Write([]byte(newMessage))
	}
}