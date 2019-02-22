// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type SecurityRuleClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*SecurityRule, error)
	Write(resource *SecurityRule, opts clients.WriteOpts) (*SecurityRule, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (SecurityRuleList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan SecurityRuleList, <-chan error, error)
}

type securityRuleClient struct {
	rc clients.ResourceClient
}

func NewSecurityRuleClient(rcFactory factory.ResourceClientFactory) (SecurityRuleClient, error) {
	return NewSecurityRuleClientWithToken(rcFactory, "")
}

func NewSecurityRuleClientWithToken(rcFactory factory.ResourceClientFactory, token string) (SecurityRuleClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &SecurityRule{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base SecurityRule resource client")
	}
	return NewSecurityRuleClientWithBase(rc), nil
}

func NewSecurityRuleClientWithBase(rc clients.ResourceClient) SecurityRuleClient {
	return &securityRuleClient{
		rc: rc,
	}
}

func (client *securityRuleClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *securityRuleClient) Register() error {
	return client.rc.Register()
}

func (client *securityRuleClient) Read(namespace, name string, opts clients.ReadOpts) (*SecurityRule, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*SecurityRule), nil
}

func (client *securityRuleClient) Write(securityRule *SecurityRule, opts clients.WriteOpts) (*SecurityRule, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(securityRule, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*SecurityRule), nil
}

func (client *securityRuleClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *securityRuleClient) List(namespace string, opts clients.ListOpts) (SecurityRuleList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToSecurityRule(resourceList), nil
}

func (client *securityRuleClient) Watch(namespace string, opts clients.WatchOpts) (<-chan SecurityRuleList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	securityrulesChan := make(chan SecurityRuleList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				securityrulesChan <- convertToSecurityRule(resourceList)
			case <-opts.Ctx.Done():
				close(securityrulesChan)
				return
			}
		}
	}()
	return securityrulesChan, errs, nil
}

func convertToSecurityRule(resources resources.ResourceList) SecurityRuleList {
	var securityRuleList SecurityRuleList
	for _, resource := range resources {
		securityRuleList = append(securityRuleList, resource.(*SecurityRule))
	}
	return securityRuleList
}
