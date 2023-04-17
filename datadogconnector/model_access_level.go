/*
 * Opal Custom App Connector API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datadogconnector

type AccessLevel struct {

	// The id of the access level in your system. Opal will provide this id when making requests for the access level to your connector.
	Id string `json:"id"`

	// The name of the access level
	Name string `json:"name"`
}
