/*
Copyright 2025.

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

package reconciler

import (
	"testing"

	"github.com/kagent-dev/kagent/go/api/v1alpha2"
	"github.com/kagent-dev/kagent/go/core/pkg/sandboxbackend"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAcceptedConditionForErrorUsesDependencyPendingReason(t *testing.T) {
	err := &sandboxbackend.DependencyPendingError{
		Reason:  "WorkerPoolNotFound",
		Message: "Waiting for WorkerPool kagent/custom-wp",
	}
	sa := &v1alpha2.SandboxAgent{ObjectMeta: metav1.ObjectMeta{Name: "hello", Namespace: "kagent", Generation: 3}}

	condition := acceptedConditionForError(sa, err)

	require.Equal(t, metav1.ConditionFalse, condition.Status)
	require.Equal(t, "WorkerPoolNotFound", condition.Reason)
	require.Equal(t, "Waiting for WorkerPool kagent/custom-wp", condition.Message)
	require.Equal(t, int64(3), condition.ObservedGeneration)
}
