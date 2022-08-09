package lab

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

type ServerConn struct {
	Connection net.Conn
	ID         string
	Open       bool
}

func ShowConnection(p *atomic.Pointer[ServerConn]) {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println(p, p.Load())
	}

}
func AtomicPointer() {
	c := make(chan bool)
	p := atomic.Pointer[ServerConn]{}
	s := ServerConn{ID: "first_conn"}
	p.Store(&s)
	go ShowConnection(&p)
	go func() {
		i := 0
		for {
			time.Sleep(13 * time.Second)
			newConn := ServerConn{ID: fmt.Sprint("new_conn ", i)}
			p.Swap(&newConn)
			i++
		}
	}()
	<-c
}
