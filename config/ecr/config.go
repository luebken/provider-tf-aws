package ecr

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/crossplane-contrib/provider-tf-aws/config/ecr"
)

// PublicRepositoryExternalName is used to set repository name from external name
func PublicRepositoryExternalName(base map[string]interface{}, externalName string) {
	base["repository_name"] = externalName
}

func init() {

	config.Store.SetForResource("aws_ecr_repository", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
			},
		},

		References: map[string]config.Reference{
			"encryption_configuration[*].kms_key": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
				Extractor: common.PathARNExtractor,
			},
		},

		// Note(turkenh): only delete has a timeout of 20 mins, we only need to
		// use async with destroy:
		// https://github.com/crossplane-contrib/terrajet/issues/102
		UseAsync: true,
	})

	config.Store.SetForResource("aws_ecrpublic_repository", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: SelfPackagePath + ".PublicRepositoryExternalName",
			OmittedFields: []string{
				"repository_name",
			},
		},

		// Note(turkenh): only delete has a timeout of 20 mins, we only need to
		// use async with destroy:
		// https://github.com/crossplane-contrib/terrajet/issues/102
		UseAsync: true,
	})
}
