package config

type Notify struct{
	Send string `yaml:"send"`
	ChatId int64 `yaml:"chat_id"`
	Token string `yaml:"token"`
}

type Config struct{
	Port int
	Host string
	TempDirName string
	NotifyCfg Notify `yaml:"notify"`
}