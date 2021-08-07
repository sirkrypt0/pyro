package cmd

import (
	"github.com/sirkrypt0/pyro/cmd/pyro/cmd/agentcmd"
	"github.com/spf13/cobra"
	"io"
)

func NewPyroCommand(inr io.Reader, outw, errw io.WriteCloser) *cobra.Command {
	root := &cobra.Command{
		Use:   "pyro",
		Short: "A cli for interacting with pyro services",
	}
	root.AddCommand(agentcmd.NewAgentCmd(inr, outw, errw))
	return root
}
