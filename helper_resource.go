// SPDX-FileCopyrightText: Copyright DB InfraGO AG and contributors
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"github.com/crossplane/function-sdk-go/resource"
	"google.golang.org/protobuf/types/known/structpb"
	"k8s.io/apimachinery/pkg/runtime"
)

func mustObjectAsStruct(o runtime.Object) *structpb.Struct {
	s, err := resource.AsStruct(o)
	if err != nil {
		panic(err.Error())
	}
	return s
}

func mustStructValue(in any) *structpb.Value {
	val, err := structpb.NewValue(in)
	if err != nil {
		panic(err.Error())
	}
	return val
}
