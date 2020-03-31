package builtin

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestNakedPodCheck_Check(t *testing.T) {
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
		want want
	}{
		{
			name: "pods should return msg",
			args: args{
				obj: &unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "pod",
						"metadata": map[string]interface{}{
							"name":      "test-pod-name",
							"namespace": "test-pod-namespace",
						},
					},
				},
			},
			want: want{
				message:    fmt.Sprintf("NakedPodCheck - %s - v1/pod test-pod-namespace/test-pod-name", NakedPodCheckStr),
				isCritical: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NakedPodCheck(tt.args.obj)
			if got.String() != tt.want.message {
				t.Errorf("Check().String() = %v, want %v", got.String(), tt.want.message)
			}
			if got.IsCritical() != tt.want.isCritical {
				t.Errorf("Check().IsCritical() = %v, want %v", got.IsCritical(), tt.want.isCritical)

			}
		})
	}
}
