package messages

import "golang.org/x/text/language"

func PtBr() map[string]string {
	return map[string]string{
		"language": language.BrazilianPortuguese.String(),

		"Admin.MigrateModels.Response.Success": "Models migradas com sucesso",
		"HealthCheck.Response.Success":         "Orare API em operação!",
	}
}
