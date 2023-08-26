package main

import (
	"fmt"
	"time"
)

func main() {

	s := NewServer()

	go s.Start() // start in a goroutine because we don't want this to block

	<-time.After(time.Second * 10)
	// time.Sleep(time.Second * 10) // works !!
	// main routine waits for 10 seconds

	s.Shutdown()

}

type Server struct {
	quitChan chan int
}

func NewServer() *Server {
	return &Server{
		quitChan: make(chan int, 1),
	}
	// channel is a buffered channel we don't want it to hold
	// after signal is passed
}

func (s *Server) Start() {

	// closing a closed channel twice is a problem
	for {

		select {
		case <-s.quitChan:
			fmt.Println("Server shutting down gracefully!!")
			// dont need worry about closing the channel twice
			return
		default:
			fmt.Println("Server Running!!")
			time.Sleep(time.Second * 2)
		}

	}
}

func (s *Server) Shutdown() {
	// some shutdown logic
	fmt.Println("Shutdown server called !!!")
	close(s.quitChan)
}
