package client

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	xnet "github.com/Bl1tz23/wow/pkg/net"
)

type Client struct {
	serverAddr string

	solver solver
}

type solver interface {
	Solve(ctx context.Context, token []byte) ([]byte, error)
}

func NewClient(addr string, s solver) *Client {
	return &Client{
		serverAddr: addr,
		solver:     s,
	}
}

func (c *Client) GetQuote(reqMsg []byte) (string, error) {
	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		return "", fmt.Errorf("failed to establish a connection with server: %w", err)
	}
	defer conn.Close()

	err = xnet.WriteConn(conn, reqMsg)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	task, err := xnet.ReadConn(conn)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	solution, err := c.solver.Solve(ctx, task)
	if err != nil {
		return "", fmt.Errorf("failed to find a solution: %w", err)
	}

	err = xnet.WriteConn(conn, solution)
	if err != nil {
		return "", fmt.Errorf("failed to write solution to conn: %w", err)
	}

	quote, err := xnet.ReadConn(conn)
	if err != nil {
		return "", fmt.Errorf("failed to read quote from server: %w", err)
	}

	if strings.Contains(string(quote), "wrong answer") {
		fmt.Printf("token: %v solution: %v\n", task, solution)
	}

	if len(quote) == 0 {
		return "", fmt.Errorf("empty quote string in response")
	}

	return string(quote), nil
}
