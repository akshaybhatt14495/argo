// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	workflowv1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	versioned "github.com/argoproj/argo/pkg/client/clientset/versioned"
	internalinterfaces "github.com/argoproj/argo/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/argoproj/argo/pkg/client/listers/workflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CronWorkflowInformer provides access to a shared informer and lister for
// CronWorkflows.
type CronWorkflowInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CronWorkflowLister
}

type cronWorkflowInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCronWorkflowInformer constructs a new informer for CronWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCronWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCronWorkflowInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCronWorkflowInformer constructs a new informer for CronWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCronWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().CronWorkflows(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().CronWorkflows(namespace).Watch(context.TODO(), options)
			},
		},
		&workflowv1alpha1.CronWorkflow{},
		resyncPeriod,
		indexers,
	)
}

func (f *cronWorkflowInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCronWorkflowInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cronWorkflowInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workflowv1alpha1.CronWorkflow{}, f.defaultInformer)
}

func (f *cronWorkflowInformer) Lister() v1alpha1.CronWorkflowLister {
	return v1alpha1.NewCronWorkflowLister(f.Informer().GetIndexer())
}
