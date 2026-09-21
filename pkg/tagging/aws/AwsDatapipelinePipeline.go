package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsDatapipelinePipeline struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsDatapipelinePipeline{}
)

func GetAwsDatapipelinePipeline() core.ResourceInterface {
	this := &AwsDatapipelinePipeline{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Data Pipeline",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::DataPipeline::Pipeline",
			Short: "dp",
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
			ResourceName: "aws_datapipeline_pipeline",
		},
	}

	return this
}
