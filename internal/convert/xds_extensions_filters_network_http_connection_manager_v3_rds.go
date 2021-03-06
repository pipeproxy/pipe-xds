package convert

import (
	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"

	"log"
)

func Convert_extensions_filters_network_http_connection_manager_v3_Rds(conf *config.ConfigCtx, c *envoy_extensions_filters_network_http_connection_manager_v3.Rds) (bind.HTTPHandler, error) {
	name := c.RouteConfigName
	if name != "" {
		return conf.RDS(name), nil
	}

	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_filters_network_http_connection_manager_v3.Rds %s\n", string(data))
	return nil, nil
}
