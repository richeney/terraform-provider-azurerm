package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type DataFactoryLinkedServiceId struct {
	ResourceGroup string
	Name          string
	DataFactory   string
}

func DataFactoryLinkedServiceID(input string) (*DataFactoryLinkedServiceId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Data Factory Linked Service ID %q: %+v", input, err)
	}

	dataFactoryIntegrationRuntime := DataFactoryLinkedServiceId{
		ResourceGroup: id.ResourceGroup,
	}

	if dataFactoryIntegrationRuntime.DataFactory, err = id.PopSegment("factories"); err != nil {
		return nil, err
	}

	if dataFactoryIntegrationRuntime.Name, err = id.PopSegment("linkedservices"); err != nil {
		return nil, err
	}

	return &dataFactoryIntegrationRuntime, nil
}
