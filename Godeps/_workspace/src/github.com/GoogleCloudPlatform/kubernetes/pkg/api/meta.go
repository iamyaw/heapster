/*
Copyright 2014 Google Inc. All rights reserved.

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

package api

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/conversion"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
)

// FillObjectMetaSystemFields populates fields that are managed by the system on ObjectMeta.
func FillObjectMetaSystemFields(ctx Context, meta *ObjectMeta) {
	meta.CreationTimestamp = util.Now()
	meta.UID = util.NewUUID()
	meta.SelfLink = ""
}

// HasObjectMetaSystemFieldValues returns true if fields that are managed by the system on ObjectMeta have values.
func HasObjectMetaSystemFieldValues(meta *ObjectMeta) bool {
	return !meta.CreationTimestamp.Time.IsZero() ||
		len(meta.UID) != 0
}

// ObjectMetaFor returns a pointer to a provided object's ObjectMeta.
// TODO: allow runtime.Unknown to extract this object
func ObjectMetaFor(obj runtime.Object) (*ObjectMeta, error) {
	v, err := conversion.EnforcePtr(obj)
	if err != nil {
		return nil, err
	}
	var meta *ObjectMeta
	err = runtime.FieldPtr(v, "ObjectMeta", &meta)
	return meta, err
}

// ListMetaFor returns a pointer to a provided object's ListMeta,
// or an error if the object does not have that pointer.
// TODO: allow runtime.Unknown to extract this object
func ListMetaFor(obj runtime.Object) (*ListMeta, error) {
	v, err := conversion.EnforcePtr(obj)
	if err != nil {
		return nil, err
	}
	var meta *ListMeta
	err = runtime.FieldPtr(v, "ListMeta", &meta)
	return meta, err
}
