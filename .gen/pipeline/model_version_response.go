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

// Pipeline build and deployment info
type VersionResponse struct {
	Version string `json:"version,omitempty"`
	CommitHash string `json:"commit_hash,omitempty"`
	BuildDate string `json:"build_date,omitempty"`
	GoVersion string `json:"go_version,omitempty"`
	Os string `json:"os,omitempty"`
	Arch string `json:"arch,omitempty"`
	Compiler string `json:"compiler,omitempty"`
	InstanceUuid string `json:"instance_uuid,omitempty"`
}
