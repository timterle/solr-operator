/*
Copyright 2019 Bloomberg Finance LP.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	solrv1beta1 "github.com/bloomberg/solr-operator/pkg/apis/solr/v1beta1"
	"github.com/bloomberg/solr-operator/pkg/client/clientset/versioned"
	"github.com/bloomberg/solr-operator/pkg/client/informers/externalversions/internalinterfaces"
	"github.com/bloomberg/solr-operator/pkg/client/listers/solr/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

// SolrBackupInformer provides access to a shared informer and lister for
// SolrBackups.
type SolrBackupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.SolrBackupLister
}

type solrBackupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSolrBackupInformer constructs a new informer for SolrBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSolrBackupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSolrBackupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSolrBackupInformer constructs a new informer for SolrBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSolrBackupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SolrV1beta1().SolrBackups(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SolrV1beta1().SolrBackups(namespace).Watch(options)
			},
		},
		&solrv1beta1.SolrBackup{},
		resyncPeriod,
		indexers,
	)
}

func (f *solrBackupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSolrBackupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *solrBackupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&solrv1beta1.SolrBackup{}, f.defaultInformer)
}

func (f *solrBackupInformer) Lister() v1beta1.SolrBackupLister {
	return v1beta1.NewSolrBackupLister(f.Informer().GetIndexer())
}
