/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.29.1
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type GetClusterBootstrapResponse struct {
	Token string `json:"token,omitempty"`
	DiscoveryTokenCaCertHash string `json:"discoveryTokenCaCertHash,omitempty"`
	MasterAddress string `json:"masterAddress,omitempty"`
}
