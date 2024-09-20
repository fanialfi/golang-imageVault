package config

import "os"

type config struct {
	dbProtocol string
	dbPasswd   string
	dbHost     string
	dbName     string
	dbUser     string
	dbPort     string
	port       string
}

// LoadConfig akan meload semua envoronment variabel yang diperlukan ke struct
func LoadConfig() (*config, error) {
	var cfg config

	cfg.SetDbHost(getEnv("DB_HOST", ""))
	cfg.SetDbPort(getEnv("DB_PORT", "3306"))
	cfg.SetDbName(getEnv("DB_NAME", ""))
	cfg.SetDbPasswd(getEnv("DB_PASSWD", ""))
	cfg.SetDbUser(getEnv("DB_USER", "root"))
	cfg.SetPort(getEnv("PORT", "8080"))
	cfg.SetDbProtocol(getEnv("DB_PROTOCOL", "tcp"))

	return &cfg, nil
}

// methof for set field struct dbProtocol
func (cfg *config) SetDbProtocol(key string) {
	cfg.dbProtocol = key
}

// method for get value of field struct dbProtocol
func (cfg *config) GetDbProtocol() *string {
	return &cfg.dbProtocol
}

// method for set field struct dbHost
func (cfg *config) SetDbHost(key string) {
	cfg.dbHost = key
}

// method for get value of field struct dbHost
func (cfg *config) GetDbHost() *string {
	return &cfg.dbHost
}

// method for set field struct dbName
func (cfg *config) SetDbName(key string) {
	cfg.dbName = key
}

// method for get value of field struct dbName
func (cfg *config) GetDbName() *string {
	return &cfg.dbName
}

// method for set field struct dbUser
func (cfg *config) SetDbUser(key string) {
	cfg.dbUser = key
}

// method for get value of field struct dbUser
func (cfg *config) GetDbUser() *string {
	return &cfg.dbUser
}

// method for set field struct dbPort
func (cfg *config) SetDbPort(key string) {
	cfg.dbPort = key
}

// method for get value of field struct dbPort
func (cfg *config) GetDbPort() *string {
	return &cfg.dbPort
}

// method for set field struct port
func (cfg *config) SetPort(key string) {
	cfg.port = key
}

// method for get value of field struct port
func (cfg *config) GetPort() *string {
	return &cfg.port
}

// method for set field struct dbPasswd
func (cfg *config) SetDbPasswd(key string) {
	cfg.dbPasswd = key
}

// method for get value of field struct dbPasswd
func (cfg *config) GetDbPasswd() *string {
	return &cfg.dbPasswd
}

// getEnv akan mengambil nilai dari environment variabel berdasarkan key
// jika tidak ada maka akan mengembalikan defaultValue
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
