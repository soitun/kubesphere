/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubesphere.io/api/tenant/v1alpha1"
)

// WorkspaceLister helps list Workspaces.
type WorkspaceLister interface {
	// List lists all Workspaces in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Workspace, err error)
	// Get retrieves the Workspace from the index for a given name.
	Get(name string) (*v1alpha1.Workspace, error)
	WorkspaceListerExpansion
}

// workspaceLister implements the WorkspaceLister interface.
type workspaceLister struct {
	indexer cache.Indexer
}

// NewWorkspaceLister returns a new WorkspaceLister.
func NewWorkspaceLister(indexer cache.Indexer) WorkspaceLister {
	return &workspaceLister{indexer: indexer}
}

// List lists all Workspaces in the indexer.
func (s *workspaceLister) List(selector labels.Selector) (ret []*v1alpha1.Workspace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Workspace))
	})
	return ret, err
}

// Get retrieves the Workspace from the index for a given name.
func (s *workspaceLister) Get(name string) (*v1alpha1.Workspace, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workspace"), name)
	}
	return obj.(*v1alpha1.Workspace), nil
}
