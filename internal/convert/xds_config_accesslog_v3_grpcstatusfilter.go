package convert

import (
	"log"

	envoy_config_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
)

func Convert_config_accesslog_v3_GrpcStatusFilter(conf *config.ConfigCtx, c *envoy_config_accesslog_v3.GrpcStatusFilter) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_accesslog_v3.GrpcStatusFilter %s\n", string(data))
	return "", nil
}
