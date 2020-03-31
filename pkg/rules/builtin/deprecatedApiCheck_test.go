package builtin

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func Test_depractedAPICheck(t *testing.T) {
	type args struct {
		newVersion string
		obj        *unstructured.Unstructured
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
			name: "It should fail when passed an object",
			args: args{
				newVersion: "fake.v1/stuff",
				obj: &unstructured.Unstructured{Object: map[string]interface{}{
					"apiVersion": "fake.v1beta1",
					"kind":       "stuff",
					"metadata": map[string]interface{}{
						"name":      "fake-stuff",
						"namespace": "fake-namespace",
					},
				}},
			},
			want: &want{
				message:    "DeprecatedAPICheck - fake.v1beta1/stuff should not be used, use fake.v1/stuff - fake.v1beta1/stuff fake-namespace/fake-stuff",
				isCritical: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := depractedAPICheck(tt.args.newVersion)
			got := rule(tt.args.obj)
			if (tt.want == nil) && (got == nil) {
				return
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
