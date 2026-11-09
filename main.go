package main
import (
	"net"
	"log"
	"github.com/knutekje/lloggo/handlers"
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
		go handlers.HandleMessage(conn)

	}
}

