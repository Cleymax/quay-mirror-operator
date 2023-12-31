/*
Copyright 2023.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	clementperrinfrv1alpha1 "github.com/cleymax/quay-mirror-operator/api/v1alpha1"
)

// QuayMirrorRepositoryReconciler reconciles a QuayMirrorRepository object
type QuayMirrorRepositoryReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=clementperrin.fr,resources=quaymirrorrepositories,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=clementperrin.fr,resources=quaymirrorrepositories/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=clementperrin.fr,resources=quaymirrorrepositories/finalizers,verbs=update

func (r *QuayMirrorRepositoryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *QuayMirrorRepositoryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&clementperrinfrv1alpha1.QuayMirrorRepository{}).
		Complete(r)
}
