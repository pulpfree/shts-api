package config

// defaults struct
type defaults struct {
	AWSRegion  string `yaml:"AWSRegion"`
	DBHost     string `yaml:"DBHost"`
	DBName     string `yaml:"DBName"`
	DBPassword string `yaml:"DBPassword"`
	DBUser     string `yaml:"DBUser"`
	SsmPath    string `yaml:"SsmPath"`
	Stage      string `yaml:"Stage"`
}

type config struct {
	AWSRegion    string
	DBConnectURL string
	DBName       string
	Stage        StageEnvironment
}
