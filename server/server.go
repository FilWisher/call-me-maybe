package main

import (
  "net"
  "net/rpc"
  "log"
  "os"
  "os/signal"
  "syscall"
  "github.com/filwisher/call-me-maybe/text"
)

func main() {

  editor := new(text.Editor)
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
