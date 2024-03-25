package servicemodels

import "github.com/sbarter/sbarter_be_base_examples/sbarternetwork"

type GetVersionResponse struct {
	sbarternetwork.BaseDTO
	ApplicationInfo *ApplicationInfo
}
