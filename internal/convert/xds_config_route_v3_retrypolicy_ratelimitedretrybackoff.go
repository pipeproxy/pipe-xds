package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
)

func Convert_config_route_v3_RetryPolicy_RateLimitedRetryBackOff(conf *config.ConfigCtx, c *envoy_config_route_v3.RetryPolicy_RateLimitedRetryBackOff) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RetryPolicy_RateLimitedRetryBackOff %s\n", string(data))
	return "", nil
}
