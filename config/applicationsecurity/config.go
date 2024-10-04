package applicationsecurity

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_vulnerability_third_party",
		"dynatrace_vulnerability_code",
		"dynatrace_appsec_notification",
		"dynatrace_attack_alerting",
		"dynatrace_attack_allowlist",
		"dynatrace_vulnerability_settings",
		"dynatrace_attack_settings",
		"dynatrace_attack_rules",
		"dynatrace_management_zone_v2",
		"dynatrace_vulnerability_alerting",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "applicationsecurity"
		})
	}
}
