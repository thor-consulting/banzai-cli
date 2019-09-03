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

type GetDeploymentResponse struct {
	ReleaseName string `json:"releaseName,omitempty"`
	Chart string `json:"chart,omitempty"`
	ChartName string `json:"chartName,omitempty"`
	ChartVersion string `json:"chartVersion,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Version int32 `json:"version,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Status string `json:"status,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	// deployment notes in base64 encoded format
	Notes string `json:"notes,omitempty"`
	// current values of the deployment
	Values map[string]interface{} `json:"values,omitempty"`
}
