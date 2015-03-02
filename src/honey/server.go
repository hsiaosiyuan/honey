package honey

import (
	"net"
	"log"
)

type Server struct {
	Conf *Config
}

func (s *Server) ListenAndServe() error {
	var (
		err error
		l net.Listener
	)

	if l, err = net.ListenTCP("tcp", s.Conf.la); err != nil {
		return err
	}

	return s.Serve(l)
}

func (s *Server) Serve(l net.Listener) error {
	defer l.Close()

	var (
		err error
		nc net.Conn
		ne net.Error
		ok bool
		c *conn
	)

	for {
		if nc, err = l.Accept(); err != nil {
			if ne, ok = err.(net.Error); ok && ne.Temporary() {
				continue
			}

			return err
		}

		if c, err = newConn(nc, s); err != nil {
			continue
		}

		go c.serve()
	}
}

func (s *Server) Log(args ...interface{}) {
	log.Println(args...)
}
