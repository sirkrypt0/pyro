package client

import (
	agentv1 "github.com/sirkrypt0/pyro/api/agent/v1"
)

type Client struct {
	agent agentv1.AgentServiceClient
}

func NewClient(agent agentv1.AgentServiceClient) (*Client, error) {
	return &Client{agent: agent}, nil
}
