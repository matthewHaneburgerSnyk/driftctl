package google

import (
	"github.com/snyk/driftctl/enumeration/resource"
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

const GoogleComputeGlobalAddressResourceType = "google_compute_global_address"

func initGoogleComputeGlobalAddressMetadata(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(GoogleComputeGlobalAddressResourceType, func(res *resource.Resource) map[string]string {
		return map[string]string{
			"Name":    *res.Attributes().GetString("name"),
			"Address": *res.Attributes().GetString("address"),
		}
	})
}
