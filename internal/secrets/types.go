package secrets

// Application secrets
type Secrets struct {
	AppPort   string   `json:"app_port" validate:"required"`
	Database  Database `json:"database" validate:"required"`
	JWTSecret string   `json:"jwt_secret" validate:"required"`
}

// Database secrets
type Database struct {
	Host string `json:"host" validate:"required"`
	Port string `json:"port" validate:"required"`
	User string `json:"user" validate:"required"`
	Pass string `json:"pass" validate:"required"`
	Name string `json:"name" validate:"required"`
}
