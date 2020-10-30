package convert

import (
	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"

	"log"
)

func Convert_extensions_filters_network_http_connection_manager_v3_ScopedRoutes_ScopeKeyBuilder_FragmentBuilder_HeaderValueExtractor(conf *config.ConfigCtx, c *envoy_extensions_filters_network_http_connection_manager_v3.ScopedRoutes_ScopeKeyBuilder_FragmentBuilder_HeaderValueExtractor) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_filters_network_http_connection_manager_v3.ScopedRoutes_ScopeKeyBuilder_FragmentBuilder_HeaderValueExtractor %s\n", string(data))
	return "", nil
}