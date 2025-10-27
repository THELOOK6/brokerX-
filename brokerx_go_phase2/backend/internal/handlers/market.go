
package handlers

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/gofiber/fiber/v2"
)

func GetQuote(c *fiber.Ctx) error {
	symbol := c.Query("symbol", "AAPL")
	price := 100.0 + float64(rand.Intn(100))/10.0
	return c.JSON(fiber.Map{"symbol": symbol, "price": fmt.Sprintf("%.2f", price), "ts": time.Now().UnixMilli()})
}

func StreamQuotes(c *fiber.Ctx) error {
	symbol := c.Query("symbol", "AAPL")
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	for i:=0; i<1000; i++ {
		price := 100.0 + float64(rand.Intn(100))/10.0
		if _, err := c.Write([]byte(fmt.Sprintf("data: {\"symbol\":\"%s\", \"price\": %.2f, \"ts\": %d}\n\n", symbol, price, time.Now().UnixMilli()))); err != nil { break }
		time.Sleep(1 * time.Second)
	}
	return nil
}
