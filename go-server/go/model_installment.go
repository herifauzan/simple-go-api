/*
 * Loan Service API
 *
 * This is a simple server to create and get information about Loan service.
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Installment struct {

	// date in format YYYY-MM-DD
	Date string `json:"date,omitempty"`

	// amount in Rupiah with only number character allowed (0..9)
	Amount string `json:"amount,omitempty"`

	// period in month
	Period int32 `json:"period,omitempty"`
}
