package main
import (
	"net"
	"log"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", ":6020")
	if err != nil {
		log.Println("FAILED TO START", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("NOPE", err)
			continue
		}
		go handleMessage(conn)
	}
}

func handleMessage(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Println("ERROR",err)
		return
	}
	log.Println(message)


}
