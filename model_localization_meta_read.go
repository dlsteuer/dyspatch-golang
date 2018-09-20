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

type LocalizationMetaRead struct {
	// An opaque, unique identifier for a localization
	Id string `json:"id,omitempty"`
	// A language identifier comprised of a language and a country identifier.  See [supported languages](https://docs.dyspatch.io/localization/supported_languages/). 
	Language string `json:"language,omitempty"`
	// The user-specified name of a localization
	Name string `json:"name,omitempty"`
	// The API url for a specific localization
	Url string `json:"url,omitempty"`
}
