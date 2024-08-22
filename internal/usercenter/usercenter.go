package usercenter

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/jinzhu/copier"

	"github.com/dawn303/cc/internal/pkg/bootstrap"
	"github.com/dawn303/cc/internal/usercenter/server"
	"github.com/dawn303/cc/pkg/db"
	"github.com/dawn303/cc/pkg/log"
	genericoptions "github.com/dawn303/cc/pkg/options"
	"github.com/dawn303/cc/pkg/version"
)

var (
	Name  = "cc-usercenter"
	ID, _ = os.Hostname()
)

// Config represents the configuration of the service.
type Config struct {
	HTTPOptions  *genericoptions.HTTPOptions
	TLSOptions   *genericoptions.TLSOptions
	JWTOptions   *genericoptions.JWTOptions
	MySQLOptions *genericoptions.MySQLOptions
	RedisOptions *genericoptions.RedisOptions
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() completedConfig {
	return completedConfig{cfg}
}

// completedConfig holds the configuration after it has been completed.
type completedConfig struct {
	*Config
}

// New returns a new instance of Server from the given config.
func (c completedConfig) New(stopCh <-chan struct{}) (*Server, error) {
	appInfo := bootstrap.NewAppInfo(ID, Name, version.Get().String())

	conf := &server.Config{
		HTTP: *c.HTTPOptions,
		TLS:  *c.TLSOptions,
	}

	var dbOptions db.MySQLOptions
	_ = copier.Copy(&dbOptions, c.MySQLOptions)

	// Initialize Kratos application with the provided configurations.
	app, cleanup, err := wireApp(appInfo, conf, &dbOptions)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	return &Server{app: app}, nil
}

// Server represents the server.
type Server struct {
	app *kratos.App
}

// Run is a method of the Server struct that starts the server.
func (s *Server) Run(stopCh <-chan struct{}) error {
	go func() {
		if err := s.app.Run(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-stopCh

	log.Infof("Gracefully shutting down server ...")

	if err := s.app.Stop(); err != nil {
		log.Errorw(err, "Failed to gracefully shutdown kratos application")
		return err
	}

	return nil
}
