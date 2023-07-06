package main

import (
	"log"

	"github.com/clly/gossip-glomers/node"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type EchoRequest struct {
	maelstrom.MessageBody
	Echo string `json:"echo"`
}

func main() {
	// ...
	n := maelstrom.NewNode()

	n.Handle("echo", node.NewHandler(func(req node.MaelstromRequest[EchoRequest]) error {
		b := req.Body
		b.Type = "echo_ok"
		return n.Reply(req.MaelstromMessage(), b)
	}))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
