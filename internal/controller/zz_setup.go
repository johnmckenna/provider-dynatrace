// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	activegatetoken "github.com/johnmckenna/provider-dynatrace/internal/controller/accesstokens/activegatetoken"
	agtoken "github.com/johnmckenna/provider-dynatrace/internal/controller/accesstokens/agtoken"
	apitoken "github.com/johnmckenna/provider-dynatrace/internal/controller/accesstokens/apitoken"
	tokensettings "github.com/johnmckenna/provider-dynatrace/internal/controller/accesstokens/tokensettings"
	connectivityalerts "github.com/johnmckenna/provider-dynatrace/internal/controller/alerting/connectivityalerts"
	maintenance "github.com/johnmckenna/provider-dynatrace/internal/controller/alerting/maintenance"
	awsanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/awsanomalies"
	customappanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/customappanomalies"
	customappcrashrate "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/customappcrashrate"
	databaseanomaliesv2 "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/databaseanomaliesv2"
	davisanomalydetectors "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/davisanomalydetectors"
	diskanomaliesv2 "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/diskanomaliesv2"
	diskanomalyrules "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/diskanomalyrules"
	diskedgeanomalydetectors "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/diskedgeanomalydetectors"
	diskspecificanomaliesv2 "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/diskspecificanomaliesv2"
	frequentissues "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/frequentissues"
	hostanomaliesv2 "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/hostanomaliesv2"
	k8sclusteranomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/k8sclusteranomalies"
	k8snamespaceanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/k8snamespaceanomalies"
	k8snodeanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/k8snodeanomalies"
	k8spvcanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/k8spvcanomalies"
	k8sworkloadanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/k8sworkloadanomalies"
	metricevents "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/metricevents"
	mobileappanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/mobileappanomalies"
	mobileappcrashrate "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/mobileappcrashrate"
	serviceanomaliesv2 "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/serviceanomaliesv2"
	vmwareanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/vmwareanomalies"
	webappanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/anomalydetection/webappanomalies"
	dbappfeatureflags "github.com/johnmckenna/provider-dynatrace/internal/controller/appengine/dbappfeatureflags"
	discoverydefaultrules "github.com/johnmckenna/provider-dynatrace/internal/controller/appengine/discoverydefaultrules"
	discoveryfeatureflags "github.com/johnmckenna/provider-dynatrace/internal/controller/appengine/discoveryfeatureflags"
	infraopsappfeatureflags "github.com/johnmckenna/provider-dynatrace/internal/controller/appengine/infraopsappfeatureflags"
	infraopsappsettings "github.com/johnmckenna/provider-dynatrace/internal/controller/appengine/infraopsappsettings"
	appsecnotification "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/appsecnotification"
	attackalerting "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/attackalerting"
	attackallowlist "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/attackallowlist"
	attackrules "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/attackrules"
	attacksettings "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/attacksettings"
	vulnerabilityalerting "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/vulnerabilityalerting"
	vulnerabilitycode "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/vulnerabilitycode"
	vulnerabilitysettings "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/vulnerabilitysettings"
	vulnerabilitythirdparty "github.com/johnmckenna/provider-dynatrace/internal/controller/applicationsecurity/vulnerabilitythirdparty"
	automationbusinesscalendar "github.com/johnmckenna/provider-dynatrace/internal/controller/automation/automationbusinesscalendar"
	automationschedulingrule "github.com/johnmckenna/provider-dynatrace/internal/controller/automation/automationschedulingrule"
	automationworkflow "github.com/johnmckenna/provider-dynatrace/internal/controller/automation/automationworkflow"
	automationworkflowawsconnections "github.com/johnmckenna/provider-dynatrace/internal/controller/automation/automationworkflowawsconnections"
	automationworkflowjira "github.com/johnmckenna/provider-dynatrace/internal/controller/automation/automationworkflowjira"
	automationworkflowk8sconnections "github.com/johnmckenna/provider-dynatrace/internal/controller/automation/automationworkflowk8sconnections"
	automationworkflowslack "github.com/johnmckenna/provider-dynatrace/internal/controller/automation/automationworkflowslack"
	sitereliabilityguardian "github.com/johnmckenna/provider-dynatrace/internal/controller/automation/sitereliabilityguardian"
	browsermonitor "github.com/johnmckenna/provider-dynatrace/internal/controller/browsermonitors/browsermonitor"
	browsermonitoroutage "github.com/johnmckenna/provider-dynatrace/internal/controller/browsermonitors/browsermonitoroutage"
	browsermonitorperformance "github.com/johnmckenna/provider-dynatrace/internal/controller/browsermonitors/browsermonitorperformance"
	businesseventsbuckets "github.com/johnmckenna/provider-dynatrace/internal/controller/businessevents/businesseventsbuckets"
	businesseventsmetrics "github.com/johnmckenna/provider-dynatrace/internal/controller/businessevents/businesseventsmetrics"
	businesseventsoneagent "github.com/johnmckenna/provider-dynatrace/internal/controller/businessevents/businesseventsoneagent"
	businesseventsoneagentoutgoing "github.com/johnmckenna/provider-dynatrace/internal/controller/businessevents/businesseventsoneagentoutgoing"
	businesseventsprocessing "github.com/johnmckenna/provider-dynatrace/internal/controller/businessevents/businesseventsprocessing"
	businesseventssecuritycontext "github.com/johnmckenna/provider-dynatrace/internal/controller/businessevents/businesseventssecuritycontext"
	cloudfoundry "github.com/johnmckenna/provider-dynatrace/internal/controller/cloudplatforms/cloudfoundry"
	k8smonitoring "github.com/johnmckenna/provider-dynatrace/internal/controller/cloudplatforms/k8smonitoring"
	kubernetes "github.com/johnmckenna/provider-dynatrace/internal/controller/cloudplatforms/kubernetes"
	kubernetesapp "github.com/johnmckenna/provider-dynatrace/internal/controller/cloudplatforms/kubernetesapp"
	kubernetesenrichment "github.com/johnmckenna/provider-dynatrace/internal/controller/cloudplatforms/kubernetesenrichment"
	managedbackup "github.com/johnmckenna/provider-dynatrace/internal/controller/clustermanagement/managedbackup"
	managedinternetproxy "github.com/johnmckenna/provider-dynatrace/internal/controller/clustermanagement/managedinternetproxy"
	managednetworkzones "github.com/johnmckenna/provider-dynatrace/internal/controller/clustermanagement/managednetworkzones"
	managedpublicendpoints "github.com/johnmckenna/provider-dynatrace/internal/controller/clustermanagement/managedpublicendpoints"
	managedremoteaccess "github.com/johnmckenna/provider-dynatrace/internal/controller/clustermanagement/managedremoteaccess"
	managedsmtp "github.com/johnmckenna/provider-dynatrace/internal/controller/clustermanagement/managedsmtp"
	containerbuiltinrule "github.com/johnmckenna/provider-dynatrace/internal/controller/containers/containerbuiltinrule"
	containerregistry "github.com/johnmckenna/provider-dynatrace/internal/controller/containers/containerregistry"
	containerrule "github.com/johnmckenna/provider-dynatrace/internal/controller/containers/containerrule"
	containertechnology "github.com/johnmckenna/provider-dynatrace/internal/controller/containers/containertechnology"
	awscredentials "github.com/johnmckenna/provider-dynatrace/internal/controller/credentials/awscredentials"
	awsservice "github.com/johnmckenna/provider-dynatrace/internal/controller/credentials/awsservice"
	azurecredentials "github.com/johnmckenna/provider-dynatrace/internal/controller/credentials/azurecredentials"
	azureservice "github.com/johnmckenna/provider-dynatrace/internal/controller/credentials/azureservice"
	credentials "github.com/johnmckenna/provider-dynatrace/internal/controller/credentials/credentials"
	dashboardsallowlist "github.com/johnmckenna/provider-dynatrace/internal/controller/dashboards/dashboardsallowlist"
	dashboardsgeneral "github.com/johnmckenna/provider-dynatrace/internal/controller/dashboards/dashboardsgeneral"
	dashboardspresets "github.com/johnmckenna/provider-dynatrace/internal/controller/dashboards/dashboardspresets"
	jsondashboard "github.com/johnmckenna/provider-dynatrace/internal/controller/dashboards/jsondashboard"
	jsondashboardbase "github.com/johnmckenna/provider-dynatrace/internal/controller/dashboards/jsondashboardbase"
	alertingprofile "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/alertingprofile"
	applicationanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/applicationanomalies"
	autotag "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/autotag"
	cloudfoundrycredentials "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/cloudfoundrycredentials"
	customanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/customanomalies"
	dashboard "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/dashboard"
	databaseanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/databaseanomalies"
	ddupool "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/ddupool"
	diskanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/diskanomalies"
	hostanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/hostanomalies"
	k8scredentials "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/k8scredentials"
	maintenancewindow "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/maintenancewindow"
	managementzone "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/managementzone"
	notification "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/notification"
	pganomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/pganomalies"
	resourceattributes "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/resourceattributes"
	serviceanomalies "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/serviceanomalies"
	slo "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/slo"
	spanevents "github.com/johnmckenna/provider-dynatrace/internal/controller/deprecated/spanevents"
	devobsagentoptin "github.com/johnmckenna/provider-dynatrace/internal/controller/developerobservability/devobsagentoptin"
	devobsdatamasking "github.com/johnmckenna/provider-dynatrace/internal/controller/developerobservability/devobsdatamasking"
	devobsgitonprem "github.com/johnmckenna/provider-dynatrace/internal/controller/developerobservability/devobsgitonprem"
	directshares "github.com/johnmckenna/provider-dynatrace/internal/controller/documents/directshares"
	document "github.com/johnmckenna/provider-dynatrace/internal/controller/documents/document"
	appmonitoring "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/appmonitoring"
	auditlog "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/auditlog"
	dataprivacy "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/dataprivacy"
	diskoptions "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/diskoptions"
	eulasettings "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/eulasettings"
	hubpermissions "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/hubpermissions"
	hubsubscriptions "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/hubsubscriptions"
	ipaddressmasking "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/ipaddressmasking"
	limitoutboundconnections "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/limitoutboundconnections"
	oneagentdefaultmode "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/oneagentdefaultmode"
	oneagentdefaultversion "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/oneagentdefaultversion"
	oneagentfeatures "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/oneagentfeatures"
	oneagentsidemasking "github.com/johnmckenna/provider-dynatrace/internal/controller/environmentsettings/oneagentsidemasking"
	extensionexecutioncontroller "github.com/johnmckenna/provider-dynatrace/internal/controller/extensions/extensionexecutioncontroller"
	extensionexecutionremote "github.com/johnmckenna/provider-dynatrace/internal/controller/extensions/extensionexecutionremote"
	hubextensionactiveversion "github.com/johnmckenna/provider-dynatrace/internal/controller/extensions/hubextensionactiveversion"
	hubextensionconfig "github.com/johnmckenna/provider-dynatrace/internal/controller/extensions/hubextensionconfig"
	failuredetectionparameters "github.com/johnmckenna/provider-dynatrace/internal/controller/failuredetection/failuredetectionparameters"
	failuredetectionrules "github.com/johnmckenna/provider-dynatrace/internal/controller/failuredetection/failuredetectionrules"
	servicefailure "github.com/johnmckenna/provider-dynatrace/internal/controller/failuredetection/servicefailure"
	servicehttpfailure "github.com/johnmckenna/provider-dynatrace/internal/controller/failuredetection/servicehttpfailure"
	aixextension "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/aixextension"
	crashdumpanalytics "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/crashdumpanalytics"
	diskanalytics "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/diskanalytics"
	ebpfservicediscovery "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/ebpfservicediscovery"
	hostmonitoring "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/hostmonitoring"
	hostmonitoringadvanced "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/hostmonitoringadvanced"
	hostmonitoringmode "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/hostmonitoringmode"
	hostprocessgroupmonitoring "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/hostprocessgroupmonitoring"
	nettracer "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/nettracer"
	networktraffic "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/networktraffic"
	osservices "github.com/johnmckenna/provider-dynatrace/internal/controller/hostmonitoring/osservices"
	httpmonitor "github.com/johnmckenna/provider-dynatrace/internal/controller/httpmonitors/httpmonitor"
	httpmonitorcookies "github.com/johnmckenna/provider-dynatrace/internal/controller/httpmonitors/httpmonitorcookies"
	httpmonitoroutage "github.com/johnmckenna/provider-dynatrace/internal/controller/httpmonitors/httpmonitoroutage"
	httpmonitorperformance "github.com/johnmckenna/provider-dynatrace/internal/controller/httpmonitors/httpmonitorperformance"
	httpmonitorscript "github.com/johnmckenna/provider-dynatrace/internal/controller/httpmonitors/httpmonitorscript"
	iamgroup "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/iamgroup"
	iampermission "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/iampermission"
	iampolicy "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/iampolicy"
	iampolicybindings "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/iampolicybindings"
	iampolicybindingsv2 "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/iampolicybindingsv2"
	iamuser "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/iamuser"
	mgmzpermission "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/mgmzpermission"
	policy "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/policy"
	policybindings "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/policybindings"
	user "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/user"
	usergroup "github.com/johnmckenna/provider-dynatrace/internal/controller/iam/usergroup"
	goldenstate "github.com/johnmckenna/provider-dynatrace/internal/controller/incubator/goldenstate"
	issuetracking "github.com/johnmckenna/provider-dynatrace/internal/controller/integrations/issuetracking"
	remoteenvironments "github.com/johnmckenna/provider-dynatrace/internal/controller/integrations/remoteenvironments"
	logbuckets "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logbuckets"
	logcustomattribute "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logcustomattribute"
	logcustomsource "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logcustomsource"
	logdebugsettings "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logdebugsettings"
	logevents "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logevents"
	logmetrics "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logmetrics"
	logoneagent "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logoneagent"
	logprocessing "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logprocessing"
	logsecuritycontext "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logsecuritycontext"
	logsensitivedatamasking "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logsensitivedatamasking"
	logstorage "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logstorage"
	logtimestamp "github.com/johnmckenna/provider-dynatrace/internal/controller/logmonitoring/logtimestamp"
	ibmmqfilters "github.com/johnmckenna/provider-dynatrace/internal/controller/mainframe/ibmmqfilters"
	imsbridges "github.com/johnmckenna/provider-dynatrace/internal/controller/mainframe/imsbridges"
	mainframetransactionmonitoring "github.com/johnmckenna/provider-dynatrace/internal/controller/mainframe/mainframetransactionmonitoring"
	queuemanager "github.com/johnmckenna/provider-dynatrace/internal/controller/mainframe/queuemanager"
	queuesharinggroups "github.com/johnmckenna/provider-dynatrace/internal/controller/mainframe/queuesharinggroups"
	transactionstartfilters "github.com/johnmckenna/provider-dynatrace/internal/controller/mainframe/transactionstartfilters"
	customunits "github.com/johnmckenna/provider-dynatrace/internal/controller/metrics/customunits"
	histogrammetrics "github.com/johnmckenna/provider-dynatrace/internal/controller/metrics/histogrammetrics"
	metricmetadata "github.com/johnmckenna/provider-dynatrace/internal/controller/metrics/metricmetadata"
	metricquery "github.com/johnmckenna/provider-dynatrace/internal/controller/metrics/metricquery"
	calculatedmobilemetric "github.com/johnmckenna/provider-dynatrace/internal/controller/mobilecustomapplications/calculatedmobilemetric"
	customappenablement "github.com/johnmckenna/provider-dynatrace/internal/controller/mobilecustomapplications/customappenablement"
	mobileappenablement "github.com/johnmckenna/provider-dynatrace/internal/controller/mobilecustomapplications/mobileappenablement"
	mobileappkeyperformance "github.com/johnmckenna/provider-dynatrace/internal/controller/mobilecustomapplications/mobileappkeyperformance"
	mobileapplication "github.com/johnmckenna/provider-dynatrace/internal/controller/mobilecustomapplications/mobileapplication"
	mobileapprequesterrors "github.com/johnmckenna/provider-dynatrace/internal/controller/mobilecustomapplications/mobileapprequesterrors"
	customdevice "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredentities/customdevice"
	monitoredtechnologiesapache "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesapache"
	monitoredtechnologiesdotnet "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesdotnet"
	monitoredtechnologiesgo "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesgo"
	monitoredtechnologiesiis "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesiis"
	monitoredtechnologiesjava "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesjava"
	monitoredtechnologiesnginx "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesnginx"
	monitoredtechnologiesnodejs "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesnodejs"
	monitoredtechnologiesopentracing "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesopentracing"
	monitoredtechnologiesphp "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesphp"
	monitoredtechnologiesvarnish "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologiesvarnish"
	monitoredtechnologieswsmb "github.com/johnmckenna/provider-dynatrace/internal/controller/monitoredtechnologies/monitoredtechnologieswsmb"
	networkmonitor "github.com/johnmckenna/provider-dynatrace/internal/controller/networkavailabilitymonitors/networkmonitor"
	networkmonitoroutage "github.com/johnmckenna/provider-dynatrace/internal/controller/networkavailabilitymonitors/networkmonitoroutage"
	alerting "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/alerting"
	ansibletowernotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/ansibletowernotification"
	emailnotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/emailnotification"
	jiranotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/jiranotification"
	mobilenotifications "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/mobilenotifications"
	opsgenienotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/opsgenienotification"
	pagerdutynotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/pagerdutynotification"
	servicenownotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/servicenownotification"
	slacknotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/slacknotification"
	trellonotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/trellonotification"
	victoropsnotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/victoropsnotification"
	webhooknotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/webhooknotification"
	xmattersnotification "github.com/johnmckenna/provider-dynatrace/internal/controller/notifications/xmattersnotification"
	attributeallowlist "github.com/johnmckenna/provider-dynatrace/internal/controller/opentelemetryopentracing/attributeallowlist"
	attributeblocklist "github.com/johnmckenna/provider-dynatrace/internal/controller/opentelemetryopentracing/attributeblocklist"
	attributemasking "github.com/johnmckenna/provider-dynatrace/internal/controller/opentelemetryopentracing/attributemasking"
	attributespreferences "github.com/johnmckenna/provider-dynatrace/internal/controller/opentelemetryopentracing/attributespreferences"
	opentelemetrymetrics "github.com/johnmckenna/provider-dynatrace/internal/controller/opentelemetryopentracing/opentelemetrymetrics"
	spancapturerule "github.com/johnmckenna/provider-dynatrace/internal/controller/opentelemetryopentracing/spancapturerule"
	spancontextpropagation "github.com/johnmckenna/provider-dynatrace/internal/controller/opentelemetryopentracing/spancontextpropagation"
	spanentrypoint "github.com/johnmckenna/provider-dynatrace/internal/controller/opentelemetryopentracing/spanentrypoint"
	ownershipconfig "github.com/johnmckenna/provider-dynatrace/internal/controller/ownership/ownershipconfig"
	ownershipteams "github.com/johnmckenna/provider-dynatrace/internal/controller/ownership/ownershipteams"
	genericsetting "github.com/johnmckenna/provider-dynatrace/internal/controller/platform/genericsetting"
	grailmetricsallowall "github.com/johnmckenna/provider-dynatrace/internal/controller/platform/grailmetricsallowall"
	grailmetricsallowlist "github.com/johnmckenna/provider-dynatrace/internal/controller/platform/grailmetricsallowlist"
	platformbucket "github.com/johnmckenna/provider-dynatrace/internal/controller/platform/platformbucket"
	cloudappworkloaddetection "github.com/johnmckenna/provider-dynatrace/internal/controller/processgroupmonitoring/cloudappworkloaddetection"
	declarativegrouping "github.com/johnmckenna/provider-dynatrace/internal/controller/processgroupmonitoring/declarativegrouping"
	pgalerting "github.com/johnmckenna/provider-dynatrace/internal/controller/processgroupmonitoring/pgalerting"
	processgroupdetection "github.com/johnmckenna/provider-dynatrace/internal/controller/processgroupmonitoring/processgroupdetection"
	processgroupdetectionflags "github.com/johnmckenna/provider-dynatrace/internal/controller/processgroupmonitoring/processgroupdetectionflags"
	processgroupmonitoring "github.com/johnmckenna/provider-dynatrace/internal/controller/processgroupmonitoring/processgroupmonitoring"
	processgroupnaming "github.com/johnmckenna/provider-dynatrace/internal/controller/processgroupmonitoring/processgroupnaming"
	processgroupsimpledetection "github.com/johnmckenna/provider-dynatrace/internal/controller/processgroupmonitoring/processgroupsimpledetection"
	builtinprocessmonitoring "github.com/johnmckenna/provider-dynatrace/internal/controller/processmonitoring/builtinprocessmonitoring"
	processavailability "github.com/johnmckenna/provider-dynatrace/internal/controller/processmonitoring/processavailability"
	processmonitoring "github.com/johnmckenna/provider-dynatrace/internal/controller/processmonitoring/processmonitoring"
	processmonitoringrule "github.com/johnmckenna/provider-dynatrace/internal/controller/processmonitoring/processmonitoringrule"
	processvisibility "github.com/johnmckenna/provider-dynatrace/internal/controller/processmonitoring/processvisibility"
	providerconfig "github.com/johnmckenna/provider-dynatrace/internal/controller/providerconfig"
	geolocation "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/geolocation"
	processgrouprum "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/processgrouprum"
	rumadvancedcorrelation "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/rumadvancedcorrelation"
	rumhostheaders "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/rumhostheaders"
	rumipdetermination "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/rumipdetermination"
	rumiplocations "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/rumiplocations"
	rumoverloadprevention "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/rumoverloadprevention"
	rumproviderbreakdown "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/rumproviderbreakdown"
	usabilityanalytics "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/usabilityanalytics"
	useractionmetrics "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/useractionmetrics"
	userexperiencescore "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/userexperiencescore"
	usersessionmetrics "github.com/johnmckenna/provider-dynatrace/internal/controller/realusermonitoring/usersessionmetrics"
	customservice "github.com/johnmckenna/provider-dynatrace/internal/controller/servicedetection/customservice"
	managementzonev2 "github.com/johnmckenna/provider-dynatrace/internal/controller/servicedetection/managementzonev2"
	serviceexternalwebrequest "github.com/johnmckenna/provider-dynatrace/internal/controller/servicedetection/serviceexternalwebrequest"
	serviceexternalwebservice "github.com/johnmckenna/provider-dynatrace/internal/controller/servicedetection/serviceexternalwebservice"
	servicefullwebrequest "github.com/johnmckenna/provider-dynatrace/internal/controller/servicedetection/servicefullwebrequest"
	servicefullwebservice "github.com/johnmckenna/provider-dynatrace/internal/controller/servicedetection/servicefullwebservice"
	unifiedservicesmetrics "github.com/johnmckenna/provider-dynatrace/internal/controller/servicedetection/unifiedservicesmetrics"
	unifiedservicesopentel "github.com/johnmckenna/provider-dynatrace/internal/controller/servicedetection/unifiedservicesopentel"
	slonormalization "github.com/johnmckenna/provider-dynatrace/internal/controller/servicelevelobjective/slonormalization"
	slov2 "github.com/johnmckenna/provider-dynatrace/internal/controller/servicelevelobjective/slov2"
	apidetection "github.com/johnmckenna/provider-dynatrace/internal/controller/servicemonitoring/apidetection"
	calculatedservicemetric "github.com/johnmckenna/provider-dynatrace/internal/controller/servicemonitoring/calculatedservicemetric"
	mutedrequests "github.com/johnmckenna/provider-dynatrace/internal/controller/servicemonitoring/mutedrequests"
	requestattribute "github.com/johnmckenna/provider-dynatrace/internal/controller/servicemonitoring/requestattribute"
	servicenaming "github.com/johnmckenna/provider-dynatrace/internal/controller/servicemonitoring/servicenaming"
	urlbasedsampling "github.com/johnmckenna/provider-dynatrace/internal/controller/servicemonitoring/urlbasedsampling"
	daviscopilot "github.com/johnmckenna/provider-dynatrace/internal/controller/servicesettings/daviscopilot"
	sessionreplayresourcecapture "github.com/johnmckenna/provider-dynatrace/internal/controller/sessionreplay/sessionreplayresourcecapture"
	sessionreplaywebprivacy "github.com/johnmckenna/provider-dynatrace/internal/controller/sessionreplay/sessionreplaywebprivacy"
	calculatedsyntheticmetric "github.com/johnmckenna/provider-dynatrace/internal/controller/synthetic/calculatedsyntheticmetric"
	syntheticavailability "github.com/johnmckenna/provider-dynatrace/internal/controller/synthetic/syntheticavailability"
	syntheticlocation "github.com/johnmckenna/provider-dynatrace/internal/controller/synthetic/syntheticlocation"
	autotagrules "github.com/johnmckenna/provider-dynatrace/internal/controller/tags/autotagrules"
	autotagv2 "github.com/johnmckenna/provider-dynatrace/internal/controller/tags/autotagv2"
	customtags "github.com/johnmckenna/provider-dynatrace/internal/controller/tags/customtags"
	genericrelationships "github.com/johnmckenna/provider-dynatrace/internal/controller/topologymodel/genericrelationships"
	generictypes "github.com/johnmckenna/provider-dynatrace/internal/controller/topologymodel/generictypes"
	grailsecuritycontext "github.com/johnmckenna/provider-dynatrace/internal/controller/topologymodel/grailsecuritycontext"
	activegateupdates "github.com/johnmckenna/provider-dynatrace/internal/controller/updates/activegateupdates"
	oneagentupdates "github.com/johnmckenna/provider-dynatrace/internal/controller/updates/oneagentupdates"
	updatewindows "github.com/johnmckenna/provider-dynatrace/internal/controller/updates/updatewindows"
	usersettings "github.com/johnmckenna/provider-dynatrace/internal/controller/usersettings/usersettings"
	vmware "github.com/johnmckenna/provider-dynatrace/internal/controller/virtualization/vmware"
	applicationdetectionrulev2 "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/applicationdetectionrulev2"
	calculatedwebmetric "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/calculatedwebmetric"
	keyuseraction "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/keyuseraction"
	webappbeaconendpoint "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappbeaconendpoint"
	webappbeaconorigins "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappbeaconorigins"
	webappcustomconfigproperties "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappcustomconfigproperties"
	webappcustomerrors "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappcustomerrors"
	webappcustominjection "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappcustominjection"
	webappenablement "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappenablement"
	webappinjectioncookie "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappinjectioncookie"
	webappjavascriptupdates "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappjavascriptupdates"
	webappjavascriptversion "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappjavascriptversion"
	webappkeyperformancecustom "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappkeyperformancecustom"
	webappkeyperformanceload "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappkeyperformanceload"
	webappkeyperformancexhr "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappkeyperformancexhr"
	webapplication "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webapplication"
	webapprequesterrors "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webapprequesterrors"
	webappresourcecleanup "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappresourcecleanup"
	webappresourcetypes "github.com/johnmckenna/provider-dynatrace/internal/controller/webapplications/webappresourcetypes"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activegatetoken.Setup,
		agtoken.Setup,
		apitoken.Setup,
		tokensettings.Setup,
		connectivityalerts.Setup,
		maintenance.Setup,
		awsanomalies.Setup,
		customappanomalies.Setup,
		customappcrashrate.Setup,
		databaseanomaliesv2.Setup,
		davisanomalydetectors.Setup,
		diskanomaliesv2.Setup,
		diskanomalyrules.Setup,
		diskedgeanomalydetectors.Setup,
		diskspecificanomaliesv2.Setup,
		frequentissues.Setup,
		hostanomaliesv2.Setup,
		k8sclusteranomalies.Setup,
		k8snamespaceanomalies.Setup,
		k8snodeanomalies.Setup,
		k8spvcanomalies.Setup,
		k8sworkloadanomalies.Setup,
		metricevents.Setup,
		mobileappanomalies.Setup,
		mobileappcrashrate.Setup,
		serviceanomaliesv2.Setup,
		vmwareanomalies.Setup,
		webappanomalies.Setup,
		dbappfeatureflags.Setup,
		discoverydefaultrules.Setup,
		discoveryfeatureflags.Setup,
		infraopsappfeatureflags.Setup,
		infraopsappsettings.Setup,
		appsecnotification.Setup,
		attackalerting.Setup,
		attackallowlist.Setup,
		attackrules.Setup,
		attacksettings.Setup,
		vulnerabilityalerting.Setup,
		vulnerabilitycode.Setup,
		vulnerabilitysettings.Setup,
		vulnerabilitythirdparty.Setup,
		automationbusinesscalendar.Setup,
		automationschedulingrule.Setup,
		automationworkflow.Setup,
		automationworkflowawsconnections.Setup,
		automationworkflowjira.Setup,
		automationworkflowk8sconnections.Setup,
		automationworkflowslack.Setup,
		sitereliabilityguardian.Setup,
		browsermonitor.Setup,
		browsermonitoroutage.Setup,
		browsermonitorperformance.Setup,
		businesseventsbuckets.Setup,
		businesseventsmetrics.Setup,
		businesseventsoneagent.Setup,
		businesseventsoneagentoutgoing.Setup,
		businesseventsprocessing.Setup,
		businesseventssecuritycontext.Setup,
		cloudfoundry.Setup,
		k8smonitoring.Setup,
		kubernetes.Setup,
		kubernetesapp.Setup,
		kubernetesenrichment.Setup,
		managedbackup.Setup,
		managedinternetproxy.Setup,
		managednetworkzones.Setup,
		managedpublicendpoints.Setup,
		managedremoteaccess.Setup,
		managedsmtp.Setup,
		containerbuiltinrule.Setup,
		containerregistry.Setup,
		containerrule.Setup,
		containertechnology.Setup,
		awscredentials.Setup,
		awsservice.Setup,
		azurecredentials.Setup,
		azureservice.Setup,
		credentials.Setup,
		dashboardsallowlist.Setup,
		dashboardsgeneral.Setup,
		dashboardspresets.Setup,
		jsondashboard.Setup,
		jsondashboardbase.Setup,
		alertingprofile.Setup,
		applicationanomalies.Setup,
		autotag.Setup,
		cloudfoundrycredentials.Setup,
		customanomalies.Setup,
		dashboard.Setup,
		databaseanomalies.Setup,
		ddupool.Setup,
		diskanomalies.Setup,
		hostanomalies.Setup,
		k8scredentials.Setup,
		maintenancewindow.Setup,
		managementzone.Setup,
		notification.Setup,
		pganomalies.Setup,
		resourceattributes.Setup,
		serviceanomalies.Setup,
		slo.Setup,
		spanevents.Setup,
		devobsagentoptin.Setup,
		devobsdatamasking.Setup,
		devobsgitonprem.Setup,
		directshares.Setup,
		document.Setup,
		appmonitoring.Setup,
		auditlog.Setup,
		dataprivacy.Setup,
		diskoptions.Setup,
		eulasettings.Setup,
		hubpermissions.Setup,
		hubsubscriptions.Setup,
		ipaddressmasking.Setup,
		limitoutboundconnections.Setup,
		oneagentdefaultmode.Setup,
		oneagentdefaultversion.Setup,
		oneagentfeatures.Setup,
		oneagentsidemasking.Setup,
		extensionexecutioncontroller.Setup,
		extensionexecutionremote.Setup,
		hubextensionactiveversion.Setup,
		hubextensionconfig.Setup,
		failuredetectionparameters.Setup,
		failuredetectionrules.Setup,
		servicefailure.Setup,
		servicehttpfailure.Setup,
		aixextension.Setup,
		crashdumpanalytics.Setup,
		diskanalytics.Setup,
		ebpfservicediscovery.Setup,
		hostmonitoring.Setup,
		hostmonitoringadvanced.Setup,
		hostmonitoringmode.Setup,
		hostprocessgroupmonitoring.Setup,
		nettracer.Setup,
		networktraffic.Setup,
		osservices.Setup,
		httpmonitor.Setup,
		httpmonitorcookies.Setup,
		httpmonitoroutage.Setup,
		httpmonitorperformance.Setup,
		httpmonitorscript.Setup,
		iamgroup.Setup,
		iampermission.Setup,
		iampolicy.Setup,
		iampolicybindings.Setup,
		iampolicybindingsv2.Setup,
		iamuser.Setup,
		mgmzpermission.Setup,
		policy.Setup,
		policybindings.Setup,
		user.Setup,
		usergroup.Setup,
		goldenstate.Setup,
		issuetracking.Setup,
		remoteenvironments.Setup,
		logbuckets.Setup,
		logcustomattribute.Setup,
		logcustomsource.Setup,
		logdebugsettings.Setup,
		logevents.Setup,
		logmetrics.Setup,
		logoneagent.Setup,
		logprocessing.Setup,
		logsecuritycontext.Setup,
		logsensitivedatamasking.Setup,
		logstorage.Setup,
		logtimestamp.Setup,
		ibmmqfilters.Setup,
		imsbridges.Setup,
		mainframetransactionmonitoring.Setup,
		queuemanager.Setup,
		queuesharinggroups.Setup,
		transactionstartfilters.Setup,
		customunits.Setup,
		histogrammetrics.Setup,
		metricmetadata.Setup,
		metricquery.Setup,
		calculatedmobilemetric.Setup,
		customappenablement.Setup,
		mobileappenablement.Setup,
		mobileappkeyperformance.Setup,
		mobileapplication.Setup,
		mobileapprequesterrors.Setup,
		customdevice.Setup,
		monitoredtechnologiesapache.Setup,
		monitoredtechnologiesdotnet.Setup,
		monitoredtechnologiesgo.Setup,
		monitoredtechnologiesiis.Setup,
		monitoredtechnologiesjava.Setup,
		monitoredtechnologiesnginx.Setup,
		monitoredtechnologiesnodejs.Setup,
		monitoredtechnologiesopentracing.Setup,
		monitoredtechnologiesphp.Setup,
		monitoredtechnologiesvarnish.Setup,
		monitoredtechnologieswsmb.Setup,
		networkmonitor.Setup,
		networkmonitoroutage.Setup,
		alerting.Setup,
		ansibletowernotification.Setup,
		emailnotification.Setup,
		jiranotification.Setup,
		mobilenotifications.Setup,
		opsgenienotification.Setup,
		pagerdutynotification.Setup,
		servicenownotification.Setup,
		slacknotification.Setup,
		trellonotification.Setup,
		victoropsnotification.Setup,
		webhooknotification.Setup,
		xmattersnotification.Setup,
		attributeallowlist.Setup,
		attributeblocklist.Setup,
		attributemasking.Setup,
		attributespreferences.Setup,
		opentelemetrymetrics.Setup,
		spancapturerule.Setup,
		spancontextpropagation.Setup,
		spanentrypoint.Setup,
		ownershipconfig.Setup,
		ownershipteams.Setup,
		genericsetting.Setup,
		grailmetricsallowall.Setup,
		grailmetricsallowlist.Setup,
		platformbucket.Setup,
		cloudappworkloaddetection.Setup,
		declarativegrouping.Setup,
		pgalerting.Setup,
		processgroupdetection.Setup,
		processgroupdetectionflags.Setup,
		processgroupmonitoring.Setup,
		processgroupnaming.Setup,
		processgroupsimpledetection.Setup,
		builtinprocessmonitoring.Setup,
		processavailability.Setup,
		processmonitoring.Setup,
		processmonitoringrule.Setup,
		processvisibility.Setup,
		providerconfig.Setup,
		geolocation.Setup,
		processgrouprum.Setup,
		rumadvancedcorrelation.Setup,
		rumhostheaders.Setup,
		rumipdetermination.Setup,
		rumiplocations.Setup,
		rumoverloadprevention.Setup,
		rumproviderbreakdown.Setup,
		usabilityanalytics.Setup,
		useractionmetrics.Setup,
		userexperiencescore.Setup,
		usersessionmetrics.Setup,
		customservice.Setup,
		managementzonev2.Setup,
		serviceexternalwebrequest.Setup,
		serviceexternalwebservice.Setup,
		servicefullwebrequest.Setup,
		servicefullwebservice.Setup,
		unifiedservicesmetrics.Setup,
		unifiedservicesopentel.Setup,
		slonormalization.Setup,
		slov2.Setup,
		apidetection.Setup,
		calculatedservicemetric.Setup,
		mutedrequests.Setup,
		requestattribute.Setup,
		servicenaming.Setup,
		urlbasedsampling.Setup,
		daviscopilot.Setup,
		sessionreplayresourcecapture.Setup,
		sessionreplaywebprivacy.Setup,
		calculatedsyntheticmetric.Setup,
		syntheticavailability.Setup,
		syntheticlocation.Setup,
		autotagrules.Setup,
		autotagv2.Setup,
		customtags.Setup,
		genericrelationships.Setup,
		generictypes.Setup,
		grailsecuritycontext.Setup,
		activegateupdates.Setup,
		oneagentupdates.Setup,
		updatewindows.Setup,
		usersettings.Setup,
		vmware.Setup,
		applicationdetectionrulev2.Setup,
		calculatedwebmetric.Setup,
		keyuseraction.Setup,
		webappbeaconendpoint.Setup,
		webappbeaconorigins.Setup,
		webappcustomconfigproperties.Setup,
		webappcustomerrors.Setup,
		webappcustominjection.Setup,
		webappenablement.Setup,
		webappinjectioncookie.Setup,
		webappjavascriptupdates.Setup,
		webappjavascriptversion.Setup,
		webappkeyperformancecustom.Setup,
		webappkeyperformanceload.Setup,
		webappkeyperformancexhr.Setup,
		webapplication.Setup,
		webapprequesterrors.Setup,
		webappresourcecleanup.Setup,
		webappresourcetypes.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
