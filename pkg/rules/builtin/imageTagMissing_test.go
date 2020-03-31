package builtin

import (
	"github.com/madvikinggod/kubelint/pkg/rules/builtin/testdata"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func TestImageTagMissingCheckDeployment(t *testing.T) {
	type args struct {
		obj *unstructured.Unstructured
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
			name: "Should fail if Image tag is empty",
			args: args{testdata.NoImageTagUnstructured()},
			want: &want{
				message:    "ImageTagMissing - No tag detected on contaner noImageTagContainer - apps/v1/Deployment noImageTag/noImageTag",
				isCritical: false,
			},
		},
		{
			name: "Should fail if Image tag is latest",
			args: args{testdata.LatestImageTagUnstructured()},
			want: &want{
				message:    "ImageTagMissing - 'latest' tag on contaner latestImageTagContainer - apps/v1/Deployment latestImageTag/latestImageTag",
				isCritical: false,
			},
		},
		{
			name: "Should pass if Image has a tag",
			args: args{testdata.HasImageTagUnstructured()},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ImageTagMissingCheckDeployment(tt.args.obj)
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
