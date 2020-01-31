package convert

import (
	envoy_api_v2_endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_endpoint_LbEndpoint(conf *config.ConfigCtx, c *envoy_api_v2_endpoint.LbEndpoint) (string, error) {
	switch h := c.HostIdentifier.(type) {
	case *envoy_api_v2_endpoint.LbEndpoint_Endpoint:
		return Convert_api_v2_core_AddressForward(conf, h.Endpoint.Address)
	case *envoy_api_v2_endpoint.LbEndpoint_EndpointName:
	}

	logger.Todof("%#v", c)
	return "", nil
}