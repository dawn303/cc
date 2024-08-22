package server

import (
	genericoptions "github.com/dawn303/cc/pkg/options"
)

type Config struct {
	HTTP genericoptions.HTTPOptions
	TLS  genericoptions.TLSOptions
}
