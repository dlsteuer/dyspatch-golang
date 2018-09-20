/*
 * Dyspatch API
 *
 * # Introduction  The Dyspatch API is based on the REST paradigm, and features resource based URLs with standard HTTP response codes to indicate errors. We use standard HTTP authentication and request verbs, and all responses are JSON formatted. See our [Implementation Guide](https://docs.dyspatch.io/development/implementing_dyspatch/) for more details on how to implement Dyspatch.  ## API Client Libraries  Dyspatch provides API Clients for popular languages and web frameworks.   - [Java](https://github.com/getdyspatch/dyspatch-java) - [Javascript](https://github.com/getdyspatch/dyspatch-javascript) - [Python](https://github.com/getdyspatch/dyspatch-python) 
 *
 * API version: 2018.08
 * Contact: support@dyspatch.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dyspatch_client
import (
	"time"
)

type TemplateMetaRead struct {
	// An opaque, unique identifier for a template
	Id string `json:"id,omitempty"`
	// The name of a template
	Name string `json:"name,omitempty"`
	// A description of the template
	Description string `json:"description,omitempty"`
	// The API url for a specific template
	Url string `json:"url,omitempty"`
	// A list of the template's available localization objects
	Localizations []LocalizationMetaRead `json:"localizations,omitempty"`
	// The time of initial creation
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// The time of last update
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}