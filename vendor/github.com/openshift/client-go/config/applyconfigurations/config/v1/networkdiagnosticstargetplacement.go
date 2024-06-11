// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// NetworkDiagnosticsTargetPlacementApplyConfiguration represents an declarative configuration of the NetworkDiagnosticsTargetPlacement type for use
// with apply.
type NetworkDiagnosticsTargetPlacementApplyConfiguration struct {
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	Tolerations  []v1.Toleration   `json:"tolerations,omitempty"`
}

// NetworkDiagnosticsTargetPlacementApplyConfiguration constructs an declarative configuration of the NetworkDiagnosticsTargetPlacement type for use with
// apply.
func NetworkDiagnosticsTargetPlacement() *NetworkDiagnosticsTargetPlacementApplyConfiguration {
	return &NetworkDiagnosticsTargetPlacementApplyConfiguration{}
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *NetworkDiagnosticsTargetPlacementApplyConfiguration) WithNodeSelector(entries map[string]string) *NetworkDiagnosticsTargetPlacementApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *NetworkDiagnosticsTargetPlacementApplyConfiguration) WithTolerations(values ...v1.Toleration) *NetworkDiagnosticsTargetPlacementApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}
