// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.

type boardcast struct {
	channel string
	msg     []byte
}

// Hub type
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan boardcast

	// Register requests from the clients.
	register chan *Client

	// UnRegister requests from clients.
	unRegister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan boardcast),
		register:   make(chan *Client),
		unRegister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			fmt.Println("register")
			h.clients[client] = true
		case client := <-h.unRegister:
			fmt.Println("unRegister")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Println(string(message.msg))
			for client := range h.clients {
				if message.channel == client.channel {
					fmt.Println(message.channel, "::", string(message.msg))
					select {
					case client.send <- string(message.msg):
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
