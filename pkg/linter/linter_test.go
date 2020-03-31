package linter

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

//func Test_linter_lintObjects(t *testing.T) {
//
//	type msg struct {
//		message    string
//		isCritical bool
//	}
//	tests := []struct {
//		name string
//		objs []*unstructured.Unstructured
//		want []msg
//	}{
//		{
//			name: "No Objects should return no messages",
//			want: []msg{},
//		},
//		{
//			name: "Unknown object should return not found warning",
//			objs: []*unstructured.Unstructured{
//				&unstructured.Unstructured{Object: map[string]interface{}{
//					"kind": "Carp",
//					"metadata": map[string]interface{}{
//						"name": "bob",
//					},
//				}},
//			},
//			want: []msg{
//				{
//					message:    "NoRuleFound - Did not find any rules for - /Carp /bob",
//					isCritical: false,
//				},
//			},
//		},
//		{
//			name: "Pod should return NakedPodCheck message",
//			objs: []*unstructured.Unstructured{
//				&unstructured.Unstructured{Object: map[string]interface{}{
//					"apiVersion": "v1",
//					"kind":       "pod",
//					"metadata": map[string]interface{}{
//						"name": "pod-name",
//					},
//				}},
//			},
//			want: []msg{
//				{
//					message:    "NakedPodCheck - Pods should not be used directly. Use a deployment instead - v1/pod /pod-name",
//					isCritical: true,
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			l := &linter{
//				cfg: Config{Rules: builtin.DefaultRules},
//			}
//			got := l.lintObjects(tt.objs)
//			if len(got) != len(tt.want) {
//				t.Errorf("lintObjects() returned %d items, wanted %d, got = %v, wanted = %v", len(got), len(tt.want), got, tt.want)
//				return
//			}
//			for i := range got {
//				if got[i].String() != tt.want[i].message {
//					t.Errorf("Message missmatch: got = %s, want = %s", got[i].String(), tt.want[i].message)
//				}
//				if got[i].IsCritical() != tt.want[i].isCritical {
//					t.Errorf("Message critical: got = %t, want = %t", got[i].IsCritical(), tt.want[i].isCritical)
//				}
//			}
//		})
//	}
//}

func Test_readObjects(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{
			name: "Can read One yaml file",
			args: args{
				reader: strings.NewReader(`apiVersion: v1
kind: test`),
			},
			wantLen: 1,
			wantErr: false,
		},
		{
			name: "Can read One yaml file starting with ---",
			args: args{reader: strings.NewReader(`---
apiVersion: v1
kind: test`)},
			wantLen: 1,
			wantErr: false,
		},
		{
			name: "can read multiple yaml files",
			args: args{reader: strings.NewReader(`apiVersion: v1
kind: test
---
apiVersion: v2
kind: otherTest`)},
			wantLen: 2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readObjects(tt.args.reader)
			for _, n := range got {
				fmt.Println(n.MustString())
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("readObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("readObjects() len(got) = %d, want %v", len(got), tt.wantLen)
			}
		})
	}
}
