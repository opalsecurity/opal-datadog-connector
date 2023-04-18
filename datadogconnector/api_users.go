/*
 * Opal Custom App Connector API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datadogconnector

import (
	"fmt"
	"math"
	"net/http"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/gin-gonic/gin"
)

const (
	userPageSize = int64(5_000) // maximum allowed number by Datadog API
)

// GetUsers -
func GetUsers(c *gin.Context) {
	ctx := datadog.NewDefaultContext(c.Request.Context())
	apiClient := GetDatadogClient()
	api := datadogV2.NewUsersApi(apiClient)

	var pageNumber int64 = 0
	if encodedCursor := c.Query("cursor"); encodedCursor != "" {
		cursor, err := decodeCursor(encodedCursor)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &Error{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("Invalid cursor: %s", err.Error()),
			})
			return
		}
		pageNumber = int64(cursor.NextPage)
	}

	usersResponse, resp, err := api.ListUsers(ctx, datadogV2.ListUsersOptionalParameters{
		PageNumber: &pageNumber,
		PageSize:   datadog.PtrInt64(userPageSize),
	})
	if err != nil {
		c.AbortWithStatusJSON(resp.StatusCode, &Error{
			Code:    int32(resp.StatusCode),
			Message: fmt.Sprintf("Error listing users: %s", err.Error()),
		})
		return
	}

	var users []User
	for _, user := range usersResponse.GetData() {
		var email string
		if user.GetAttributes().Email != nil {
			email = *user.GetAttributes().Email
		}
		users = append(users, User{
			Id:    user.GetId(),
			Email: email,
		})
	}

	var nextCursor *string
	if page, ok := usersResponse.Meta.GetPageOk(); ok {
		if totalUsers := page.TotalFilteredCount; totalUsers != nil && *totalUsers > userPageSize {
			totalPages := math.Ceil(float64(*totalUsers) / float64(userPageSize))
			if pageNumber < int64(totalPages)-1 {
				nextPage := pageNumber + 1
				encodedCursor, err := encodeCursor(&Cursor{
					NextPage: int(nextPage),
				})
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, &Error{
						Code:    http.StatusInternalServerError,
						Message: fmt.Sprintf("Error encoding cursor: %s", err.Error()),
					})
					return
				}
				nextCursor = &encodedCursor
			}
		}
	}

	c.JSON(http.StatusOK, UsersResponse{
		Users:      users,
		NextCursor: nextCursor,
	})
}
