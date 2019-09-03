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

type NodePoolStatusGoogle struct {
	Autoscaling bool `json:"autoscaling,omitempty"`
	Count int32 `json:"count,omitempty"`
	MinCount int32 `json:"minCount,omitempty"`
	MaxCount int32 `json:"maxCount,omitempty"`
	InstanceType string `json:"instanceType,omitempty"`
}
