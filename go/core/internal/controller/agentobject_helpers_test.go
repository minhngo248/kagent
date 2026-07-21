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

package controller

import (
	"testing"

	"github.com/kagent-dev/kagent/go/api/v1alpha2"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestSandboxAgentUsesWorkerPool(t *testing.T) {
	t.Parallel()

	defaultWorkerPool := types.NamespacedName{Namespace: "kagent", Name: "default-wp"}

	for _, tt := range []struct {
		name string
		sa   *v1alpha2.SandboxAgent
		wp   types.NamespacedName
		want bool
	}{
		{
			name: "explicit substrate workerpool matches in sandboxagent namespace",
			sa:   sandboxAgentForWorkerPool("hello", &v1alpha2.TypedLocalReference{Name: "custom-wp"}),
			wp:   types.NamespacedName{Namespace: "kagent", Name: "custom-wp"},
			want: true,
		},
		{
			name: "explicit substrate workerpool does not match other pool",
			sa:   sandboxAgentForWorkerPool("hello", &v1alpha2.TypedLocalReference{Name: "custom-wp"}),
			wp:   types.NamespacedName{Namespace: "kagent", Name: "other-wp"},
			want: false,
		},
		{
			name: "substrate sandboxagent without explicit ref depends on default workerpool",
			sa:   sandboxAgentForWorkerPool("hello", nil),
			wp:   defaultWorkerPool,
			want: true,
		},
		{
			name: "non substrate sandboxagent does not match default workerpool",
			sa: &v1alpha2.SandboxAgent{
				ObjectMeta: metav1.ObjectMeta{Namespace: "kagent", Name: "hello"},
				Spec:       v1alpha2.SandboxAgentSpec{Platform: v1alpha2.SandboxPlatformAgentSandbox},
			},
			wp:   defaultWorkerPool,
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, sandboxAgentUsesWorkerPool(tt.sa, tt.wp, defaultWorkerPool))
		})
	}
}

func sandboxAgentForWorkerPool(name string, ref *v1alpha2.TypedLocalReference) *v1alpha2.SandboxAgent {
	return &v1alpha2.SandboxAgent{
		ObjectMeta: metav1.ObjectMeta{Namespace: "kagent", Name: name},
		Spec: v1alpha2.SandboxAgentSpec{
			Platform: v1alpha2.SandboxPlatformSubstrate,
			Substrate: &v1alpha2.SandboxSubstrateSpec{
				WorkerPoolRef: ref,
			},
		},
	}
}
