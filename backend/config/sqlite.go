package config

type SQLiteConfig struct {
	Path string
}

func GetSQLiteConfig() SQLiteConfig {
	return SQLiteConfig{
		Path: "./db.db",
	}
}
