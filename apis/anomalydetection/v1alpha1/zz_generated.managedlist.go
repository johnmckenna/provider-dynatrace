// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import resource "github.com/crossplane/crossplane-runtime/pkg/resource"

// GetItems of this AwsAnomaliesList.
func (l *AwsAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this CustomAppAnomaliesList.
func (l *CustomAppAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this CustomAppCrashRateList.
func (l *CustomAppCrashRateList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this DatabaseAnomaliesV2List.
func (l *DatabaseAnomaliesV2List) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this DavisAnomalyDetectorsList.
func (l *DavisAnomalyDetectorsList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this DiskAnomaliesV2List.
func (l *DiskAnomaliesV2List) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this DiskAnomalyRulesList.
func (l *DiskAnomalyRulesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this DiskEdgeAnomalyDetectorsList.
func (l *DiskEdgeAnomalyDetectorsList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this DiskSpecificAnomaliesV2List.
func (l *DiskSpecificAnomaliesV2List) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this FrequentIssuesList.
func (l *FrequentIssuesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this HostAnomaliesV2List.
func (l *HostAnomaliesV2List) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this K8SClusterAnomaliesList.
func (l *K8SClusterAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this K8SNamespaceAnomaliesList.
func (l *K8SNamespaceAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this K8SNodeAnomaliesList.
func (l *K8SNodeAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this K8SPvcAnomaliesList.
func (l *K8SPvcAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this K8SWorkloadAnomaliesList.
func (l *K8SWorkloadAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this MetricEventsList.
func (l *MetricEventsList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this MobileAppAnomaliesList.
func (l *MobileAppAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this MobileAppCrashRateList.
func (l *MobileAppCrashRateList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this ServiceAnomaliesV2List.
func (l *ServiceAnomaliesV2List) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this VmwareAnomaliesList.
func (l *VmwareAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this WebAppAnomaliesList.
func (l *WebAppAnomaliesList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}
