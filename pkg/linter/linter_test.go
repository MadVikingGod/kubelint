package linter

import (
	"github.com/madvikinggod/kubelint/pkg/rules/builtin"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func Test_linter_lintObjects(t *testing.T) {

	type msg struct {
		message string
		isCritical bool
	}
	tests := []struct {
		name string
		objs []*unstructured.Unstructured
		want []msg
	}{
		{
			name: "No Objects should return no messages",
			want: []msg{},
		},
		{
			name: "Unknown object should return not found warning",
			objs: []*unstructured.Unstructured{
				&unstructured.Unstructured{Object: map[string]interface{}{
					"kind": "Carp",
					"metadata": map[string]interface{}{
						"name": "bob",
					},
				}},
			},
			want: []msg{
				{
					message: "NoRuleFound - Did not find any rules for - /Carp /bob",
					isCritical: false,
				},
			},
		},
		{
			name: "Pod should return NakedPodCheck message",
			objs: []*unstructured.Unstructured{
				&unstructured.Unstructured{Object: map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "pod",
					"metadata": map[string]interface{}{
						"name": "pod-name",
					},
				}},
			},
			want: []msg{
				{
					message: "NakedPodCheck - Pods should not be used directly. Use a deployment instead - v1/pod /pod-name",
					isCritical: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &linter{
				cfg: LinterConfig{Rules: builtin.DefaultRules},
			}
			got := l.lintObjects(tt.objs)
			if len(got) != len(tt.want) {
				t.Errorf("lintObjects() returned %d items, wanted %d, got = %v, wanted = %v", len(got), len(tt.want), got, tt.want)
				return
			}
			for i := range got {
				if got[i].String() != tt.want[i].message {
					t.Errorf("Message missmatch: got = %s, want = %s", got[i].String() , tt.want[i].message)
				}
				if got[i].IsCritical() != tt.want[i].isCritical {
					t.Errorf("Message critical: got = %t, want = %t", got[i].IsCritical() , tt.want[i].isCritical)
				}
			}
		})
	}
}
