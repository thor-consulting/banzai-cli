/*
 * Cluster Recommender.
 *
 * This project can be used to recommend instance type groups on different cloud providers consisting of regular and spot/preemptible instances. The main goal is to provide and continuously manage a cost-effective but still stable cluster layout that's built up from a diverse set of regular and spot instances.
 *
 * API version: 0.5.3
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package telescopes
// RecommendClusterScaleOutRequest ClusterScaleoutRecommendationReq encapsulates the recommendation input data
type RecommendClusterScaleOutRequest struct {
	// Description of the current cluster layout in:body
	ActualLayout []NodePoolDesc `json:"actualLayout,omitempty"`
	// Total desired number of CPUs in the cluster after the scale out
	DesiredCpu float64 `json:"desiredCpu,omitempty"`
	// Total desired number of GPUs in the cluster after the scale out
	DesiredGpu int64 `json:"desiredGpu,omitempty"`
	// Total desired memory (GB) in the cluster after the scale out
	DesiredMem float64 `json:"desiredMem,omitempty"`
	// Excludes is a blacklist - a slice with vm types to be excluded from the recommendation
	Excludes []string `json:"excludes,omitempty"`
	// Percentage of regular (on-demand) nodes among the scale out nodes
	OnDemandPct int64 `json:"onDemandPct,omitempty"`
	// Availability zone to be included in the recommendation
	Zone string `json:"zone,omitempty"`
}
