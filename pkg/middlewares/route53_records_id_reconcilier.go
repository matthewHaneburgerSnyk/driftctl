package middlewares

import (
	"strings"

	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/aws"
	"github.com/sirupsen/logrus"
)

// Since AWS returns the FQDN as the name of the remote record, we must change the Id of the
// state record to be equivalent (ZoneId_FQDN_Type_SetIdentifier)
// For a TXT record toto for zone example.com with Id 1234
// From AWS provider, we retrieve: 1234_toto.example.com_TXT
// From Terraform state, we retrieve: 1234_toto_TXT
type Route53RecordIDReconcilier struct{}

func NewRoute53RecordIDReconcilier() Route53RecordIDReconcilier {
	return Route53RecordIDReconcilier{}
}

func (m Route53RecordIDReconcilier) Execute(_, resourcesFromState *[]resource.Resource) error {

	for _, stateResource := range *resourcesFromState {

		if stateResource.TerraformType() != aws.AwsRoute53RecordResourceType {
			continue
		}

		record, _ := stateResource.(*resource.AbstractResource)

		vars := []string{
			(*record.Attrs)["zone_id"].(string),
			(*record.Attrs)["fqdn"].(string),
			(*record.Attrs)["type"].(string),
		}
		newId := strings.Join(vars, "_")
		if newId != record.Id {
			record.Id = newId
			_ = record.Attrs.SafeSet([]string{"id"}, newId)
			logrus.WithFields(logrus.Fields{
				"old_id": record.TerraformId(),
				"new_id": newId,
			}).Debug("Normalized route53 record ID")
		}
	}

	return nil
}
