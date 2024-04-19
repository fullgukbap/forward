package netcode

import (
	"fmt"
	"net"
)

const (
	maxMsgSize = 512
)

type Client struct {
	Hub *Hub

	Conn net.Conn

	Send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	buf := make([]byte, maxMsgSize)
	for {
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("%s Client의 메시지 수신을 완료하지 못했습니다: %v", c.Conn.RemoteAddr(), err)
			continue
		}
		c.Hub.Input <- buf
	}
}

func (c *Client) writePump() {
	defer c.Conn.Close()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				return
			}

			_, err := c.Conn.Write(msg)
			if err != nil {
				fmt.Printf("%s Client에게 메시지를 송신을 완료하지 못했습니다: %v", c.Conn.RemoteAddr(), err)
				continue
			}
		}
	}
}

func (c *Client) Init() {
	go c.readPump()
	go c.writePump()
}
