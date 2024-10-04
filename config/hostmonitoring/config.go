package hostmonitoring

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_aix_extension",
		"dynatrace_host_monitoring",
		"dynatrace_ebpf_service_discovery",
		"dynatrace_disk_analytics",
		"dynatrace_host_monitoring_mode",
		"dynatrace_host_monitoring_advanced",
		"dynatrace_network_traffic",
		"dynatrace_os_services",
		"dynatrace_host_process_group_monitoring",
		"dynatrace_nettracer",
		"dynatrace_crashdump_analytics",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "hostmonitoring"
		})
	}
}
