package spinnakerservice

import (
	spinnakerv1alpha1 "github.com/armory-io/spinnaker-operator/pkg/apis/spinnaker/v1alpha1"
	"github.com/armory-io/spinnaker-operator/pkg/halconfig"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Transformers tracks the list of transformers
var Transformers []TransformerGenerator

func init() {
	Transformers = append(Transformers, &ownerTransformerGenerator{})
}

// Transformer affects how Spinnaker is deployed.
// It can change the Spinnaker configuration itself with TransformConfig.
// It can also change the manifests before they are updated.
type Transformer interface {
	TransformConfig(hc *halconfig.SpinnakerCompleteConfig) error
	TransformManifests(scheme *runtime.Scheme, hc *halconfig.SpinnakerCompleteConfig, manifests []runtime.Object, status *spinnakerv1alpha1.SpinnakerServiceStatus) error
}

// TransformerGenerator generates transformers for the given SpinnakerService
type TransformerGenerator interface {
	NewTransformer(svc spinnakerv1alpha1.SpinnakerService, client client.Client) (Transformer, error)
}