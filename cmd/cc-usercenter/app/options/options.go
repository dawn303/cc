package options

import (
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/dawn303/cc/internal/pkg/client"
	"github.com/dawn303/cc/internal/pkg/feature"
	"github.com/dawn303/cc/internal/usercenter"
	"github.com/dawn303/cc/pkg/app"
	"github.com/dawn303/cc/pkg/log"
	genericoptions "github.com/dawn303/cc/pkg/options"
)

var _ app.CliOptions = (*Options)(nil)

type Options struct {
	HTTPOptions  *genericoptions.HTTPOptions  `json:"http" mapstructure:"http"`
	TLSOptions   *genericoptions.TLSOptions   `json:"tls" mapstructure:"tls"`
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	RedisOptions *genericoptions.RedisOptions `json:"redis" mapstructure:"redis"`
	JWTOptions   *genericoptions.JWTOptions   `json:"jwt" mapstructure:"jwt"`
	Log          *log.Options                 `json:"log" mapstructure:"log"`
}

// NewOptions returns initialized Options.
func NewOptions() *Options {
	o := &Options{
		// GenericOptions: genericoptions.NewOptions(),
		HTTPOptions:  genericoptions.NewHTTPOptions(),
		TLSOptions:   genericoptions.NewTLSOptions(),
		MySQLOptions: genericoptions.NewMySQLOptions(),
		RedisOptions: genericoptions.NewRedisOptions(),
		JWTOptions:   genericoptions.NewJWTOptions(),
		Log:          log.NewOptions(),
	}

	return o
}

// Flags returns flags for a specific server by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	o.TLSOptions.AddFlags(fss.FlagSet("tls"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.JWTOptions.AddFlags(fss.FlagSet("jwt"))
	o.Log.AddFlags(fss.FlagSet("log"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	client.AddFlags(fs)
	feature.DefaultMutableFeatureGate.AddFlag(fs)

	return fss
}

// Complete completes all the required options.
func (o *Options) Complete() error {
	return nil
}

// Validate validates all the required options.
func (o *Options) Validate() error {
	errs := []error{}

	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.RedisOptions.Validate()...)
	errs = append(errs, o.JWTOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)

	return utilerrors.NewAggregate(errs)
}

func (o *Options) ApplyTo(c *usercenter.Config) error {
	c.HTTPOptions = o.HTTPOptions
	c.TLSOptions = o.TLSOptions
	c.MySQLOptions = o.MySQLOptions
	c.RedisOptions = o.RedisOptions
	c.JWTOptions = o.JWTOptions
	return nil
}

func (o *Options) Config() (*usercenter.Config, error) {
	c := &usercenter.Config{}

	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}

	return c, nil
}
