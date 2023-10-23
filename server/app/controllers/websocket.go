package controllers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

type WsMessage struct {
	Action   string `json:"action,omitempty"`
	Button   string `json:"button,omitempty"`
	ClientId string `json:"c,omitempty"`
	RobotId  string `json:"r,omitempty"`
	Base64   string `json:"data,omitempty"`
}

type Data struct {
	Base64 string `json:"data,omitempty"`
}

var connections = make(map[string]*websocket.Conn)
var robots = make(map[string]string)
var users = make(map[string]string)

func Websocket(conn *websocket.Conn) {
	log.Println(conn.Locals("Host"))

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", msg)

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
				roboClient.WriteMessage(mt, msg)
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

		//err = conn.WriteMessage(mt, msg)
		//
		//if err != nil {
		//	log.Println("write:", err)
		//	break
		//}
	}
}
