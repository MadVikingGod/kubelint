package builtin

import (
	"testing"

	"github.com/madvikinggod/kubelint/pkg/rules/builtin/testdata"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func TestPodResourcesCheck(t *testing.T) {
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
			name: "Should fail if there are no pod resources",
			args: args{yaml: testdata.NoResources},
			want: &want{
				message:    "PodResourcesCheck - container noResourceContainer does not have resources - {noResource noResource apps/v1 Deployment}",
				isCritical: true,
			},
		},
		{
			name: "Should pass if there are pod resources",
			args: args{yaml: testdata.HasResources},
			want: nil,
		},
		{
			name: "Should fail if there are no pod resources in initContainers",
			args: args{yaml: testdata.NoResourcesInitContainers},
			want: &want{
				message:    "PodResourcesCheck - container noResourceInitContainer does not have resources - {noResourceInit noResourceInit apps/v1 Deployment}",
				isCritical: true,
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj, _ := yaml.Parse(tt.args.yaml)
			meta, _ := obj.GetMeta()

			got := PodResourcesCheck(obj, meta.GetIdentifier())

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
