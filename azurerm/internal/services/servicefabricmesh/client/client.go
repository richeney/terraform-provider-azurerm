package client

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/servicefabricmesh/mgmt/2018-09-01-preview/servicefabricmesh"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	ApplicationClient *servicefabricmesh.ApplicationClient
	ServiceClient     *servicefabricmesh.ServiceClient
}

func NewClient(o *common.ClientOptions) *Client {
	applicationsClient := servicefabricmesh.NewApplicationClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&applicationsClient.Client, o.ResourceManagerAuthorizer)

	servicesClient := servicefabricmesh.NewServiceClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&servicesClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		ApplicationClient: &applicationsClient,
		ServiceClient:     &servicesClient,
	}
}
