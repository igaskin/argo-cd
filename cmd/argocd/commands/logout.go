package commands

import (
	"fmt"

	"github.com/argoproj/argo-cd/errors"
	argocdclient "github.com/argoproj/argo-cd/pkg/apiclient"
	"github.com/argoproj/argo-cd/util/localconfig"
	"github.com/spf13/cobra"
)

var (
	server string
)

// NewLogoutCommand delete token from user
func NewLogoutCommand(globalClientOpts *argocdclient.ClientOptions) *cobra.Command {
	var command = &cobra.Command{
		Use:   "logout SERVER",
		Short: "Log out of Argo CD",
		Long:  "Log out of Argo CD",
		Run: func(c *cobra.Command, args []string) {
			localCfg, err := localconfig.ReadLocalConfig(globalClientOpts.ConfigPath)
			errors.CheckError(err)
			if len(args) > 0 {
				server = args[0]
			} else {
				server = localCfg.CurrentContext
			}
			currentContext, err := localCfg.ResolveContext(server)
			errors.CheckError(err)
			localCfg.UpsertUser(localconfig.User{
				Name: currentContext.Name,
			})
			err = localconfig.WriteLocalConfig(*localCfg, globalClientOpts.ConfigPath)
			errors.CheckError(err)
			fmt.Printf("Logged out from '%s'\n", server)
		},
	}

	return command
}
