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

type CreateUpdateDeploymentResponse struct {
	ReleaseName string `json:"releaseName,omitempty"`
	// deployment notes in base64 encoded format
	Notes string `json:"notes,omitempty"`
}
