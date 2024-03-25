package handlers

import (
	"github.com/sbarter/sbarter_be_base_examples/sbarternetwork"
	servicemodels "github.com/sbarter/sbarter_be_base_examples/sbarterservice/models"
	config "github.com/spf13/viper"
)

// BuildHandler
type BuildHandler struct {
	buildDate string
	buildHash string
}

// GetVersion returns service version to be pinged by infrastructure
func (handler *BuildHandler) GetVersion(request sbarternetwork.BaseDTO, response *servicemodels.GetVersionResponse) error {

	// Set application info
	applicationInfo := &servicemodels.ApplicationInfo{
		Name:        config.GetString("app_name"),
		Version:     config.GetString("version"),
		Service:     config.GetString("service_name"),
		Environment: config.GetString("stage"),
		CommitDate:  config.GetString("commit_date"),
		BuildDate:   handler.buildDate,
		BuildHash:   handler.buildHash,
	}

	// Set response
	response.ApplicationInfo = applicationInfo

	return nil
}
