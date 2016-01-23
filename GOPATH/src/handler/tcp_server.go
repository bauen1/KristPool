package handler

import(
		"net"
		"fmt"
		"bufio"
	)
func readFromConn(c net.Conn) string{
	message,_ := bufio.NewReader(c).ReadString('\n')
	return string(message)
}
func writeToConn(c net.Conn, s string){
	c.Write([]byte(s + "\n"))
}
func handleConn(c net.Conn) {
	for {
		fmt.Println(readFromConn(c))
	}
}

func HandleTCP(){
	s,err:= net.Listen("tcp",":8392") //8392 is the port the miner will connect on (Don't ask)
	if(err != nil){
		fmt.Println(err)
	}
	for {
		conn, _ := s.Accept()
		go handleConn(conn)
	}
}