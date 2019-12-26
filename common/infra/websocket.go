package infra

import (
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"go-socket/modules/message/dto"
	"net/http"
	"time"
)

type M map[string]interface{}

var CurrentConnection WebSocketConnection

var connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	ID        uuid.UUID
	Message   string
	CreatedAt time.Time
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Handler WebSocket
func Wshandler(w http.ResponseWriter, r *http.Request) {
	currentGorillaConn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	username := r.URL.Query().Get("username")
	currentConn := WebSocketConnection{Conn: currentGorillaConn, Username: username}
	CurrentConnection = currentConn
	connections = append(connections, &currentConn)
}

// BrodacastMessage to send data to client
func BrodacastMessage(currentConn *WebSocketConnection, message *dto.Message) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		err := eachConn.WriteJSON(SocketResponse{
			ID:        message.ID,
			Message:   message.Message,
			CreatedAt: message.CreatedAt,
		})

		if err != nil {
			break
		}
	}
}

// GetCurrentConnection get current connection detail
func GetCurrentConnection() *WebSocketConnection {
	return &CurrentConnection
}
