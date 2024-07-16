package handlers

import (
    "fmt"
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
)

type WebSocketHandler struct {
    Upgrader websocket.Upgrader
}

func (h *WebSocketHandler) ServeHTTP(c echo.Context) error {
    conn, err := h.Upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer conn.Close()

    for {
        // Read message from browser
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            return err
        }

        // Print the message to the console
        fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

        // Write message back to browser
        if err = conn.WriteMessage(msgType, msg); err != nil {
            return err
        }
    }
}
