package config

type (
	Config struct {
		Databases   []Database `yaml:"databases"`
		Translator  Translator `yaml:"translator"`
		Logging     Logging    `yaml:"logging"`
		BotToken    string     `yaml:"bot_token"`
		DeveloperID string     `yaml:"developer_id"`
		Debug       bool       `yaml:"debug"`
		PWD         string     `yaml:"pwd"`
	}

	Database struct {
		Name     string `yaml:"name"`
		Type     string `yaml:"type"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
		TimeZone string `yaml:"time_zone"`
		Charset  string `yaml:"charset"`
	}

	Translator struct {
		Path string `yaml:"path"`
	}

	Logging struct {
		Path         string `yaml:"path"`
		Pattern      string `yaml:"pattern"`
		MaxAge       string `yaml:"max_age"`
		RotationTime string `yaml:"rotation_time"`
		RotationSize string `yaml:"rotation_size"`
	}
)
