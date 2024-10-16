// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// RegistryLocationApplyConfiguration represents a declarative configuration of the RegistryLocation type for use
// with apply.
type RegistryLocationApplyConfiguration struct {
	DomainName *string `json:"domainName,omitempty"`
	Insecure   *bool   `json:"insecure,omitempty"`
}

// RegistryLocationApplyConfiguration constructs a declarative configuration of the RegistryLocation type for use with
// apply.
func RegistryLocation() *RegistryLocationApplyConfiguration {
	return &RegistryLocationApplyConfiguration{}
}

// WithDomainName sets the DomainName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DomainName field is set to the value of the last call.
func (b *RegistryLocationApplyConfiguration) WithDomainName(value string) *RegistryLocationApplyConfiguration {
	b.DomainName = &value
	return b
}

// WithInsecure sets the Insecure field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Insecure field is set to the value of the last call.
func (b *RegistryLocationApplyConfiguration) WithInsecure(value bool) *RegistryLocationApplyConfiguration {
	b.Insecure = &value
	return b
}