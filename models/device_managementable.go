package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementable 
type DeviceManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable)
    GetAuditEvents()([]AuditEventable)
    GetComplianceManagementPartners()([]ComplianceManagementPartnerable)
    GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable)
    GetDetectedApps()([]DetectedAppable)
    GetDeviceCategories()([]DeviceCategoryable)
    GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable)
    GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable)
    GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable)
    GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable)
    GetDeviceConfigurations()([]DeviceConfigurationable)
    GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable)
    GetDeviceManagementPartners()([]DeviceManagementPartnerable)
    GetExchangeConnectors()([]DeviceManagementExchangeConnectorable)
    GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable)
    GetIntuneAccountId()(*string)
    GetIntuneBrand()(IntuneBrandable)
    GetIosUpdateStatuses()([]IosUpdateDeviceStatusable)
    GetManagedDeviceOverview()(ManagedDeviceOverviewable)
    GetManagedDevices()([]ManagedDeviceable)
    GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable)
    GetNotificationMessageTemplates()([]NotificationMessageTemplateable)
    GetRemoteAssistancePartners()([]RemoteAssistancePartnerable)
    GetReports()(DeviceManagementReportsable)
    GetResourceOperations()([]ResourceOperationable)
    GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable)
    GetRoleDefinitions()([]RoleDefinitionable)
    GetSettings()(DeviceManagementSettingsable)
    GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable)
    GetSubscriptionState()(*DeviceManagementSubscriptionState)
    GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable)
    GetTermsAndConditions()([]TermsAndConditionsable)
    GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable)
    GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable)
    GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable)
    GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable)
    SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)()
    SetAuditEvents(value []AuditEventable)()
    SetComplianceManagementPartners(value []ComplianceManagementPartnerable)()
    SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)()
    SetDetectedApps(value []DetectedAppable)()
    SetDeviceCategories(value []DeviceCategoryable)()
    SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)()
    SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)()
    SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)()
    SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)()
    SetDeviceConfigurations(value []DeviceConfigurationable)()
    SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)()
    SetDeviceManagementPartners(value []DeviceManagementPartnerable)()
    SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)()
    SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)()
    SetIntuneAccountId(value *string)()
    SetIntuneBrand(value IntuneBrandable)()
    SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)()
    SetManagedDeviceOverview(value ManagedDeviceOverviewable)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)()
    SetNotificationMessageTemplates(value []NotificationMessageTemplateable)()
    SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)()
    SetReports(value DeviceManagementReportsable)()
    SetResourceOperations(value []ResourceOperationable)()
    SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)()
    SetRoleDefinitions(value []RoleDefinitionable)()
    SetSettings(value DeviceManagementSettingsable)()
    SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)()
    SetSubscriptionState(value *DeviceManagementSubscriptionState)()
    SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)()
    SetTermsAndConditions(value []TermsAndConditionsable)()
    SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)()
    SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)()
    SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)()
    SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)()
}
