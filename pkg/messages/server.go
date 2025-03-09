package messages

import (
	"fmt"
	"strconv"

	"github.com/lunaris/p10go/pkg/types"
)

type Server struct {
	Name           string
	HopCount       int
	StartTimestamp int64
	LinkTimestamp  int64
	Protocol       Protocol
	Numeric        types.ServerNumeric
	MaxConnections types.ClientNumeric
	Description    string
}

func ParseServer(tokens []string) (*Server, error) {
	if len(tokens) < 9 {
		return nil, fmt.Errorf("SERVER: expected at least 9 tokens; received %d", len(tokens))
	}

	if tokens[0] != "SERVER" {
		return nil, fmt.Errorf("SERVER: expected SERVER; received %s", tokens[0])
	}

	hopCount, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse hop count: %w", err)
	}

	startTimestamp, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse start timestamp: %w", err)
	}

	linkTimestamp, err := strconv.ParseInt(tokens[4], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse link timestamp: %w", err)
	}

	protocol, err := ParseProtocol(tokens[5])
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse protocol: %w", err)
	}

	maxClientID, err := types.ParseClientID(tokens[6])
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse server numeric and maximum connection count: %w", err)
	}

	// tokens[7] is an unused placeholder

	description := lastParameter(tokens[8:])

	return &Server{
		Name:           tokens[1],
		HopCount:       hopCount,
		StartTimestamp: startTimestamp,
		LinkTimestamp:  linkTimestamp,
		Protocol:       protocol,
		Numeric:        maxClientID.Server,
		MaxConnections: maxClientID.Client,
		Description:    description,
	}, nil
}

func (s *Server) String() string {
	return fmt.Sprintf(
		"SERVER %s %d %d %d %s %s%s 0 :%s",
		s.Name,
		s.HopCount,
		s.StartTimestamp,
		s.LinkTimestamp,
		s.Protocol,
		s.Numeric,
		s.MaxConnections,
		s.Description,
	)
}
