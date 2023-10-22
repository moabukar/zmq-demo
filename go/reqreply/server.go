package main

import (
	"github.com/gofiber/fiber/v2"
	zmq "github.com/pebbe/zmq4"
)

type Input struct {
	Text string `json:"text"`
}

func main() {
	app := fiber.New()
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REQ)
	s.Bind("tcp://*:6666")

	app.Post("/", func(c *fiber.Ctx) error {
		input := new(Input)
		if err := c.BodyParser(input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Failed to parse body"})
		}

		if _, err := s.Send(input.Text, 0); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to send message"})
		}

		msg, err := s.Recv(0)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to receive message"})
		}

		return c.JSON(fiber.Map{"Length": msg})
	})

	app.Listen(":3000")
}
