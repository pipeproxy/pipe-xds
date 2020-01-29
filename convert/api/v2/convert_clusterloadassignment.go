package convert_api_v2

import (
	"encoding/json"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
	convert_api_v2_endpoint "github.com/wzshiming/envoy/convert/api/v2/endpoint"
)

func Convert_ClusterLoadAssignment(conf *config.ConfigCtx, c *envoy_api_v2.ClusterLoadAssignment) (string, error) {
	switch len(c.Endpoints) {
	case 1:
		name, err := convert_api_v2_endpoint.Convert_LocalityLbEndpoints(conf, c.Endpoints[0])
		if err != nil {
			return "", err
		}

		ref, err := config.MarshalRef(name)
		if err != nil {
			return "", err
		}

		name = config.XdsName(c.ClusterName)

		return conf.RegisterComponents(name, ref)

	default:
		list := []json.RawMessage{}
		for _, endpoint := range c.Endpoints {
			name, err := convert_api_v2_endpoint.Convert_LocalityLbEndpoints(conf, endpoint)
			if err != nil {
				return "", err
			}
			ref, err := config.MarshalRef(name)
			if err != nil {
				return "", err
			}

			list = append(list, ref)
		}

		d, err := config.MarshalKindStreamHandlerPoller("round_robin", list)
		if err != nil {
			return "", err
		}

		name := config.XdsName(c.ClusterName)

		return conf.RegisterComponents(name, d)
	}
}