package env

// Data returns environment variables
func Data() map[string]string {
	return map[string]string{"PORT": "60020", "GIN_MODE": "debug"}
}
