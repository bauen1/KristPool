package handler

import(
		"net"
		"fmt"
	)
func handleConn(c net.Conn) {
	fmt.Println(c.RemoteAddr())
	c.Close()
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