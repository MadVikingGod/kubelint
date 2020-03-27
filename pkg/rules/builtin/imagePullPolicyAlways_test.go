package builtin

import (
	"github.com/madvikinggod/kubelint/pkg/rules/builtin/testdata"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)




func TestImagePullPolicyAlwaysCheck(t *testing.T) {
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
			name: "Should fail if Image pull policy is empty",
			args: args{testdata.NoImagePullPolicyUnstructured()},
			want: &want{
				message:    "ImagePullPolicyAlways - container noImagePullPolicyContainer has an image pull policy of , it should be Alaways - apps/v1/Deployment noImagePullPolicy/noImagePullPolicy",
				isCritical: true,
			},
		},
		{
			name: "StatefulSets Should fail if Image pull policy is empty",
			args: args{testdata.SSNoImagePullPolicyUnstructured()},
			want: &want{
				message:    "ImagePullPolicyAlways - container noImagePullPolicyContainer has an image pull policy of , it should be Alaways - apps/v1/Deployment noImagePullPolicy/noImagePullPolicy",
				isCritical: true,
			},
		},
		{
			name: "Should fail if Image pull policy is not always",
			args: args{testdata.NeverImagePullPolicyUnstructured()},
			want: &want{
				message:    "ImagePullPolicyAlways - container neverImagePullPolicyContainer has an image pull policy of Never, it should be Alaways - apps/v1/Deployment neverImagePullPolicy/neverImagePullPolicy",
				isCritical: true,
			},
		},
		{
			name: "Should pass if Image pull policy is Always",
			args: args{testdata.ImagePullPolicyAlwaysUnstructured()},
			want: nil,
		},
		{
			name: "Should fail if Image pull policy for any continer is empty",
			args: args{testdata.ImagePullPolicyMultiUnstructured()},
			want: &want{
				message:    "ImagePullPolicyAlways - container multiImagePullPolicyContainer-fail has an image pull policy of Never, it should be Alaways - apps/v1/Deployment multiImagePullPolicy/multiImagePullPolicy",
				isCritical: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ImagePullPolicyAlwaysCheck(tt.args.obj)
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
