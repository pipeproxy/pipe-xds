package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_auth_SdsSecretConfig(conf *config.ConfigCtx, c *envoy_api_v2_auth.SdsSecretConfig) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}