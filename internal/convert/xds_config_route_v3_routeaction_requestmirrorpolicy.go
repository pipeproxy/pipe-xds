package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_config_route_v3_RouteAction_RequestMirrorPolicy(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteAction_RequestMirrorPolicy) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RouteAction_RequestMirrorPolicy %s\n", string(data))
	return "", nil
}
