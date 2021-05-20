package client

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Client julius client structure
type Client struct {
	conn      net.Conn
	connClose func()
	channel   chan string
}

// NewClient create new Client
func NewClient(network, address string) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
		connClose: func() {
			if err := conn.Close(); err != nil {
				log.Fatal(err)
			}
		},
		channel: make(chan string),
	}, nil
}

// receive Receiving message and send to channel
func (c *Client) receive(_ context.Context) {
	for {
		buf := make([]byte, 4096)
		n, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		c.channel <- string(buf[:n])
	}
}

// handle Handling for received messages
func (c *Client) handle(ctx context.Context) {
	for msg := range c.channel {
		if !isRecogout(msg) {
			continue
		}
		result, err := parseMessage(msg)
		if err != nil {
			log.Println(err)
		}
		for _, handler := range RegisteredHandlers {
			if handler.Filter(ctx, result) {
				if err := handler.Do(ctx, result); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

// Run starting connection
func (c *Client) Run() {

	ctx, cancel := context.WithCancel(context.Background())

	go c.handle(ctx)
	go c.receive(ctx)
	go c.close(ctx)

	// Wait until the force-quit (such as ctrl + C) command comes
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit
	cancel()
}

// close terminate the connection
func (c *Client) close(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			if err := c.conn.Close(); err != nil {
				fmt.Printf("failed to close connection: %+v", err)
			}
			return
		}
	}
}
