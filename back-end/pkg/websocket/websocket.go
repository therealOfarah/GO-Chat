package websocket

import (
    "fmt"
    "io"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}
type Pool struct{
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}
type Client struct{
	ID string
	Conn *websocket.Conn
	Pool *Pool
}
type Message struct{
	Type int
	Body string
}
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}
func (pool *Pool)Start(){
	for {
		select {
		case client := <-pool.Register:
				pool.Clients[client] = true
				fmt.Println("Size of Connection Pool: ", len(pool.Clients))
				for client, _ := range pool.Clients {
						fmt.Println(client)
						client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
						break
				}
		case client := <-pool.Unregister:
				delete(pool.Clients, client)
				fmt.Println("Size of Connection Pool: ", len(pool.Clients))
				for client, _ := range pool.Clients {
						client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
						break
				}
		case message := <-pool.Broadcast:
				fmt.Println("Sending message to all clients in Pool")
				for client, _ := range pool.Clients {
						if err := client.Conn.WriteJSON(message); err != nil {
								fmt.Println(err)
								return
						}
				}	
		}
	}
}
func (c *Client)Read(){
	defer func(){
		//passes a value from a channel to a refrence
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for{
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    return conn, nil
}

func Reader(conn *websocket.Conn) {
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}

func Writer(conn *websocket.Conn) {
    for {
        fmt.Println("Sending")
        messageType, r, err := conn.NextReader()
        if err != nil {
            fmt.Println(err)
            return
        }
        w, err := conn.NextWriter(messageType)
        if err != nil {
            fmt.Println(err)
            return
        }
        if _, err := io.Copy(w, r); err != nil {
            fmt.Println(err)
            return
        }
        if err := w.Close(); err != nil {
            fmt.Println(err)
            return
        }
    }
}