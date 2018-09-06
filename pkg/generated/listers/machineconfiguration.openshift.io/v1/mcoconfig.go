// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MCOConfigLister helps list MCOConfigs.
type MCOConfigLister interface {
	// List lists all MCOConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1.MCOConfig, err error)
	// MCOConfigs returns an object that can list and get MCOConfigs.
	MCOConfigs(namespace string) MCOConfigNamespaceLister
	MCOConfigListerExpansion
}

// mCOConfigLister implements the MCOConfigLister interface.
type mCOConfigLister struct {
	indexer cache.Indexer
}

// NewMCOConfigLister returns a new MCOConfigLister.
func NewMCOConfigLister(indexer cache.Indexer) MCOConfigLister {
	return &mCOConfigLister{indexer: indexer}
}

// List lists all MCOConfigs in the indexer.
func (s *mCOConfigLister) List(selector labels.Selector) (ret []*v1.MCOConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MCOConfig))
	})
	return ret, err
}

// MCOConfigs returns an object that can list and get MCOConfigs.
func (s *mCOConfigLister) MCOConfigs(namespace string) MCOConfigNamespaceLister {
	return mCOConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MCOConfigNamespaceLister helps list and get MCOConfigs.
type MCOConfigNamespaceLister interface {
	// List lists all MCOConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MCOConfig, err error)
	// Get retrieves the MCOConfig from the indexer for a given namespace and name.
	Get(name string) (*v1.MCOConfig, error)
	MCOConfigNamespaceListerExpansion
}

// mCOConfigNamespaceLister implements the MCOConfigNamespaceLister
// interface.
type mCOConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MCOConfigs in the indexer for a given namespace.
func (s mCOConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.MCOConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MCOConfig))
	})
	return ret, err
}

// Get retrieves the MCOConfig from the indexer for a given namespace and name.
func (s mCOConfigNamespaceLister) Get(name string) (*v1.MCOConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("mcoconfig"), name)
	}
	return obj.(*v1.MCOConfig), nil
}