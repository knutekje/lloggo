package connectionmanager

import (
	"sync"
	"net"
)

type ConnectionManager struct {
	mu sync.RWMutex 
	connections map[string]net.Conn
	subscriptions map[string][]string
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[string]net.Conn),
		subscriptions: make(map[string][]string),
	
	}
}

func (cm *ConnectionManager) Add(conn net.Conn) string {
	key := conn.RemoteAddr().String()
	cm.mu.Lock()
	cm.connections[key] = conn
	cm.mu.Unlock()
	return key
}

func (cm *ConnectionManager) Remove(key string) {
	
	cm.mu.Lock()
	delete(cm.connections, key)
	cm.mu.Unlock()
}

func (cm *ConnectionManager) Subscribe(channel string) {
	cm.mu.Lock()

	cm.mu.Unlock()
}

func (cm *ConnectionManager) Broadcast(message []byte) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	for _, conn := range cm.connections {
		go conn.Write(message)
	}
}
