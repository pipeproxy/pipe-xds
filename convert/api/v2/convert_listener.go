package convert_api_v2

import (
	"encoding/json"
	"fmt"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
	convert_api_v2_core "github.com/wzshiming/envoy/convert/api/v2/core"
	convert_api_v2_listener "github.com/wzshiming/envoy/convert/api/v2/listener"
)

func Convert_Listener(conf *config.ConfigCtx, c *envoy_api_v2.Listener) (string, error) {
	name, err := convert_api_v2_core.Convert_AddressListener(conf, c.Address)
	if err != nil {
		return "", err
	}
	d, err := config.MarshalRef(name)
	if err != nil {
		return "", err
	}

	name = config.XdsName(c.Name + ".listener")
	listenerName, err := conf.RegisterComponents(name, d)
	if err != nil {
		return "", err
	}

	if len(c.FilterChains) == 0 || len(c.FilterChains[0].Filters) == 0 {
		return "", fmt.Errorf("not filter chains")
	}

	var handlerName string
	switch len(c.FilterChains) {
	case 1:
		name, err := convert_api_v2_listener.Convert_FilterChain(conf, c.FilterChains[0])
		if err != nil {
			return "", err
		}

		//TODO remove this
		if name == "" {
			return "", nil
		}

		handlerName = name
	default:
		list := []json.RawMessage{}
		for _, filterChain := range c.FilterChains {
			name, err := convert_api_v2_listener.Convert_FilterChain(conf, filterChain)
			if err != nil {
				return "", err
			}
			ref, err := config.MarshalRef(name)
			if err != nil {
				return "", err
			}

			list = append(list, ref)
		}
		d, err = config.MarshalKindStreamHandlerMulti(list)
		if err != nil {
			return "", err
		}
		name = config.XdsName(c.Name + ".filter-chains")
		handlerName, err = conf.RegisterComponents(name, d)
		if err != nil {
			return "", err
		}
	}

	listenerRef, err := config.MarshalRef(listenerName)
	if err != nil {
		return "", err
	}

	handlerRef, err := config.MarshalRef(handlerName)
	if err != nil {
		return "", err
	}

	d, err = config.MarshalKindServiceServer(listenerRef, handlerRef)
	if err != nil {
		return "", err
	}

	name = config.XdsName(c.Name)
	return conf.RegisterComponents(name, d)
}