package gorm

type PostgresConfig struct {
	Host     string `yaml:"postgres_host" env:"POSTGRES_HOST" env-default:"localhost"`
	User     string `yaml:"postgres_user" env:"POSTGRES_USER" env-default:"postgres"`
	Password string `yaml:"postgres_password" env:"POSTGRES_PASSWORD" env-default:"2006"`
	Database string `yaml:"postgres_database" env:"POSTGRES_DATABASE" env-default:"effective-api-data"`
	Port     string `yaml:"postgres_port" env:"POSTGRES_PORT" env-default:"5432"`
	SSLMode  string `yaml:"postgres_sslmode" env:"POSTGRES_SSLMODE" env-default:"disable"`
	Timezone string `yaml:"postgres_timezone" env:"POSTGRES_TIMEZONE" env-default:"Europe/Moscow"`
}

//TODO реализация подключения к горм
//// Connect to postgres database with gorm
//func Connect(cfg PostgresConfig) (*gorm.DB, error) {
//	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
//		cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port, cfg.SSLMode, cfg.Timezone)
//
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
//	if err != nil {
//		panic(err)
//	}
//	// Migrate the schema
//	db.AutoMigrate(&Product{})
//
//}
