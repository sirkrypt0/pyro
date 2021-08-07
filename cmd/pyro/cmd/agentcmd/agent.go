package agentcmd

import (
	"fmt"
	agentv1 "github.com/sirkrypt0/pyro/api/agent/v1"
	"github.com/sirkrypt0/pyro/pkg/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"io"
)

var (
	agentAddr string
	apiClient *client.Client
	closeConn func() error
)

func NewAgentCmd(inr io.Reader, outw, errw io.WriteCloser) *cobra.Command {
	cmd := &cobra.Command{
		Use:                "agent",
		Short:              "Interact with the pyro agent",
		PersistentPreRunE:  newAgentServiceClient,
		PersistentPostRunE: closeAgentServiceClient,
	}
	cmd.PersistentFlags().StringVarP(&agentAddr, "addr", "a", "127.0.0.1:3000", "Address of the pyro agent")

	cmd.AddCommand(newExecCmd(inr, outw, errw))

	return cmd
}

func newAgentServiceClient(_ *cobra.Command, _ []string) error {
	cc, err := grpc.Dial(agentAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("error connecting to agent: %w", err)
	}
	closeConn = cc.Close
	agent := agentv1.NewAgentServiceClient(cc)
	apiClient, err = client.NewClient(agent)
	if err != nil {
		return fmt.Errorf("error creating new api client: %w", err)
	}
	return nil
}

func closeAgentServiceClient(_ *cobra.Command, _ []string) error {
	if err := closeConn(); err != nil {
		return fmt.Errorf("error closing connection to agent: %w", err)
	}
	return nil
}
