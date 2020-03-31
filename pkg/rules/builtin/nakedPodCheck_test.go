package builtin

import (
	"sigs.k8s.io/kustomize/kyaml/yaml"
	"testing"
)

func TestNakedPodCheck(t *testing.T) {
	type args struct {
		yaml string
	}
	type want struct {
		message    string
		isCritical bool
	}
	tests := []struct {
		name string
		args args
		want *want
	}{
		{
			name: "Pods should always fail",
			args: args{
				yaml: `apiVersion: v1
kind: Pod
metadata:
  name: test-pod
  namespace: pod-namespace
`,
			},
			want: &want{
				message:    "NakedPodCheck - Pods should not be used directly, apps/v1 Deployments are recommended - {test-pod pod-namespace v1 Pod}",
				isCritical: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj, _ := yaml.Parse(tt.args.yaml)
			meta, _ := obj.GetMeta()

			got := NakedPodCheck(obj, meta.GetIdentifier())

			if (tt.want == nil) && (got == nil) {
				return
			}
			if (tt.want == nil) != (got == nil) {
				t.Fatalf("Check() = %t, want %t", tt.want == nil, got == nil)

			}
			if got.String() != tt.want.message {
				t.Errorf("Check().String() = %v, want %v", got.String(), tt.want.message)
			}
			if got.IsCritical() != tt.want.isCritical {
				t.Errorf("Check().IsCritical() = %v, want %v", got.IsCritical(), tt.want.isCritical)

			}
		})
	}
}
