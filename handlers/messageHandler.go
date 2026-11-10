
package handlers

import (
	"net"
	"fmt"
	"log"
	"github.com/knutekje/lloggo/managers"
	"strings"
)

func HandleConnection(cm *managers.ConnectionManager, conn net.Conn) {
	key := cm.Add(conn)
	fmt.Println(key)

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("FAILED TO CLOSE")
		}
		cm.Remove(key)
		fmt.Printf("Connection removed: %s\n", key)
	}()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		word := strings.TrimSpace(string(buffer[:n])) 
		if word == "subscribe" {
			fmt.Println("YO")
		}
		cm.Broadcast(buffer)
		fmt.Printf("[%s]: %s\n", key, string(buffer[:n]))
	}


}
