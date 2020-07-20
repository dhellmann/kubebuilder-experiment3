package v1alpha2

import (
	"github.com/dhellmann/kubebuilder-experiment3/api/v1alpha1"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var log = ctrl.Log.WithName("simple-conversion-v1alpha2")

// ConvertTo converts this Simple to the Hub version (v1alpha1).
func (src *Simple) ConvertTo(dstRaw conversion.Hub) error {

	dst := dstRaw.(*v1alpha1.Simple)

	log.Info("ConvertTo")

	// renamed field
	dst.Spec.Foo = src.Spec.LessSillyName

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha1) to this version.
func (dst *Simple) ConvertFrom(srcRaw conversion.Hub) error {

	src := srcRaw.(*v1alpha1.Simple)

	log.Info("ConvertFrom")

	// renamed field
	dst.Spec.LessSillyName = src.Spec.Foo

	return nil
}
