package cmd

const (
	LOGIN   string = "login"
	UPDATE_PASSWORD  string = "update-pwd"
	DELETE  string = "delete"
	HISTORY string = "history"
	MIGRATE string = "migrate"
)

var operations_arr = []string{LOGIN, UPDATE_PASSWORD, DELETE, HISTORY, MIGRATE}
