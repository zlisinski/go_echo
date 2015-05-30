# go_echo
A simple echo server written in Go

To build, run:

    golang-go build

In one terminal run ```./go_echo``` to have the server listen on the default port 8000. To change the port, use the ```-p``` flag, for example:

    ./go_echo -p 3000
    
  or 
  
    ./go_echo -p=3000
    
to listen on port 3000.

In another terminal, run

    telnet localhost 8000

Replace the ```8000``` with whatever port you used for the server.

Type whatever you want, and it will be echoed back to you along with a timestamp of when it was received by the server.
