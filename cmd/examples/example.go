package main

import (
	"context"
	"fmt"

	"github.com/magicbell/magicbell-go-project-client/pkg/broadcasts"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
)

func main() {
	config := magicbellprojectclientconfig.NewConfig()
	client := magicbellprojectclient.NewMagicbellProjectClient(config)

	params := broadcasts.ListBroadcastsRequestParams{}

	response, err := client.Broadcasts.ListBroadcasts(context.Background(), params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}
