package merge

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "merge"

func init() {
	decode.Register(name, NewMergeWithConfig)
}

type Config struct {
	Merge []tls.TLS
}

func NewMergeWithConfig(conf *Config) (tls.TLS, error) {
	ts := make([]*tls.Config, 0, len(conf.Merge))
	for _, v := range conf.Merge {
		ts = append(ts, v.TLS())
	}
	return tls.WrapTLS(NewMerge(ts)), nil
}