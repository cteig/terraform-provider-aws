package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	//"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aws.Provider})
}
