/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	auth "tkestack.io/tke/api/auth"
)

// ClientLister helps list Clients.
type ClientLister interface {
	// List lists all Clients in the indexer.
	List(selector labels.Selector) (ret []*auth.Client, err error)
	// Get retrieves the Client from the index for a given name.
	Get(name string) (*auth.Client, error)
	ClientListerExpansion
}

// clientLister implements the ClientLister interface.
type clientLister struct {
	indexer cache.Indexer
}

// NewClientLister returns a new ClientLister.
func NewClientLister(indexer cache.Indexer) ClientLister {
	return &clientLister{indexer: indexer}
}

// List lists all Clients in the indexer.
func (s *clientLister) List(selector labels.Selector) (ret []*auth.Client, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*auth.Client))
	})
	return ret, err
}

// Get retrieves the Client from the index for a given name.
func (s *clientLister) Get(name string) (*auth.Client, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(auth.Resource("client"), name)
	}
	return obj.(*auth.Client), nil
}