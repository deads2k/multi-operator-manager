// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ImageDigestMirrorSetSpecApplyConfiguration represents a declarative configuration of the ImageDigestMirrorSetSpec type for use
// with apply.
type ImageDigestMirrorSetSpecApplyConfiguration struct {
	ImageDigestMirrors []ImageDigestMirrorsApplyConfiguration `json:"imageDigestMirrors,omitempty"`
}

// ImageDigestMirrorSetSpecApplyConfiguration constructs a declarative configuration of the ImageDigestMirrorSetSpec type for use with
// apply.
func ImageDigestMirrorSetSpec() *ImageDigestMirrorSetSpecApplyConfiguration {
	return &ImageDigestMirrorSetSpecApplyConfiguration{}
}

// WithImageDigestMirrors adds the given value to the ImageDigestMirrors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImageDigestMirrors field.
func (b *ImageDigestMirrorSetSpecApplyConfiguration) WithImageDigestMirrors(values ...*ImageDigestMirrorsApplyConfiguration) *ImageDigestMirrorSetSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithImageDigestMirrors")
		}
		b.ImageDigestMirrors = append(b.ImageDigestMirrors, *values[i])
	}
	return b
}