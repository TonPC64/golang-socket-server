// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var count = 0

func main() {
	hub := newHub()
	go hub.run()

	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/:channel", func(c echo.Context) error {
		channel := c.Param("channel")
		serveWs(hub, c.Response(), c.Request(), channel)
		return nil
	})

	e.GET("/publish/:channel/:data", func(c echo.Context) error {
		channel := c.Param("channel")
		data := c.Param("data")
		hub.broadcast <- boardcast{
			channel: channel,
			msg:     []byte(data),
		}
		count++
		fmt.Printf("count %d || Channel: %s || Data: %s\n", count, channel, data)
		return c.String(200, "ok")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
