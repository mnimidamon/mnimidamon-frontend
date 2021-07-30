package global

import "mnimidamonbackend/models"

// General application config.
type Config struct {
	User     *UserConfig
	Computer *ComputerConfig
	Server   *ServerConfig
}

// Logged in user.
// JWT authentication key for current user.
type UserConfig struct {
	models.User
	Key string
}

// Logged in computer.
// JWT authentication key for current computer.
type ComputerConfig struct {
	models.Computer
	Key string
}

// Server settings.
type ServerConfig struct {
	ServerAddress string
	FolderPath    string
}
