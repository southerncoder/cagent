package root

import (
	"context"
	"fmt"

	"github.com/docker/cagent/internal/telemetry"
	registrytypes "github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/registry"
	"github.com/spf13/cobra"
)

func NewLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Authenticate to agent registry",
		Long:  `Authenticate to a registry to push and pull agents`,
		Run: func(cmd *cobra.Command, args []string) {
			telemetry.TrackCommand("login", args)
			runLogin()

		},
	}
}

func runLogin() (registrytypes.AuthenticateOKBody, error) {
	ctx := context.Background()
	auth := registrytypes.AuthConfig{
		ServerAddress: "https://index.docker.io/v1/",
	}
	userAgent := "Docker-Client/v28.4.0 (darwin)"
	fmt.Println("UserAgent:", userAgent)
	svc, err := registry.NewService(registry.ServiceOptions{})
	if err != nil {
		return registrytypes.AuthenticateOKBody{}, err
	}

	status, token, err := svc.Auth(ctx, &auth, userAgent)

	fmt.Println("Login successful status:", status, "token:", token)

	return registrytypes.AuthenticateOKBody{
		Status:        status,
		IdentityToken: token,
	}, err
}
