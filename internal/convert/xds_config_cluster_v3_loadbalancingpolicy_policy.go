package convert

import (
	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_config_cluster_v3_LoadBalancingPolicy_Policy(conf *config.ConfigCtx, c *envoy_config_cluster_v3.LoadBalancingPolicy_Policy) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_cluster_v3.LoadBalancingPolicy_Policy %s\n", string(data))
	return "", nil
}
