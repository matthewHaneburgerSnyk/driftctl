package azurerm

import "github.com/snyk/driftctl/enumeration/resource"

const AzureLoadBalancerRuleResourceType = "azurerm_lb_rule"

func initAzureLoadBalancerRuleMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(AzureLoadBalancerRuleResourceType, func(res *resource.Resource) map[string]string {
		attrs := make(map[string]string)
		if name := res.Attributes().GetString("name"); name != nil && *name != "" {
			attrs["Name"] = *name
		}
		return attrs
	})
	resourceSchemaRepository.SetFlags(AzureLoadBalancerRuleResourceType, resource.FlagDeepMode)
}
