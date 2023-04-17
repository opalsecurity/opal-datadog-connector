/*
 * Opal Custom App Connector API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datadogconnector

type AddResourceUserRequest struct {

	// The ID of the user to add.
	UserId string `json:"user_id"`

	// The identifier of your custom app.
	AppId string `json:"app_id"`

	// The ID of the access level to assign to the user.
	AccessLevelId string `json:"access_level_id,omitempty"`
}