package builtin

import (
	"github.com/madvikinggod/kubelint/pkg/rules/builtin/testdata"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func TestImagePullPolicyAlways_Check(t *testing.T) {
	type want struct {
		message string
		isCritical bool
	}
	tests := []struct {
		name string
		args *unstructured.Unstructured
		want *want
	}{
		{
			name: "Should fail if Image pull policy is empty",
			args: testdata.NoImagePullPolicyUnstructured(),
			want: &want{
				message: "ImagePullPolicyAlways - container noImagePullPolicyContainer has an image pull policy of %!s(<nil>), it should be Alaways - apps/v1/Deployment noImagePullPolicy/noImagePullPolicy",
				isCritical: true,
			},
		},
		{
			name: "Should fail if Image pull policy is not always",
			args: testdata.NeverImagePullPolicyUnstructured(),
			want: &want{
				message: "ImagePullPolicyAlways - container neverImagePullPolicyContainer has an image pull policy of Never, it should be Alaways - apps/v1/Deployment neverImagePullPolicy/neverImagePullPolicy",
				isCritical:true,
			},
		},
		{
			name: "Should pass if Image pull policy is Always",
			args: testdata.AlwaysImagePullPolicyUnstructured(),
			want: nil,
		},
		{
			name: "Should fail if Image pull policy for any continer is empty",
			args: testdata.MultiImagePullPolicyUnstructured(),
			want: &want{
				message: "ImagePullPolicyAlways - container multiImagePullPolicyContainer-fail has an image pull policy of Never, it should be Alaways - apps/v1/Deployment multiImagePullPolicy/multiImagePullPolicy",
				isCritical:true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := ImagePullPolicyAlways{}
			got := i.Check(tt.args)
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