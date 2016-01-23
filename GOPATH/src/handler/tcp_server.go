package handler

import(
		"net"
		"fmt"
		"bufio"
		//"encoding/json"
	)
type miner struct {
	address string
	prefix string
	shares int
	balance int
}
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

/*
Some code to convert miners into a json array(and maybe back)
	var jsonArray [len(miners)][4]string
	for i := 0; i < len(miners); i++ {
		jsonArray[i][0] = miners[i].address
		jsonArray[i][1] = string(miners[i].solved)
	}
	a,_ := json.Marshal(jsonArray);
	fmt.Println(string(a))
*/