package convert

import (
	"log"

	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_filters_network_http_connection_manager_v3_HttpConnectionManager(conf *config.ConfigCtx, c *envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager, tls bind.TLS) (bind.StreamHandler, error) {
	var route bind.HTTPHandler
	switch r := c.RouteSpecifier.(type) {
	case *envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager_Rds:
		handler, err := Convert_extensions_filters_network_http_connection_manager_v3_Rds(conf, r.Rds)
		if err != nil {
			return nil, err
		}
		route = handler
	case *envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager_RouteConfig:
		handler, err := Convert_config_route_v3_RouteConfiguration(conf, r.RouteConfig)
		if err != nil {
			return nil, err
		}
		route = handler
	case *envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager_ScopedRoutes:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_extensions_filters_network_http_connection_manager_v3.TcpProxy %s\n", string(data))
		return nil, nil
	}

	var err error
	for _, l := range c.AccessLog {
		route, err = Convert_config_accesslog_v3_AccessLog_HTTPHandler(conf, l, route)
		if err != nil {
			return nil, err
		}
	}

	var d = bind.HTTP2StreamHandlerConfig{
		Handler: route,
		TLS:     tls,
	}
	return d, nil
}
