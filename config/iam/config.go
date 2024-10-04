package iam

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_mgmz_permission",
		"dynatrace_user_group",
		"dynatrace_user",
		"dynatrace_policy",
		"dynatrace_policy",
		"dynatrace_policy_bindings",
		"dynatrace_policy_bindings",
		"dynatrace_iam_permission",
		"dynatrace_user_group",
		"dynatrace_policy",
		"dynatrace_policy_bindings",
		"dynatrace_policy",
		"dynatrace_policy_bindings",
		"dynatrace_user_group",
		"dynatrace_policy",
		"dynatrace_policy",
		"dynatrace_policy_bindings",
		"dynatrace_policy",
		"dynatrace_iam_policy",
		"dynatrace_user_group",
		"dynatrace_user",
		"dynatrace_policy",
		"dynatrace_policy",
		"dynatrace_policy_bindings",
		"dynatrace_policy_bindings",
		"dynatrace_iam_group",
		"dynatrace_iam_group",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy_bindings",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy_bindings",
		"dynatrace_iam_group",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy_bindings",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy",
		"dynatrace_iam_group",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy_bindings",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy_bindings_v2",
		"dynatrace_iam_group",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy",
		"dynatrace_iam_policy_bindings",
		"dynatrace_iam_user",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "iam"
		})
	}
}
