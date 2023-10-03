package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsMessage struct {
	Action   string `json:"action,omitempty"`
	Button   string `json:"button,omitempty"`
	ClientId string `json:"c,omitempty"`
	RobotId  string `json:"r,omitempty"`
}

var connections = make(map[string]*websocket.Conn)
var robots = make(map[string]string)
var users = make(map[string]string)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// r.StaticFile("/favicon.ico", "./resources/img/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "client.html", gin.H{})
	})

	r.GET("/websocket", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	r.Run(":9098")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	defer conn.Close()

	for {
		t, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println("error:", err)
			return
		}

		log.Println(string(msg))

		message := WsMessage{}

		err = json.Unmarshal(msg, &message)

		if err != nil {
			log.Println("error:", err)
			return
		}

		if message.Action != "" && message.Button != "" && message.ClientId != "" {
			log.Println("Move")
			robotId := users[message.ClientId]

			if robotId != "" {
				roboClient := connections[robotId]
				roboClient.WriteMessage(t, msg)
			}

		} else if message.RobotId != "" {
			log.Println("Robot Connected:", message.RobotId)
			connections[message.RobotId] = conn
			robots[message.RobotId] = ""

			for userId, robotId := range users {
				if robotId == "" {
					robots[message.RobotId] = userId
					users[userId] = message.RobotId
					break
				}
			}

		} else if message.ClientId != "" {
			for robotId, userId := range robots {
				log.Printf("userId: %s, robotId: %s", userId, robotId)

				if userId == "" {
					users[message.ClientId] = robotId
					robots[message.RobotId] = message.ClientId
					connections[message.RobotId] = conn
					log.Println("User assigned to the robot")
					break
				} else {
					log.Println("This robot is not available")
				}
			}
		}
	}
}
