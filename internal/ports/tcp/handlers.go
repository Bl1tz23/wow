package tcp

import (
	"fmt"
	"net"
	"time"

	xnet "github.com/Bl1tz23/wow/pkg/net"
)

const (
	defaultReadDeadline = 20 * time.Second
	defaultDeadline     = 30 * time.Second
)

func (server *TCPServer) handleRequest(conn net.Conn) error {
	defer conn.Close()

	err := conn.SetReadDeadline(time.Now().Add(defaultReadDeadline))
	if err != nil {
		return fmt.Errorf("failed to set read deadline for conn: %w", err)
	}

	err = conn.SetDeadline(time.Now().Add(defaultDeadline))
	if err != nil {
		return fmt.Errorf("failed to set deadline for conn: %w", err)
	}

	buf, err := xnet.ReadConn(conn)
	if err != nil {
		return err
	}

	log.With("req_msg", string(buf)).Debug("client request")

	task, err := server.tasksProvider.NewTask()
	if err != nil {
		return fmt.Errorf("failed to create new task: %w", err)
	}

	err = xnet.WriteConn(conn, task)
	if err != nil {
		return err
	}

	buf, err = xnet.ReadConn(conn)
	if err != nil {
		return err
	}

	log.With("req_msg", string(buf)).Debug("client solution")

	ok, err := server.tasksProvider.Verify(task, buf)
	if err != nil {
		return fmt.Errorf("failed to verify proof: %w", err)
	}
	if !ok {
		err = xnet.WriteConn(conn, []byte("wrong answer"))
		if err != nil {
			return err
		}
		return nil
	}

	quote := server.quoteBook.GetRandomQoute()

	log.With("quote", string(quote)).Debug("sending to client")

	err = xnet.WriteConn(conn, quote)
	if err != nil {
		return err
	}

	return nil
}
