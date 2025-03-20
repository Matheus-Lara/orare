package environment

import "github.com/Matheus-Lara/orare/pkg/common"

func IsDevelopment() bool {
	return common.GetEnv("GIN_MODE") == "debug" && common.GetEnv("HTTP_SERVER_HANDLER") == "default"
}
