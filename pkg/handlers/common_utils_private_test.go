/*
Copyright 2023 Riskified Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handlers

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// A hack for implementing a k8s client inside a test function.
// see: https://stackoverflow.com/questions/31362044/anonymous-interface-implementation-in-golang
type MockClient struct {
	client.Client
	getMethod  func(context.Context, types.NamespacedName, client.Object, ...client.GetOption) error
	listMethod func(context.Context, client.ObjectList, ...client.ListOption) error
}

func (m MockClient) Get(c context.Context, ns types.NamespacedName, o client.Object, _ ...client.GetOption) error {
	return m.getMethod(c, ns, o)
}

func (m MockClient) List(c context.Context, l client.ObjectList, o ...client.ListOption) error {
	return m.listMethod(c, l, o...)
}
