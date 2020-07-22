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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/dhellmann/kubebuilder-experiment3/api/v1alpha1"
	studyv1alpha1 "github.com/dhellmann/kubebuilder-experiment3/api/v1alpha1"
)

// SimpleReconciler reconciles a Simple object
type SimpleReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=study.kni-arch.org,resources=simples,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=study.kni-arch.org,resources=simples/status,verbs=get;update;patch

// Reconcile is called when a resource is updated
func (r *SimpleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("simple", req.NamespacedName)

	obj := &v1alpha1.Simple{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, obj)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			// Request object not found, could have been deleted after
			// reconcile request.  Owned objects are automatically
			// garbage collected. For additional cleanup logic use
			// finalizers.  Return and don't requeue
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, errors.Wrap(err, "could not load resource")
	}

	log.Info("reconciling", "value", obj.Spec.Foo)

	return ctrl.Result{}, nil
}

// SetupWithManager registers the controller with the manager
func (r *SimpleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&studyv1alpha1.Simple{}).
		Complete(r)
}
