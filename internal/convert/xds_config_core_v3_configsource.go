package convert

import (
	"log"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_core_v3_ConfigSource(conf *config.ConfigCtx, c *envoy_config_core_v3.ConfigSource) (bind.StreamDialer, error) {
	switch s := c.ConfigSourceSpecifier.(type) {
	case *envoy_config_core_v3.ConfigSource_Path:
	case *envoy_config_core_v3.ConfigSource_ApiConfigSource:
		return Convert_config_core_v3_ApiConfigSource(conf, s.ApiConfigSource)
	case *envoy_config_core_v3.ConfigSource_Ads:
		return conf.CDS("ads"), nil
	case *envoy_config_core_v3.ConfigSource_Self:
	}

	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.ConfigSource %s\n", string(data))
	return nil, nil
}
