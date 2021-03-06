/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
type V2beta1ObjectMetricStatus struct {

	// currentValue is the current value of the metric (as a quantity).
	CurrentValue string `json:"currentValue"`

	// metricName is the name of the metric in question.
	MetricName string `json:"metricName"`

	// target is the described Kubernetes object.
	Target *V2beta1CrossVersionObjectReference `json:"target"`
}
