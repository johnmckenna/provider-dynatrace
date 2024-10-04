package environmentsettings

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_data_privacy",
		"dynatrace_ip_address_masking",
		"dynatrace_eula_settings",
		"dynatrace_oneagent_default_version",
		"dynatrace_app_monitoring",
		"dynatrace_oneagent_default_mode",
		"dynatrace_audit_log",
		"dynatrace_hub_permissions",
		"dynatrace_limit_outbound_connections",
		"dynatrace_hub_subscriptions",
		"dynatrace_oneagent_features",
		"dynatrace_disk_options",
		"dynatrace_oneagent_side_masking",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "environmentsettings"
		})
	}
}
