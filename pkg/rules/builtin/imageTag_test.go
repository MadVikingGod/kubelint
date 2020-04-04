package builtin

import (
	"testing"

	"github.com/madvikinggod/kubelint/pkg/rules/builtin/testdata"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func TestImageTagCheck(t *testing.T) {
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
			name: "Should fail if image tag is latest",
			args: args{yaml: testdata.LatestImageTagYaml},
			want: &want{
				message:    "ImageTagCheck - container latestImageTagContainer has an image tag of latest - {latestImageTag latestImageTag apps/v1 Deployment}",
				isCritical: true,
			},
		},
		{
			name: "Should fail if image tag is missing",
			args: args{yaml: testdata.NoImageTagYaml},
			want: &want{
				message:    "ImageTagCheck - container noImageTagContainer has no image tag - {noImageTag noImageTag apps/v1 Deployment}",
				isCritical: true,
			},
		},
		{
			name: "Should pass if image tag looks valid",
			args: args{yaml: testdata.HasImageTagYaml},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj, _ := yaml.Parse(tt.args.yaml)
			meta, _ := obj.GetMeta()

			got := ImageTagCheck(obj, meta.GetIdentifier())

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
