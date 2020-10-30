package convert

import (
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"

	"log"
)

func Convert_config_core_v3_HealthCheck_TlsOptions(conf *config.ConfigCtx, c *envoy_config_core_v3.HealthCheck_TlsOptions) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.HealthCheck_TlsOptions %s\n", string(data))
	return "", nil
}