package ws

import (
	"encoding/json"
	"github.com/Defenestrationq/forum-api/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

func (hub *Hub) Chat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, []byte("Invalid thread ID"))
		conn.Close()
		return
	}

	client := &Client{
		conn:  conn,
		send:  make(chan entity.Message, 256),
		disID: id,
	}

	hub.register <- client

	go func() {
		posts, err := hub.msgUseCase.GetAllMessage(id)
		if err != nil {
			return
		}

		for _, post := range posts {
			select {
			case client.send <- post:
			default:
				hub.unregister <- client
				conn.Close()
				return
			}
		}
	}()

	go func() {
		defer func() {
			hub.unregister <- client
			conn.Close()
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				break
			}

			var msg entity.Message
			if err := json.Unmarshal(message, &msg); err != nil {
				conn.WriteJSON(map[string]string{"error": "invalid message format"})
				continue
			}

			msg.DiscussionID = id
			err = hub.msgUseCase.CreateMessage(msg)
			if err != nil {
				conn.WriteJSON(map[string]string{"error": "failed to create msg"})
				continue
			}
			hub.chat <- msg
		}
	}()

	go func() {
		defer conn.Close()

		for message := range client.send {
			postBytes, err := json.Marshal(message)
			if err != nil {
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, postBytes); err != nil {
				break
			}
		}
	}()
}
