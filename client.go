package main

import (
  "fmt"
  "log"
  "net/rpc"
  "os"
)

type Command struct {
  Action string
  Pos, Len int
}

type Text struct {
  Contents string
}

func main() {
  client, err := rpc.Dial("unix", "/tmp/echo.sock")
  if err != nil {
    log.Fatal("dialing error: ", err)
  }
  command := &Command{Action:os.Args[1],Pos:7,Len:8,}
  var response Text
  err = client.Call("Editor.Send", command, &response)
  if err != nil {
    log.Fatal("editor.send error: ", err)
  }
  fmt.Printf("Got %s\n", response.Contents)
}
