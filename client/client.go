package main

import (
  "fmt"
  "log"
  "net/rpc"
  "os"
  "github.com/filwisher/call-me-maybe/text"
)

func main() {
  client, err := rpc.Dial("unix", "/tmp/echo.sock")
  if err != nil {
    log.Fatal("dialing error: ", err)
  }
  if len(os.Args) < 2 {
    log.Fatal("client needs at least one argument")
  }
  command := &text.Command{Action:os.Args[1],Pos:7,Len:8,}
  var response text.Text
  err = client.Call("Editor.Send", command, &response)
  if err != nil {
    log.Fatal("editor.send error: ", err)
  }
  fmt.Printf("Got %s\n", response.Contents)
}
