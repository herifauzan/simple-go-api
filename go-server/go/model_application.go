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

type Application struct {

	// date in format YYYY-MM-DD
	Date string `json:"date,omitempty"`

	// ktp is a string of number (0..9) with length 16
	Ktp string `json:"ktp,omitempty"`

	// birth date in format YYYY-MM-DD
	Birthdate string `json:"birthdate,omitempty"`

	// gender value limited only for 'male' or 'female'
	Gender string `json:"gender,omitempty"`

	// amount in Rupiah with only number character allowed (0..9)
	Amount string `json:"amount,omitempty"`

	// period in month
	Period int32 `json:"period,omitempty"`
}
