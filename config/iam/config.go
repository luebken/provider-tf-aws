/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package iam

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

func init() {
	config.Store.SetForResource("aws_iam_access_key", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"user": config.Reference{
				Type: "User",
			},
		},
	})

	config.Store.SetForResource("aws_iam_group", config.Resource{
		// Cannot use kind "Group" since it conflict with generated constant
		// name in group version info
		Kind: "IAMGroup",
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
			},
		},
	})

	config.Store.SetForResource("aws_iam_group_policy", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
		References: config.References{
			"group": config.Reference{
				Type: "IAMGroup",
			},
		},
	})

	config.Store.SetForResource("aws_iam_group_policy_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"group": config.Reference{
				Type: "IAMGroup",
			},
			"policy_arn": config.Reference{
				Type:      "Policy",
				Extractor: common.PathARNExtractor,
			},
		},
	})

	config.Store.SetForResource("aws_iam_instance_profile", config.Resource{
		ExternalName: config.ExternalName{
			ConfigureFunctionPath: common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
		References: config.References{
			"role": config.Reference{
				Type: "Role",
			},
		},
	})

	config.Store.SetForResource("aws_iam_policy", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
	})

	config.Store.SetForResource("aws_iam_policy_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"users": config.Reference{
				Type: "User",
			},
			"roles": config.Reference{
				Type: "Role",
			},
			"groups": config.Reference{
				Type: "IAMGroup",
			},
			"policy_arn": config.Reference{
				Type:      "Policy",
				Extractor: common.PathARNExtractor,
			},
		},
	})

	config.Store.SetForResource("aws_iam_role", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: false,
			ConfigureFunctionPath:  common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
	})

	config.Store.SetForResource("aws_iam_role_policy", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: false,
			ConfigureFunctionPath:  common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
				"name_prefix",
			},
		},
	})
	config.Store.SetForResource("aws_iam_role_policy_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"role": config.Reference{
				Type: "Role",
			},
			"policy_arn": config.Reference{
				Type:      "Policy",
				Extractor: common.PathARNExtractor,
			},
		},
	})

	config.Store.SetForResource("aws_iam_user", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: false,
			ConfigureFunctionPath:  common.PathExternalNameAsName,
			OmittedFields: []string{
				"name",
			},
		},
	})

	config.Store.SetForResource("aws_iam_user_group_membership", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"user": config.Reference{
				Type: "User",
			},
			"groups": config.Reference{
				Type: "IAMGroup",
			},
		},
	})

	config.Store.SetForResource("aws_iam_user_policy", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"user": config.Reference{
				Type: "User",
			},
		},
	})
	config.Store.SetForResource("aws_iam_user_policy_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: config.References{
			"user": config.Reference{
				Type: "User",
			},
			"policy_arn": config.Reference{
				Type:      "Policy",
				Extractor: common.PathARNExtractor,
			},
		},
	})
}
