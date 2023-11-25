package main

import (
	"fmt"
	"os"

	"github.com/bungolow-dev/bungolow/pkg/application/plugins"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type Roku struct {
	ip   string
	port int16
}

func (r *Roku) Query() string {

	return "HELLO FROM ROKU PLUGIN"
}

func (r *Roku) Register() error {

	return fmt.Errorf("RETURNED ERROR")
}

func (r *Roku) Initialize(settings map[string]interface{}) {
	fmt.Println(settings)
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	roku := &Roku{
		ip:   "10.0.0.181",
		port: 8060,
	}

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"roku": &plugins.DevicePlugin{Impl: roku},
	}

	logger.Debug("message from plugin", "foo", "bar")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
