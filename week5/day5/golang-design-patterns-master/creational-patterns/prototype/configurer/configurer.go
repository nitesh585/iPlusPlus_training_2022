package configurer

// Config provides a configuration of microservice
type Config struct {
	workDir string
	user    string
}

// NewConfig create a new config
func NewConfig(user string, workDir string) Config {
	return Config{
		user:    user,
		workDir: workDir,
	}
}

// WithWorkDir creates a copy of Config with the provided working directory
func (c Config) WithWorkDir(dir string) Config {
	c.workDir = dir
	return c
}

// WithUser creates a copy of Config with the provided user
func (c Config) WithUser(user string) Config {
	c.user = user
	return c
}
