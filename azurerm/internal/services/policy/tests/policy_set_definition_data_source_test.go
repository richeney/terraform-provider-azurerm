package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccDataSourceAzureRMPolicySetDefinition_builtIn(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_policy_set_definition", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMPolicySetDefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureRMPolicySetDefinition_builtIn("Audit Windows VMs with a pending reboot"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "name", "c96b2a9c-6fab-4ac2-ae21-502143491cd4"),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", "Audit Windows VMs with a pending reboot"),
					resource.TestCheckResourceAttr(data.ResourceName, "policy_type", "BuiltIn"),
					resource.TestCheckResourceAttr(data.ResourceName, "parameters", ""),
					resource.TestCheckResourceAttrSet(data.ResourceName, "policy_definitions"),
					resource.TestCheckResourceAttr(data.ResourceName, "policy_definition_reference.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMPolicySetDefinition_customByName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_policy_set_definition", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMPolicySetDefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureRMPolicySetDefinition_customByName(data),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctestPolSet-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", fmt.Sprintf("acctestPolSet-display-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "policy_type", "Custom"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "parameters"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "policy_definitions"),
					resource.TestCheckResourceAttr(data.ResourceName, "policy_definition_reference.#", "1"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMPolicySetDefinition_customByDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_policy_set_definition", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMPolicySetDefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureRMPolicySetDefinition_customByDisplayName(data),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctestPolSet-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", fmt.Sprintf("acctestPolSet-display-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "policy_type", "Custom"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "parameters"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "policy_definitions"),
					resource.TestCheckResourceAttr(data.ResourceName, "policy_definition_reference.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceAzureRMPolicySetDefinition_builtIn(name string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

data "azurerm_policy_set_definition" "test" {
  display_name = "%s"
}
`, name)
}

func testAccDataSourceAzureRMPolicySetDefinition_customByName(data acceptance.TestData) string {
	template := testAzureRMPolicySetDefinition_custom(data)
	return fmt.Sprintf(`
%s

data "azurerm_policy_set_definition" "test" {
  name = azurerm_policy_set_definition.test.name
}
`, template)
}

func testAccDataSourceAzureRMPolicySetDefinition_customByDisplayName(data acceptance.TestData) string {
	template := testAzureRMPolicySetDefinition_custom(data)
	return fmt.Sprintf(`
%s

data "azurerm_policy_set_definition" "test" {
  display_name = azurerm_policy_set_definition.test.display_name
}
`, template)
}
