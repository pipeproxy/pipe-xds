package convert

import (
	envoy_config_filter_network_http_connection_manager_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_filter_network_http_connection_manager_v2_ScopedRoutes_ScopeKeyBuilder_FragmentBuilder_HeaderValueExtractor(conf *config.ConfigCtx, c *envoy_config_filter_network_http_connection_manager_v2.ScopedRoutes_ScopeKeyBuilder_FragmentBuilder_HeaderValueExtractor) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}