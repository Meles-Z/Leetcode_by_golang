package main

import "fmt"

type OptFunc func(*Opts)
type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "m1",
		tls:     false,
	}
}

func withTls(opts *Opts) {
	opts.tls = true
}
func withId(id string) OptFunc {
	return func(o *Opts) {
		o.id = id
	}
}

func maxConn(maxConn int) OptFunc {
	return func(o *Opts) {
		o.maxConn = maxConn
	}
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	df := defaultOpts()
	for _, fn := range opts {
		fn(&df)
	}
	return &Server{
		Opts: df,
	}
}

func main() {
	s := newServer(withTls, withId("m2"), maxConn(99))
	fmt.Println(s)
}
