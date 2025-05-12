// SPDX-FileCopyrightText: Copyright DB InfraGO AG and contributors
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	fnapi "github.com/crossplane/function-sdk-go/proto/v1"
	"github.com/crossplane/function-sdk-go/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func convertResultsToMap(r []*fnapi.Result) []map[string]interface{} {
	if r == nil {
		return nil
	}
	res := make([]map[string]interface{}, len(r))
	for i, rr := range r {
		if rr == nil {
			continue
		}
		res[i] = map[string]interface{}{
			"Message":  rr.GetMessage(),
			"Severity": rr.GetSeverity(),
		}
	}
	return res
}

func convertResourcesMapToUnstructured(r map[string]*fnapi.Resource) map[string]*unstructured.Unstructured {
	if r == nil {
		return nil
	}
	res := map[string]*unstructured.Unstructured{}
	for k, v := range r {
		u := convertResourceToUnstructured(v)

		// If the name annotation was the only annotation in the resource,
		// delete the entire field to avoid creating unnecessary diffs.
		if len(u.GetAnnotations()) == 0 {
			u.SetAnnotations(nil)
		}

		res[k] = u
	}
	return res
}

func convertResourceToUnstructured(r *fnapi.Resource) *unstructured.Unstructured {
	if r == nil {
		return nil
	}
	u := &unstructured.Unstructured{}
	if err := resource.AsObject(r.GetResource(), u); err != nil {
		panic(err)
	}
	return u
}

func ConvertDesiredCompositeToObject(r *fnapi.RunFunctionResponse, o runtime.Object) {
	if err := resource.AsObject(r.GetDesired().GetComposite().GetResource(), o); err != nil {
		panic(err)
	}
}

func ConvertDesiredResourceToObject(r *fnapi.RunFunctionResponse, name string, o runtime.Object) {
	if obj, exists := r.GetDesired().GetResources()[name]; !exists {
		panic("could not get resource from response: " + name)
	} else {
		if err := resource.AsObject(obj.GetResource(), o); err != nil {
			panic(err)
		}
	}
}
