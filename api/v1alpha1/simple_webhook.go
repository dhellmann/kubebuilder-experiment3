/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var simplelog = logf.Log.WithName("simple-resource")

func (r *Simple) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-study-kni-arch-org-v1alpha1-simple,mutating=true,failurePolicy=fail,groups=study.kni-arch.org,resources=simples,verbs=create;update,versions=v1alpha1,name=msimple.kb.io

var _ webhook.Defaulter = &Simple{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Simple) Default() {
	simplelog.Info("default", "name", r.Name)

	if r.Spec.Foo == "" {
		r.Spec.Foo = "a good default"
	}

}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-study-kni-arch-org-v1alpha1-simple,mutating=false,failurePolicy=fail,groups=study.kni-arch.org,resources=simples,versions=v1alpha1,name=vsimple.kb.io

var _ webhook.Validator = &Simple{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Simple) ValidateCreate() error {
	simplelog.Info("validate create", "name", r.Name)

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Simple) ValidateUpdate(old runtime.Object) error {
	simplelog.Info("validate update", "name", r.Name)

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Simple) ValidateDelete() error {
	simplelog.Info("validate delete", "name", r.Name)

	return nil
}
