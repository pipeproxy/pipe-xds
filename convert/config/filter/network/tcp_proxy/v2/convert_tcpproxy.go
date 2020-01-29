package convert_config_filter_network_tcp_proxy_v2

import (
	envoy_config_filter_network_tcp_proxy_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_TcpProxy(conf *config.ConfigCtx, c *envoy_config_filter_network_tcp_proxy_v2.TcpProxy) (string, error) {
	switch c.StatPrefix {
	case "tcp":
		switch s := c.ClusterSpecifier.(type) {
		case *envoy_config_filter_network_tcp_proxy_v2.TcpProxy_Cluster:
			return config.XdsName(s.Cluster), nil
		case *envoy_config_filter_network_tcp_proxy_v2.TcpProxy_WeightedClusters:
		}
	}
	logger.Todof("%#v", c)
	return "", nil
}