/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"strings"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	tjname "github.com/crossplane/upjet/pkg/types/name"

	"github.com/johnmckenna/provider-dynatrace/config/accesstokens"
	"github.com/johnmckenna/provider-dynatrace/config/alerting"
	"github.com/johnmckenna/provider-dynatrace/config/anomalydetection"
	"github.com/johnmckenna/provider-dynatrace/config/appengine"
	"github.com/johnmckenna/provider-dynatrace/config/applicationsecurity"
	"github.com/johnmckenna/provider-dynatrace/config/automation"
	"github.com/johnmckenna/provider-dynatrace/config/browsermonitors"
	"github.com/johnmckenna/provider-dynatrace/config/businessevents"
	"github.com/johnmckenna/provider-dynatrace/config/cloudplatforms"
	"github.com/johnmckenna/provider-dynatrace/config/clustermanagement"
	"github.com/johnmckenna/provider-dynatrace/config/containers"
	"github.com/johnmckenna/provider-dynatrace/config/credentials"
	"github.com/johnmckenna/provider-dynatrace/config/dashboards"
	"github.com/johnmckenna/provider-dynatrace/config/deprecated"
	"github.com/johnmckenna/provider-dynatrace/config/developerobservability"
	"github.com/johnmckenna/provider-dynatrace/config/documents"
	"github.com/johnmckenna/provider-dynatrace/config/environmentsettings"
	"github.com/johnmckenna/provider-dynatrace/config/extensions"
	"github.com/johnmckenna/provider-dynatrace/config/failuredetection"
	"github.com/johnmckenna/provider-dynatrace/config/hostmonitoring"
	"github.com/johnmckenna/provider-dynatrace/config/httpmonitors"
	"github.com/johnmckenna/provider-dynatrace/config/iam"
	"github.com/johnmckenna/provider-dynatrace/config/incubator"
	"github.com/johnmckenna/provider-dynatrace/config/integrations"
	"github.com/johnmckenna/provider-dynatrace/config/logmonitoring"
	"github.com/johnmckenna/provider-dynatrace/config/mainframe"
	"github.com/johnmckenna/provider-dynatrace/config/managementzones"
	"github.com/johnmckenna/provider-dynatrace/config/metrics"
	"github.com/johnmckenna/provider-dynatrace/config/mobilecustomapplications"
	"github.com/johnmckenna/provider-dynatrace/config/monitoredentities"
	"github.com/johnmckenna/provider-dynatrace/config/monitoredtechnologies"
	"github.com/johnmckenna/provider-dynatrace/config/networkavailabilitymonitors"
	"github.com/johnmckenna/provider-dynatrace/config/notifications"
	"github.com/johnmckenna/provider-dynatrace/config/opentelemetryopentracing"
	"github.com/johnmckenna/provider-dynatrace/config/ownership"
	"github.com/johnmckenna/provider-dynatrace/config/platform"
	"github.com/johnmckenna/provider-dynatrace/config/processgroupmonitoring"
	"github.com/johnmckenna/provider-dynatrace/config/processmonitoring"
	"github.com/johnmckenna/provider-dynatrace/config/realusermonitoring"
	"github.com/johnmckenna/provider-dynatrace/config/servicedetection"
	"github.com/johnmckenna/provider-dynatrace/config/servicelevelobjective"
	"github.com/johnmckenna/provider-dynatrace/config/servicemonitoring"
	"github.com/johnmckenna/provider-dynatrace/config/servicesettings"
	"github.com/johnmckenna/provider-dynatrace/config/sessionreplay"
	"github.com/johnmckenna/provider-dynatrace/config/synthetic"
	"github.com/johnmckenna/provider-dynatrace/config/tags"
	"github.com/johnmckenna/provider-dynatrace/config/topologymodel"
	"github.com/johnmckenna/provider-dynatrace/config/updates"
	"github.com/johnmckenna/provider-dynatrace/config/usersettings"
	"github.com/johnmckenna/provider-dynatrace/config/virtualization"
	"github.com/johnmckenna/provider-dynatrace/config/webapplications"
)

const (
	resourcePrefix = "dynatrace"
	modulePath     = "github.com/johnmckenna/provider-dynatrace"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// Dynatrace naming convention does not match Crossplane defaults
func KindOverrides() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		r.Kind = tjname.NewFromSnake(strings.Join(strings.Split(r.Name, "_")[1:], "_")).Camel
	}
}

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("dynatrace.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			KindOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		accesstokens.Configure,
		alerting.Configure,
		anomalydetection.Configure,
		appengine.Configure,
		applicationsecurity.Configure,
		automation.Configure,
		browsermonitors.Configure,
		businessevents.Configure,
		cloudplatforms.Configure,
		clustermanagement.Configure,
		containers.Configure,
		credentials.Configure,
		dashboards.Configure,
		deprecated.Configure,
		developerobservability.Configure,
		documents.Configure,
		environmentsettings.Configure,
		extensions.Configure,
		failuredetection.Configure,
		hostmonitoring.Configure,
		httpmonitors.Configure,
		iam.Configure,
		incubator.Configure,
		integrations.Configure,
		logmonitoring.Configure,
		mainframe.Configure,
		managementzones.Configure,
		metrics.Configure,
		mobilecustomapplications.Configure,
		monitoredentities.Configure,
		monitoredtechnologies.Configure,
		networkavailabilitymonitors.Configure,
		notifications.Configure,
		opentelemetryopentracing.Configure,
		ownership.Configure,
		platform.Configure,
		processgroupmonitoring.Configure,
		processmonitoring.Configure,
		realusermonitoring.Configure,
		servicedetection.Configure,
		servicelevelobjective.Configure,
		servicemonitoring.Configure,
		servicesettings.Configure,
		sessionreplay.Configure,
		synthetic.Configure,
		tags.Configure,
		topologymodel.Configure,
		updates.Configure,
		usersettings.Configure,
		virtualization.Configure,
		webapplications.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
