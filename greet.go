package test



type Service string

type Args struct {
	Name string
}

type Reply string

func (t *Service) Greet(args *Args, reply *Reply) error {
	*reply = "Greetings!" + args.Name
	return nil
}