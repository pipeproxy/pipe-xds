package convert

import (
	envoy_config_filter_network_tcp_proxy_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_filter_network_tcp_proxy_v2_TcpProxy_WeightedCluster_ClusterWeight(conf *config.ConfigCtx, c *envoy_config_filter_network_tcp_proxy_v2.TcpProxy_WeightedCluster_ClusterWeight) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}