
package handlers

import (
	"net"
	"log" 
	"bufio"
)
func HandleMessage(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	bytes, err := reader.ReadBytes(10)

	if err != nil {
		log.Println("ERROR",err)
		return
	}


	message := string(bytes)

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("ERROR WRITING:", err)
		return
	}



	log.Println(message)

	log.Println(bytes)



}
