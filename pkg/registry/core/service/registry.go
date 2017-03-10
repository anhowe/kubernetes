/*
Copyright 2014 The Kubernetes Authors.

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

package service

import (
	"fmt"

	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/api"
)

// Registry is an interface for things that know how to store services.
type Registry interface {
	ListServices(ctx genericapirequest.Context, options *metainternalversion.ListOptions) (*api.ServiceList, error)
	CreateService(ctx genericapirequest.Context, svc *api.Service) (*api.Service, error)
	GetService(ctx genericapirequest.Context, name string, options *metav1.GetOptions) (*api.Service, error)
	DeleteService(ctx genericapirequest.Context, name string) error
	UpdateService(ctx genericapirequest.Context, svc *api.Service) (*api.Service, error)
	WatchServices(ctx genericapirequest.Context, options *metainternalversion.ListOptions) (watch.Interface, error)
	ExportService(ctx genericapirequest.Context, name string, options metav1.ExportOptions) (*api.Service, error)
}

// storage puts strong typing around storage calls
type storage struct {
	rest.StandardStorage
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched
// types will panic.
func NewRegistry(s rest.StandardStorage) Registry {
	return &storage{s}
}

func (s *storage) ListServices(ctx genericapirequest.Context, options *metainternalversion.ListOptions) (*api.ServiceList, error) {
	glog.V(2).Infof("[ListServices,")
	defer glog.V(2).Infof("ListServices]")
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*api.ServiceList), nil
}

func (s *storage) CreateService(ctx genericapirequest.Context, svc *api.Service) (*api.Service, error) {
	glog.V(2).Infof("[CreateService,")
	defer glog.V(2).Infof("CreateService]")
	obj, err := s.Create(ctx, svc)
	if err != nil {
		return nil, err
	}
	return obj.(*api.Service), nil
}

func (s *storage) GetService(ctx genericapirequest.Context, name string, options *metav1.GetOptions) (*api.Service, error) {
	glog.V(2).Infof("[GetService,")
	defer glog.V(2).Infof("GetService]")
	obj, err := s.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}
	return obj.(*api.Service), nil
}

func (s *storage) DeleteService(ctx genericapirequest.Context, name string) error {
	glog.V(2).Infof("[DeleteService,")
	defer glog.V(2).Infof("DeleteService]")
	_, _, err := s.Delete(ctx, name, nil)
	return err
}

func (s *storage) UpdateService(ctx genericapirequest.Context, svc *api.Service) (*api.Service, error) {
	glog.V(2).Infof("[UpdateService,")
	defer glog.V(2).Infof("UpdateService]")
	obj, _, err := s.Update(ctx, svc.Name, rest.DefaultUpdatedObjectInfo(svc, api.Scheme))
	if err != nil {
		return nil, err
	}
	return obj.(*api.Service), nil
}

func (s *storage) WatchServices(ctx genericapirequest.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	glog.V(2).Infof("[DeleteService,")
	defer glog.V(2).Infof("DeleteService]")
	return s.Watch(ctx, options)
}

// If StandardStorage implements rest.Exporter, returns exported service.
// Otherwise export is not supported.
func (s *storage) ExportService(ctx genericapirequest.Context, name string, options metav1.ExportOptions) (*api.Service, error) {
	glog.V(2).Infof("[ExportService,")
	defer glog.V(2).Infof("ExportService]")
	exporter, isExporter := s.StandardStorage.(rest.Exporter)
	if !isExporter {
		return nil, fmt.Errorf("export is not supported")
	}
	obj, err := exporter.Export(ctx, name, options)
	if err != nil {
		return nil, err
	}
	return obj.(*api.Service), nil
}
