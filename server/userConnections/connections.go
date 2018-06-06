package userConnections

import (
	"github.com/gorilla/websocket"
	"net"
	"sync"
)

type Connections struct{
	WSConnectionsMutex *sync.Mutex
	WSConnections map[*websocket.Conn]string // connection:login
	TCPConnectionsMutex *sync.Mutex
	TCPConnections map[net.Conn]string // connection:login
	OutChan chan *Message
}

func (c *Connections) AddTCPConn(conn net.Conn, userName string, outChan *Message) Connections{
	newStr := Connections{}
	newStr.TCPConnections[conn] = userName
	newStr.OutChan <- outChan
	return newStr
}

func (c *Connections) GetAllTCPConnections() map[net.Conn]string {
	c.TCPConnectionsMutex.Lock()
	defer c.TCPConnectionsMutex.Unlock()
	return c.TCPConnections
}

func (c *Connections) GetAllWSConnections() map[*websocket.Conn]string {
	c.WSConnectionsMutex.Lock()
	defer c.WSConnectionsMutex.Unlock()
	return c.WSConnections
}