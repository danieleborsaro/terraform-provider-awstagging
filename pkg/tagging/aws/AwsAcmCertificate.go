package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsAcmCertificate struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsAcmCertificate{}
)

func GetAwsAcmCertificate() core.ResourceInterface {
	this := &AwsAcmCertificate{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS:ACM",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::CertificateManager::Certificate",
			Short: "acmc",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           255,
			IsLowerCase:         false,
			IsEnforceVersioning: false,
			Components:          []string{},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_acm_certificate",
		},
	}

	return this
}
