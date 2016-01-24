package handler

import(
		"net"
		"fmt"
		"bufio"
		"io/ioutil"
		"encoding/json"
	)
type user struct {
	name string
	password string //actually the password hash
}
type miner struct {
	address string
	prefix string
	user string
	shares int
}
var miners [1000]miner //I can expand this if needed
var users [1000]user //I can expand this if needed
func writeLists() { //FIXME: This exports it to an infinite long string.(See miners.json)
	var minerArray [len(miners)][4]string
	var userArray [len(users)][2]string
	for i := 0; i < len(miners); i++ {
		minerArray[i][0] = miners[i].address
		minerArray[i][1] = miners[i].prefix
		minerArray[i][2] = miners[i].user
		minerArray[i][3] = string(miners[i].shares)
	}
	a,_ := json.Marshal(minerArray);
	bytes := []byte(string(a))
	err := ioutil.WriteFile("miners.json",bytes,0644)
	if(err != nil) {
		fmt.Println("Error while saving miners!")
	}
	for i := 0; i < len(users); i++ {
		userArray[i][0] = users[i].name
		userArray[i][1] = users[i].password
	}
	a,_ = json.Marshal(userArray)
	bytes = []byte(string(a))
	err = ioutil.WriteFile("users.json",bytes,0644)
	if(err != nil){
		fmt.Println("Error while saving users")
	}
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
	miners[0] = miner{"klucaund08","7s","Luca",100}
	users[0] = user{"Luca","someshastring"}
	writeLists()
	s,err:= net.Listen("tcp",":8392") //8392 is the port the miner will connect on (Don't ask)
	if(err != nil){
		fmt.Println(err)
	}
	for {
		conn, _ := s.Accept()
		go handleConn(conn)
	}
}