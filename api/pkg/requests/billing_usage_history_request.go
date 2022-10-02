package requests

import (
	"strings"

	"github.com/NdoleStudio/httpsms/pkg/repositories"
)

// BillingUsageHistory is the payload for fetching the entities.BillingUsage history
type BillingUsageHistory struct {
	request
	Skip  string `json:"skip" query:"skip"`
	Limit string `json:"limit" query:"limit"`
}

// Sanitize sets defaults to MessageOutstanding
func (input *BillingUsageHistory) Sanitize() BillingUsageHistory {
	if strings.TrimSpace(input.Limit) == "" {
		input.Limit = "12"
	}
	input.Skip = strings.TrimSpace(input.Skip)
	if input.Skip == "" {
		input.Skip = "0"
	}
	return *input
}

// ToIndexParams converts BillingUsageHistory to repositories.IndexParams
func (input *BillingUsageHistory) ToIndexParams() repositories.IndexParams {
	return repositories.IndexParams{
		Skip:  input.getInt(input.Skip),
		Limit: input.getInt(input.Limit),
	}
}
