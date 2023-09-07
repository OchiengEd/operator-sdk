// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package release

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Test_readBoolAnnotationWithDefault(t *testing.T) {
	objBuilder := func(anno map[string]string) *unstructured.Unstructured {
		object := &unstructured.Unstructured{}
		object.SetAnnotations(anno)
		return object
	}

	type args struct {
		obj        *unstructured.Unstructured
		annotation string
		fallback   bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return value of annotation read",
			args: args{
				obj: objBuilder(map[string]string{
					"helm.sdk.operatorframework.io/rollback-force": "false",
				}),
				annotation: "helm.sdk.operatorframework.io/rollback-force",
				fallback:   true,
			},
			want: false,
		},
		{
			name: "Should return fallback when annotation is not present",
			args: args{
				obj: objBuilder(map[string]string{
					"helm.sdk.operatorframework.io/upgrade-force": "true",
				}),
				annotation: "helm.sdk.operatorframework.io/rollback-force",
				fallback:   false,
			},
			want: false,
		},
		{
			name: "Should return fallback when errors while parsing bool value",
			args: args{
				obj: objBuilder(map[string]string{
					"helm.sdk.operatorframework.io/rollback-force": "force",
				}),
				annotation: "helm.sdk.operatorframework.io/rollback-force",
				fallback:   true,
			},
			want: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := readBoolAnnotationWithDefault(tc.args.obj, tc.args.annotation, tc.args.fallback); got != tc.want {
				assert.Equal(t, tc.want, got, "readBoolAnnotationWithDefault() function")
			}
		})
	}
}
