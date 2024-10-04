package notifications

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_xmatters_notification",
		"dynatrace_alerting",
		"dynatrace_slack_notification",
		"dynatrace_alerting",
		"dynatrace_pager_duty_notification",
		"dynatrace_alerting",
		"dynatrace_victor_ops_notification",
		"dynatrace_alerting",
		"dynatrace_email_notification",
		"dynatrace_alerting",
		"dynatrace_ansible_tower_notification",
		"dynatrace_alerting",
		"dynatrace_trello_notification",
		"dynatrace_alerting",
		"dynatrace_service_now_notification",
		"dynatrace_alerting",
		"dynatrace_mobile_notifications",
		"dynatrace_webhook_notification",
		"dynatrace_alerting",
		"dynatrace_ops_genie_notification",
		"dynatrace_alerting",
		"dynatrace_jira_notification",
		"dynatrace_alerting",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "notifications"
		})
	}
}
