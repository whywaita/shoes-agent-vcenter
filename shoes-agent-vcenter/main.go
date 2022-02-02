package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-plugin"
	"github.com/whywaita/shoes-agent-vcenter/shoes-agent-vcenter/vcenter"
	"github.com/whywaita/shoes-agent/shoes-agent/pkg/shoesprovider"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	vc, err := vcenter.New(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize vCenter: %w", err)
	}

	handshake := plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "SHOES_PLUGIN_MAGIC_COOKIE",
		MagicCookieValue: "are_you_a_shoes?",
	}
	pluginMap := map[string]plugin.Plugin{
		"shoes_grpc": &shoesprovider.AgentPlugin{
			Backend:   vc,
			ShoesType: "shoes-agent-vcenter",
		},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshake,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})

	return nil
}
