
package handlers

import (
	"net"
	"fmt"
	"github.com/knutekje/lloggo/managers"
	"encoding/binary"
	"io"
)


const (
	CmdSubscribe byte = 0x01	
	CmdPublish byte = 0x02
)




func HandleConnection(cm *managers.ConnectionManager, conn net.Conn) {
	key := cm.Add(conn)
	fmt.Println(key)

	defer func() {
		conn.Close()
		cm.Remove(key)
		fmt.Printf("Connection removed: %s\n", key)
	}()

	lenBuf := make([]byte, 2)

	for {
		if _, err := io.ReadFull(conn, lenBuf); err != nil {
			return
		}
		msgLen := binary.BigEndian.Uint16(lenBuf)
		if msgLen == 0 {
			continue
		}

		payload := make([]byte, msgLen)
		if _, err := io.ReadFull(conn, payload); err != nil {
			return
		}
		sig := payload[0]
		data := payload[1:]

		switch sig {
		case CmdSubscribe:

			fmt.Printf("YO - [%s] subscribed to: %s\n", key, string(data))

		case CmdPublish:

			cm.Broadcast(data)
			fmt.Printf("[%s] Channel %d: %s\n", key, sig, string(data))
		}
	}


}
