package config

var serviceName, version, buildTime = "", "", ""

var (
	buildCfg *buildConfig
	Cfg      conf
)

type buildConfig struct {
	ServiceName string `toml:"service_name" toml:"service_name"`
	Version     string `toml:"version" toml:"version"`
	BuildTime   string `toml:"build_time" toml:"build_time"`
	ConfigFile  string `toml:"config_file" toml:"config_file"`
}

// nolint:gochecknoinits
func init() {
	buildCfg = &buildConfig{
		ServiceName: serviceName,
		Version:     version,
		BuildTime:   buildTime,
	}
}

func GetBuildCfg() *buildConfig {
	return buildCfg
}

type conf struct {
	RepoUrl    string             `toml:"repo_url"`
	ArgoConfig argoConfig         `toml:"argo_config"`
	AppList    map[string]AppInfo `toml:"app_list"` // key: appName, value: namespace
}

type AppInfo struct {
	Namespace  string `toml:"namespace"`
	SourcePath string `toml:"source_path"`
}

type argoConfig struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}
