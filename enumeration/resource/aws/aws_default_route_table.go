package aws

import "github.com/snyk/driftctl/enumeration/resource"

const AwsDefaultRouteTableResourceType = "aws_default_route_table"

func initAwsDefaultRouteTableMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetFlags(AwsDefaultRouteTableResourceType, resource.FlagDeepMode)
	resourceSchemaRepository.SetNormalizeFunc(AwsDefaultRouteTableResourceType, func(res *resource.Resource) {
		val := res.Attrs
		val.SafeDelete([]string{"timeouts"})
	})
}
