// SPDX-FileCopyrightText: Copyright DB InfraGO AG and contributors
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"encoding/json"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func mustUnstructuredFromYAML(rawYAML []byte) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	if err := yaml.Unmarshal(rawYAML, u); err != nil {
		panic(err.Error())
	}
	return u
}

func mustUnstructuredFromJSON(rawJSON []byte) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	if err := json.Unmarshal(rawJSON, u); err != nil {
		panic(err.Error())
	}
	return u
}
