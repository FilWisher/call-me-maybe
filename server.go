package main

import (
  "net"
  "net/rpc"
  "log"
  "os"
  "os/signal"
  "syscall"
  "fmt"
)

type Editor struct {
  Hello string
}

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

func main() {
  editor := new(Editor)
  rpc.Register(editor)

  l, err := net.Listen("unix", "/tmp/echo.sock")
  defer os.Remove("/tmp/echo.sock")

  if err != nil {
    log.Fatal("listen error: ", err)
  }

  go rpc.Accept(l)
  wait()
}

func wait() {
  signals := make(chan os.Signal)
  signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP)
  <-signals
}
