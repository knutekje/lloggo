package main
import (
	"net"
	"log"
	"github.com/knutekje/lloggo/handlers"
	"github.com/knutekje/lloggo/managers"
)

func main() {
	cm := managers.NewConnectionManager()
	listener, err := net.Listen("tcp", ":6020")
	if err != nil {
		log.Println("FAILED TO START", err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Println("failed to close connection")
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("NOPE", err)
			continue
		}
		go handlers.HandleConnection(cm, conn)

	}
}

