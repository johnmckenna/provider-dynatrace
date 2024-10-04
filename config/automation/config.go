package automation

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_automation_workflow_jira",
		"dynatrace_automation_workflow_slack",
		"dynatrace_automation_business_calendar",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_business_calendar",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_business_calendar",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_business_calendar",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_scheduling_rule",
		"dynatrace_automation_workflow_k8s_connections",
		"dynatrace_automation_workflow_aws_connections",
		"dynatrace_automation_business_calendar",
		"dynatrace_automation_workflow",
		"dynatrace_site_reliability_guardian",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "automation"
		})
	}
}
