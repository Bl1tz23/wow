package tcp

import (
	"fmt"
	"net"
	"sync"

	"github.com/Bl1tz23/wow/pkg/logger"
)

const (
	DefaultNetwork = "tcp"
)

var (
	log = logger.Logger().Named("tcp-server").Sugar()
)

type TCPServer struct {
	l net.Listener

	wg    sync.WaitGroup
	close chan struct{}

	tasksProvider tasksProvider
	quoteBook     quoteBook
}

type tasksProvider interface {
	NewTask() ([]byte, error)
	Verify(token, nonce []byte) (bool, error)
}

type quoteBook interface {
	GetRandomQoute() []byte
}

func New(addr string, tasksProvider tasksProvider, quoteBook quoteBook) (*TCPServer, error) {
	l, err := net.Listen(DefaultNetwork, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen with addr: %s, err: %w", addr, err)
	}

	return &TCPServer{
		l:             l,
		close:         make(chan struct{}),
		tasksProvider: tasksProvider,
		quoteBook:     quoteBook,
	}, nil
}

func (server *TCPServer) Run() error {
	for {
		conn, err := server.l.Accept()
		if err != nil {
			select {
			case <-server.close:
				return nil
			default:
				return fmt.Errorf("failed to accept connection: %w", err)
			}
		}

		server.wg.Add(1)
		go func() {
			err := server.handleRequest(conn)
			if err != nil {
				log.Error("failed to handle conn: %s", err)
			}
			server.wg.Done()
		}()
	}
}

func (server *TCPServer) Close() error {
	close(server.close)
	err := server.l.Close()
	if err != nil {
		return fmt.Errorf("failed to close tcp server: %w", err)
	}
	server.wg.Wait()
	return nil
}
