/*
Copyright 2021.

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
	"fmt"
	"github.com/go-logr/logr"

	batchv1 "github.com/golanshabi/scrambler/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Scrambler interface defines an interface used to scramble words.
type Scrambler func(x string) (string, error)

// ScrambledReconciler reconciles a Scrambled object.
type ScrambledReconciler struct {
	client.Client
	Log       logr.Logger
	Scheme    *runtime.Scheme
	Scrambler Scrambler
}

//+kubebuilder:rbac:groups=batch.github.com.golanshabi,resources=scrambleds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch.github.com.golanshabi,resources=scrambleds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=batch.github.com.golanshabi,resources=scrambleds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// the Scrambled object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *ScrambledReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("starting to scramble.")

	var scramble batchv1.Scrambled

	if err := r.Get(ctx, req.NamespacedName, &scramble); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	r.Log.Info(fmt.Sprintf("found resource, scrambling word - %s", scramble.Spec.CleanWord))

	scrambledString, err := r.Scrambler(scramble.Spec.CleanWord)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("error scrambling - %w", err)
	}

	scramble.Status.ScrambledWord = scrambledString
	err = r.Status().Update(ctx, &scramble)
	if err != nil {
		return ctrl.Result{}, err
	}

	r.Log.Info(fmt.Sprintf("scrambled word %s to %s", scramble.Spec.CleanWord, scrambledString))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ScrambledReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.Scrambled{}).
		Complete(r)
}
