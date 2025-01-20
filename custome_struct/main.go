package main

import "fmt"

// struct configration in golang
/*
 type Server struct{
    maxConn int
	id string
	tls bool
 }
func newServer(maxConn int, id string, tls bool) *Server {
	return &Server{
		maxConn: maxConn,
		id:      id,
		tls:     tls,
	}
}
*/

type OptFunc func(*Opts)
type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts{
	return Opts{
		maxConn: 10,
		id: "default",
		tls: false,
	}
}

func withTLS(opts *Opts){
	opts.tls=true
}

func withID(s string) OptFunc{
	return func(o *Opts) {
		o.id=s
	}
}

func withMaxConn(n int) OptFunc{
	return func(o *Opts) {
		o.maxConn=n
	}
}
type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o:=defaultOpts()
	for _, fn:= range opts{
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := newServer(withMaxConn(99), withID("m2"))
	fmt.Println(s)
}


