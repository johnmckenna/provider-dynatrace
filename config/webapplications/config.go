package webapplications

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_calculated_web_metric",
		"dynatrace_web_app_custom_config_properties",
		"dynatrace_web_app_resource_types",
		"dynatrace_web_app_beacon_origins",
		"dynatrace_web_app_beacon_endpoint",
		"dynatrace_web_app_custom_injection",
		"dynatrace_application_detection_rule_v2",
		"dynatrace_web_app_key_performance_xhr",
		"dynatrace_web_app_resource_cleanup",
		"dynatrace_web_app_injection_cookie",
		"dynatrace_web_app_custom_errors",
		"dynatrace_web_app_enablement",
		"dynatrace_web_application",
		"dynatrace_web_app_request_errors",
		"dynatrace_key_user_action",
		"dynatrace_web_application",
		"dynatrace_web_app_key_performance_custom",
		"dynatrace_web_app_key_performance_load",
		"dynatrace_web_app_javascript_version",
		"dynatrace_web_app_javascript_updates",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "webapplications"
		})
	}
}
