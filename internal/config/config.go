package config

type Notify struct {
	Send   bool   `yaml:"send"`
	ChatId int64  `yaml:"chat_id"`
	Token  string `yaml:"token"`
}

type Build struct {
	InputPoint  string `yaml:"input"`
	OutputPoint string `yaml:"output"`
}

type Config struct {
	Port        int
	Host        string
	TempDirName string
	BuildCfg    Build  `yaml:"build"`
	NotifyCfg   Notify `yaml:"notify"`
}
