// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	operatorv1alpha1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1alpha1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// OLMsGetter has a method to return a OLMInterface.
// A group's client should implement this interface.
type OLMsGetter interface {
	OLMs() OLMInterface
}

// OLMInterface has methods to work with OLM resources.
type OLMInterface interface {
	Create(ctx context.Context, oLM *v1alpha1.OLM, opts v1.CreateOptions) (*v1alpha1.OLM, error)
	Update(ctx context.Context, oLM *v1alpha1.OLM, opts v1.UpdateOptions) (*v1alpha1.OLM, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, oLM *v1alpha1.OLM, opts v1.UpdateOptions) (*v1alpha1.OLM, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OLM, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OLMList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OLM, err error)
	Apply(ctx context.Context, oLM *operatorv1alpha1.OLMApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.OLM, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, oLM *operatorv1alpha1.OLMApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.OLM, err error)
	OLMExpansion
}

// oLMs implements OLMInterface
type oLMs struct {
	*gentype.ClientWithListAndApply[*v1alpha1.OLM, *v1alpha1.OLMList, *operatorv1alpha1.OLMApplyConfiguration]
}

// newOLMs returns a OLMs
func newOLMs(c *OperatorV1alpha1Client) *oLMs {
	return &oLMs{
		gentype.NewClientWithListAndApply[*v1alpha1.OLM, *v1alpha1.OLMList, *operatorv1alpha1.OLMApplyConfiguration](
			"olms",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.OLM { return &v1alpha1.OLM{} },
			func() *v1alpha1.OLMList { return &v1alpha1.OLMList{} }),
	}
}