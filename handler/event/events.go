package event

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024 * 1024 * 1024,
		WriteBufferSize: 1024 * 1024 * 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Client struct {
	Id   string
	conn *websocket.Conn
	Send chan []byte
}

type Event struct {
	Timestamp string
	Topic     string
	Value     string
}

func EventsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	eventChan := make(chan Event)
	closeChan := make(chan bool)
	msgChan := make(chan []byte)

	client := Client{
		Id:   uuid.Must(uuid.New(), nil).String(),
		conn: conn,
		Send: msgChan,
	}

	go GenerateEvents(eventChan)                    // generate our random events
	go ManageConns(client)                          // manages our client connections
	go ConsumeEvents(eventChan, closeChan, msgChan) // consumes our events and writes to our message chan

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			_ = client.conn.Close()
			break
		}
		client.Send <- message
	}

}

// a function to write into some events channel
func GenerateEvents(eventChan chan Event) {
	for {
		eventChan <- Event{
			Timestamp: time.Now().String(),
			Topic:     "org1/berlin/loc1/mac001/temp1",
			Value:     fmt.Sprintf("%d C", rand.Intn(100)),
		}
		time.Sleep(2 * time.Second)
	}
}

// a function to consume events from eventchannel
func ConsumeEvents(eventChan chan Event, closeChan chan bool, msgChan chan []byte) {
	for {
		event := <-eventChan
		select {
		case <-closeChan:
			return
		default:
			jsonbytes, _ := json.Marshal(event)
			msgChan <- jsonbytes
		}
	}
}

// a function to manage connections
func ManageConns(client Client) {
	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				if err := client.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					fmt.Println("we are not able to close the connection")
				}
				return
			}
			if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Println("we are unable to write message into the channel")
			}
		default:
			continue
		}
	}
}
