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

type DeploymentListDeploymentResponse struct {
	Chart string `json:"chart,omitempty"`
	ChartName string `json:"chartName,omitempty"`
	ChartVersion string `json:"chartVersion,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	ReleaseName string `json:"releaseName,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Version int32 `json:"version,omitempty"`
}
