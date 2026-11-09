
package handlers

import (
	"net"
	"fmt"
	"log" 
	"bufio"
	"github.com/knutekje/lloggo/connectionmanager"
)

func HandleConnection(cm *connectionmanager.ConnectionManager, conn net.Conn) {
	key := cm.Add(conn)
	fmt.Println(key)

	defer func() {
		conn.Close()
		cm.Remove(key)
		fmt.Printf("Connection removed: %s\n", key)
	}()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		cm.Broadcast(buffer)
		fmt.Printf("[%s]: %s\n", key, string(buffer[:n]))
	}


}
