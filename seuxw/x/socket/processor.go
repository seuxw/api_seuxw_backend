package socket

import (
	"seuxw/x/socket/websocket"
)

type Processor interface {
	OnConnected(*Connection)
	OnMessage(*Message)
	OnDisconnected(*Connection)
}

type Message struct {
	SN      int    // server NO
	ID      int    // message id
	RawData []byte // message data
	conn    *websocket.Conn
}

type Connection websocket.Conn

type Authenticator interface {
	Auth(id, cer string) error
}

type defaultAuthenticator struct{}

func (a *defaultAuthenticator) Auth(id, cer string) error {
	return nil
}
