package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_FilterAction(conf *config.ConfigCtx, c *envoy_config_route_v3.FilterAction) (bind.HTTPHandler, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.FilterAction %s\n", string(data))
	return nil, nil
}
