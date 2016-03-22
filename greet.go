package test


import (
    "github.com/gorilla/mux"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
    "log"
    "net/http"
)

type Service string

type Args struct {
	Name string
	ID int
}

type Reply string

func (t *Service) Greet(args *Args, reply *Reply) error {
	*reply = "Greetings!" + args.Name + "with id" + args.id
	return nil
}