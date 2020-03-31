package builtin

import (
	"sigs.k8s.io/kustomize/kyaml/yaml"
	"testing"
)

func Test_depractedAPICheck(t *testing.T) {
	type args struct {
		newVersion yaml.TypeMeta
		yaml       string
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
			name: "It should error with new version",
			args: args{
				newVersion: yaml.TypeMeta{"test", "test.io/v1"},
				yaml: `apiVersion: test.io/v1beta1
kind: test
metadata:
  name: test-name
  namespace: test-namespace`,
			},
			want: &want{
				message:    "DeprecatedAPICheck - {test test.io/v1beta1} should not be used, use {test test.io/v1} - {test-name test-namespace test.io/v1beta1 test}",
				isCritical: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check := depractedAPICheck(tt.args.newVersion)
			y, _ := yaml.Parse(tt.args.yaml)
			meta, _ := y.GetMeta()
			got := check(y, meta.GetIdentifier())

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
