// +build all resource_serviceendpoint_dockerregistry
// +build !exclude_serviceendpoints

package acceptancetests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
)

// validates that an apply followed by another apply (i.e., resource update) will be reflected in AzDO and the
// underlying terraform state.
func TestAccServiceEndpointDockerRegistry_CreateAndUpdate(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	serviceEndpointNameFirst := testutils.GenerateResourceName()
	serviceEndpointNameSecond := testutils.GenerateResourceName()

	resourceType := "azuredevops_serviceendpoint_dockerregistry"
	tfSvcEpNode := resourceType + ".serviceendpoint"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testutils.PreCheck(t, &[]string{
				"AZDO_DOCKERREGISTRY_SERVICE_CONNECTION_USERNAME",
				"AZDO_DOCKERREGISTRY_SERVICE_CONNECTION_EMAIL",
				"AZDO_DOCKERREGISTRY_SERVICE_CONNECTION_PASSWORD",
			})
		},
		Providers:    testutils.GetProviders(),
		CheckDestroy: testutils.CheckServiceEndpointDestroyed(resourceType),
		Steps: []resource.TestStep{
			{
				Config: testutils.HclServiceEndpointDockerRegistryResource(projectName, serviceEndpointNameFirst),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "docker_username"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "docker_email"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "docker_password", ""),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "docker_password_hash"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "service_endpoint_name", serviceEndpointNameFirst),
					testutils.CheckServiceEndpointExistsWithName(tfSvcEpNode, serviceEndpointNameFirst),
				),
			}, {
				Config: testutils.HclServiceEndpointDockerRegistryResource(projectName, serviceEndpointNameSecond),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "docker_username"),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "docker_email"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "docker_password", ""),
					resource.TestCheckResourceAttrSet(tfSvcEpNode, "docker_password_hash"),
					resource.TestCheckResourceAttr(tfSvcEpNode, "service_endpoint_name", serviceEndpointNameSecond),
					testutils.CheckServiceEndpointExistsWithName(tfSvcEpNode, serviceEndpointNameSecond)),
			},
		},
	})
}
