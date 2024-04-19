package netcode

import "errors"

// Hub 구조체는 Conn구조체와 Engine 구조체의 중간 매개체 입니다.
// 매개체의 역활로 Conn의 연결 및 누락을 처리합니다.
// 또한 게임 로직 이벤트를 발생시켜 로직 처리를 유도합니다.
type Hub struct {
	Engine *Engine

	Clients map[*Client]bool

	Input chan []byte

	Register chan *Client

	Unregister chan *Client
}

func NewHub(e *Engine) *Hub {
	if e == nil {
		panic("the value of e is nil")
	}
	return &Hub{
		Engine:     e,
		Clients:    make(map[*Client]bool),
		Input:      make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Activating() error {
	for {
		select {

		case client := <-h.Register:
			if err := h.handleRegister(client); err != nil {
				return err
			}

		case client := <-h.Unregister:
			if err := h.handleUnregister(client); err != nil {
				return err
			}
		}

		case 

	}
}

func (h *Hub) handleRegister(c *Client) error {
	h.Clients[c] = true
	return nil
}

func (h *Hub) handleUnregister(c *Client) error {
	if _, ok := h.Clients[c]; ok {
		delete(h.Clients, c)
		close(c.Send)
		return nil
	}
	return errors.New("the client doesn't exist")
}

func handleInput(c *Client, msg []byte) error {

}
