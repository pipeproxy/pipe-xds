package convert_api_v2

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_UnimplementedScopedRoutesDiscoveryServiceServer(conf *config.ConfigCtx, c *envoy_api_v2.UnimplementedScopedRoutesDiscoveryServiceServer) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}