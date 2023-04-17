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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// GenerateSignature generates a signature for a given payload in a HTTP request
func GenerateSignature(
	signingSecret string,
	timestamp string,
	serializedBlob []byte,
) (string, error) {
	// Concatenate base string
	sigBaseString := "v0:" + timestamp + ":" + string(serializedBlob)

	// Hash base string to get signature
	hash := hmac.New(sha256.New, []byte(signingSecret))
	_, err := hash.Write([]byte(sigBaseString))
	if err != nil {
		return "", errors.Wrap(err, "error writing hash")
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func validateOpalSignature(signingSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		opalSignature := c.GetHeader("X-Opal-Signature")
		if opalSignature == "" {
			c.JSON(401, Error{
				Code:    401,
				Message: "X-Opal-Signature header is missing",
			})
			c.Abort()
			return
		}
		opalRequestTimestamp := c.GetHeader("X-Opal-Request-Timestamp")
		if opalRequestTimestamp == "" {
			c.JSON(401, Error{
				Code:    401,
				Message: "X-Opal-Request-Timestamp header is missing",
			})
			c.Abort()
			return
		}
		serializedBlob, err := json.Marshal(c.Request.Body)
		if err != nil {
			c.JSON(500, Error{
				Code:    500,
				Message: "Unable to serialize payload",
			})
			c.Abort()
			return
		}

		signature, err := GenerateSignature(signingSecret, opalRequestTimestamp, serializedBlob)
		fmt.Printf("signature %s timestamp %s body %s signature %s", opalSignature, opalRequestTimestamp, string(serializedBlob), signingSecret)
		if signature != opalSignature || err != nil {
			c.JSON(401, Error{
				Code:    401,
				Message: "Invalid signature",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(validateOpalSignature(os.Getenv("OPAL_SIGNING_SECRET")))

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

var routes = Routes{

	{
		"AddResourceUser",
		http.MethodPost,
		"/resources/:resource_id/users",
		AddResourceUser,
	},

	{
		"GetResource",
		http.MethodGet,
		"/resources/:resource_id",
		GetResource,
	},

	{
		"GetResourceAccessLevels",
		http.MethodGet,
		"/resources/:resource_id/access_levels",
		GetResourceAccessLevels,
	},

	{
		"GetResourceUsers",
		http.MethodGet,
		"/resources/:resource_id/users",
		GetResourceUsers,
	},

	{
		"GetResources",
		http.MethodGet,
		"/resources",
		GetResources,
	},

	{
		"RemoveResourceUser",
		http.MethodDelete,
		"/resources/:resource_id/users/:user_id",
		RemoveResourceUser,
	},

	{
		"GetStatus",
		http.MethodGet,
		"/status",
		GetStatus,
	},

	{
		"GetUsers",
		http.MethodGet,
		"/users",
		GetUsers,
	},
}