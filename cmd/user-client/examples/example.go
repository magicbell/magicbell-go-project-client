package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/magicbell/magicbell-go/pkg/user-client/userclient"
	"github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"

	"github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

func main() {
	loadEnv()

	config := userclientconfig.NewConfig()
	client := userclient.NewUserClient(config)

	params := channels.GetInAppInboxTokensRequestParams{}

	response, err := client.Channels.GetInAppInboxTokens(context.Background(), params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
