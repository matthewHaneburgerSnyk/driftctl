package azurerm

import (
	"github.com/snyk/driftctl/enumeration/resource"
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

const AzurePublicIPResourceType = "azurerm_public_ip"

func initAzurePublicIPMetadata(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(AzurePublicIPResourceType, func(res *resource.Resource) map[string]string {
		val := res.Attrs
		attrs := make(map[string]string)
		if name := val.GetString("name"); name != nil && *name != "" {
			attrs["Name"] = *name
		}
		return attrs
	})
}
