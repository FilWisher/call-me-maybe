package text

import (
  "fmt"
)

type Editor struct {}

type Command struct {
  Action string
  Pos, Len int
}

type Text struct {
  Contents string
}

func (e *Editor) Send(command Command, response *Text) error {
  fmt.Printf("Got command: action - %s; pos - %d; len - %d\n", command.Action, command.Pos, command.Len)

  response.Contents = "Hello there"
  return nil
}

