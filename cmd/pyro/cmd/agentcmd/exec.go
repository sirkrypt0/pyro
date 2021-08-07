package agentcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

type execFlags struct {
	interactive bool
}

func newExecCmd(inr io.Reader, outw, errw io.WriteCloser) *cobra.Command {
	flags := &execFlags{}

	cmdExec := &cobra.Command{
		Use:   "exec <command...>",
		Short: "execute a command on the agent",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			var exitCode int
			var err error
			if flags.interactive {
				exitCode, err = apiClient.ExecuteInteractively(args, ctx, inr, outw, errw)
			} else {
				exitCode, err = apiClient.Execute(args, ctx, outw, errw)
			}
			if err != nil {
				return fmt.Errorf("error executing command: %w", err)
			}
			os.Exit(exitCode)
			return nil
		},
	}

	cmdExec.Flags().BoolVarP(&flags.interactive, "interactive", "i", false, "execute command interactively")
	return cmdExec
}
