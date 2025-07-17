package main

type Args struct {
	X, Y int
}

type ServiceA struct {
}

func (sa *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}
