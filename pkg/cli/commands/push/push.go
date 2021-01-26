package push

import (
	"github.com/pkg/errors"
	"github.com/rancher/kim/pkg/client"
	"github.com/rancher/kim/pkg/client/action"
	wrangler "github.com/rancher/wrangler-cli"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return wrangler.Command(&CommandSpec{}, cobra.Command{
		Use:   "push [OPTIONS] IMAGE",
		Short: "Push an image",
	})
}

type CommandSpec struct {
	action.PushImage
}

func (s *CommandSpec) Run(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("exactly one argument is required")
	}

	k8s, err := client.DefaultConfig.Interface()
	if err != nil {
		return err
	}
	return s.PushImage.Invoke(cmd.Context(), k8s, args[0])
}
