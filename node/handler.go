package node

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type MaelstromRequest[T any] struct {
	Src      string `json:"src"`
	Dest     string `json:"dest"`
	Body     *T     `json:"body"`
	original json.RawMessage
}

func (r *MaelstromRequest[T]) MaelstromMessage() maelstrom.Message {
	return maelstrom.Message{
		Src:  r.Src,
		Dest: r.Dest,
		Body: r.original,
	}
}

type MaelstromHandlerRequest[T any] func(req MaelstromRequest[T]) error

func NewHandler[T any](fn MaelstromHandlerRequest[T]) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		var body = new(T)
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		req := MaelstromRequest[T]{
			Src:      msg.Src,
			Dest:     msg.Dest,
			Body:     body,
			original: msg.Body,
		}

		return fn(req)
	}
}
