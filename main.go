package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version = "0.0.24"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	//plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.New(version)})
	username := "buildbot@hashdata.cn"
	password := "Passw0rd"
	client_id := "bad16277-8da4-48ec-963d-b5515af3aa82"
	client_secret := "bad16277-8da4-48ec-963d-b5515af3aa82"
	end_point := "https://console.hashdata.xyz"
	//cfg := cloudmgr.NewConfiguration()
	configOauth2 := oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Endpoint: oauth2.Endpoint{
			TokenURL: end_point + "/account/oauth/token",
		},
	}
	token, err := configOauth2.PasswordCredentialsToken(context.Background(), username, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token.AccessToken)
}
