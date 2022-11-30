package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Userable 
type Userable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAboutMe()(*string)
    GetAccountEnabled()(*bool)
    GetActivities()([]UserActivityable)
    GetAgeGroup()(*string)
    GetAgreementAcceptances()([]AgreementAcceptanceable)
    GetAppRoleAssignments()([]AppRoleAssignmentable)
    GetAssignedLicenses()([]AssignedLicenseable)
    GetAssignedPlans()([]AssignedPlanable)
    GetAuthentication()(Authenticationable)
    GetAuthorizationInfo()(AuthorizationInfoable)
    GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBusinessPhones()([]string)
    GetCalendar()(Calendarable)
    GetCalendarGroups()([]CalendarGroupable)
    GetCalendars()([]Calendarable)
    GetCalendarView()([]Eventable)
    GetChats()([]Chatable)
    GetCity()(*string)
    GetCompanyName()(*string)
    GetConsentProvidedForMinor()(*string)
    GetContactFolders()([]ContactFolderable)
    GetContacts()([]Contactable)
    GetCountry()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedObjects()([]DirectoryObjectable)
    GetCreationType()(*string)
    GetDepartment()(*string)
    GetDeviceEnrollmentLimit()(*int32)
    GetDeviceManagementTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable)
    GetDirectReports()([]DirectoryObjectable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetDrives()([]Driveable)
    GetEmployeeHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmployeeId()(*string)
    GetEmployeeOrgData()(EmployeeOrgDataable)
    GetEmployeeType()(*string)
    GetEvents()([]Eventable)
    GetExtensions()([]Extensionable)
    GetExternalUserState()(*string)
    GetExternalUserStateChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFaxNumber()(*string)
    GetFollowedSites()([]Siteable)
    GetGivenName()(*string)
    GetHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIdentities()([]ObjectIdentityable)
    GetImAddresses()([]string)
    GetInferenceClassification()(InferenceClassificationable)
    GetInsights()(OfficeGraphInsightsable)
    GetInterests()([]string)
    GetIsResourceAccount()(*bool)
    GetJobTitle()(*string)
    GetJoinedTeams()([]Teamable)
    GetLastPasswordChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLegalAgeGroupClassification()(*string)
    GetLicenseAssignmentStates()([]LicenseAssignmentStateable)
    GetLicenseDetails()([]LicenseDetailsable)
    GetMail()(*string)
    GetMailboxSettings()(MailboxSettingsable)
    GetMailFolders()([]MailFolderable)
    GetMailNickname()(*string)
    GetManagedAppRegistrations()([]ManagedAppRegistrationable)
    GetManagedDevices()([]ManagedDeviceable)
    GetManager()(DirectoryObjectable)
    GetMemberOf()([]DirectoryObjectable)
    GetMessages()([]Messageable)
    GetMobilePhone()(*string)
    GetMySite()(*string)
    GetOauth2PermissionGrants()([]OAuth2PermissionGrantable)
    GetOfficeLocation()(*string)
    GetOnenote()(Onenoteable)
    GetOnlineMeetings()([]OnlineMeetingable)
    GetOnPremisesDistinguishedName()(*string)
    GetOnPremisesDomainName()(*string)
    GetOnPremisesExtensionAttributes()(OnPremisesExtensionAttributesable)
    GetOnPremisesImmutableId()(*string)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable)
    GetOnPremisesSamAccountName()(*string)
    GetOnPremisesSecurityIdentifier()(*string)
    GetOnPremisesSyncEnabled()(*bool)
    GetOnPremisesUserPrincipalName()(*string)
    GetOtherMails()([]string)
    GetOutlook()(OutlookUserable)
    GetOwnedDevices()([]DirectoryObjectable)
    GetOwnedObjects()([]DirectoryObjectable)
    GetPasswordPolicies()(*string)
    GetPasswordProfile()(PasswordProfileable)
    GetPastProjects()([]string)
    GetPeople()([]Personable)
    GetPhoto()(ProfilePhotoable)
    GetPhotos()([]ProfilePhotoable)
    GetPlanner()(PlannerUserable)
    GetPostalCode()(*string)
    GetPreferredDataLocation()(*string)
    GetPreferredLanguage()(*string)
    GetPreferredName()(*string)
    GetPresence()(Presenceable)
    GetProvisionedPlans()([]ProvisionedPlanable)
    GetProxyAddresses()([]string)
    GetRegisteredDevices()([]DirectoryObjectable)
    GetResponsibilities()([]string)
    GetSchools()([]string)
    GetScopedRoleMemberOf()([]ScopedRoleMembershipable)
    GetSecurityIdentifier()(*string)
    GetSettings()(UserSettingsable)
    GetShowInAddressList()(*bool)
    GetSignInSessionsValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSkills()([]string)
    GetState()(*string)
    GetStreetAddress()(*string)
    GetSurname()(*string)
    GetTeamwork()(UserTeamworkable)
    GetTodo()(Todoable)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    GetUsageLocation()(*string)
    GetUserPrincipalName()(*string)
    GetUserType()(*string)
    SetAboutMe(value *string)()
    SetAccountEnabled(value *bool)()
    SetActivities(value []UserActivityable)()
    SetAgeGroup(value *string)()
    SetAgreementAcceptances(value []AgreementAcceptanceable)()
    SetAppRoleAssignments(value []AppRoleAssignmentable)()
    SetAssignedLicenses(value []AssignedLicenseable)()
    SetAssignedPlans(value []AssignedPlanable)()
    SetAuthentication(value Authenticationable)()
    SetAuthorizationInfo(value AuthorizationInfoable)()
    SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBusinessPhones(value []string)()
    SetCalendar(value Calendarable)()
    SetCalendarGroups(value []CalendarGroupable)()
    SetCalendars(value []Calendarable)()
    SetCalendarView(value []Eventable)()
    SetChats(value []Chatable)()
    SetCity(value *string)()
    SetCompanyName(value *string)()
    SetConsentProvidedForMinor(value *string)()
    SetContactFolders(value []ContactFolderable)()
    SetContacts(value []Contactable)()
    SetCountry(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedObjects(value []DirectoryObjectable)()
    SetCreationType(value *string)()
    SetDepartment(value *string)()
    SetDeviceEnrollmentLimit(value *int32)()
    SetDeviceManagementTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)()
    SetDirectReports(value []DirectoryObjectable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetDrives(value []Driveable)()
    SetEmployeeHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmployeeId(value *string)()
    SetEmployeeOrgData(value EmployeeOrgDataable)()
    SetEmployeeType(value *string)()
    SetEvents(value []Eventable)()
    SetExtensions(value []Extensionable)()
    SetExternalUserState(value *string)()
    SetExternalUserStateChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFaxNumber(value *string)()
    SetFollowedSites(value []Siteable)()
    SetGivenName(value *string)()
    SetHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIdentities(value []ObjectIdentityable)()
    SetImAddresses(value []string)()
    SetInferenceClassification(value InferenceClassificationable)()
    SetInsights(value OfficeGraphInsightsable)()
    SetInterests(value []string)()
    SetIsResourceAccount(value *bool)()
    SetJobTitle(value *string)()
    SetJoinedTeams(value []Teamable)()
    SetLastPasswordChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLegalAgeGroupClassification(value *string)()
    SetLicenseAssignmentStates(value []LicenseAssignmentStateable)()
    SetLicenseDetails(value []LicenseDetailsable)()
    SetMail(value *string)()
    SetMailboxSettings(value MailboxSettingsable)()
    SetMailFolders(value []MailFolderable)()
    SetMailNickname(value *string)()
    SetManagedAppRegistrations(value []ManagedAppRegistrationable)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetManager(value DirectoryObjectable)()
    SetMemberOf(value []DirectoryObjectable)()
    SetMessages(value []Messageable)()
    SetMobilePhone(value *string)()
    SetMySite(value *string)()
    SetOauth2PermissionGrants(value []OAuth2PermissionGrantable)()
    SetOfficeLocation(value *string)()
    SetOnenote(value Onenoteable)()
    SetOnlineMeetings(value []OnlineMeetingable)()
    SetOnPremisesDistinguishedName(value *string)()
    SetOnPremisesDomainName(value *string)()
    SetOnPremisesExtensionAttributes(value OnPremisesExtensionAttributesable)()
    SetOnPremisesImmutableId(value *string)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)()
    SetOnPremisesSamAccountName(value *string)()
    SetOnPremisesSecurityIdentifier(value *string)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetOnPremisesUserPrincipalName(value *string)()
    SetOtherMails(value []string)()
    SetOutlook(value OutlookUserable)()
    SetOwnedDevices(value []DirectoryObjectable)()
    SetOwnedObjects(value []DirectoryObjectable)()
    SetPasswordPolicies(value *string)()
    SetPasswordProfile(value PasswordProfileable)()
    SetPastProjects(value []string)()
    SetPeople(value []Personable)()
    SetPhoto(value ProfilePhotoable)()
    SetPhotos(value []ProfilePhotoable)()
    SetPlanner(value PlannerUserable)()
    SetPostalCode(value *string)()
    SetPreferredDataLocation(value *string)()
    SetPreferredLanguage(value *string)()
    SetPreferredName(value *string)()
    SetPresence(value Presenceable)()
    SetProvisionedPlans(value []ProvisionedPlanable)()
    SetProxyAddresses(value []string)()
    SetRegisteredDevices(value []DirectoryObjectable)()
    SetResponsibilities(value []string)()
    SetSchools(value []string)()
    SetScopedRoleMemberOf(value []ScopedRoleMembershipable)()
    SetSecurityIdentifier(value *string)()
    SetSettings(value UserSettingsable)()
    SetShowInAddressList(value *bool)()
    SetSignInSessionsValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSkills(value []string)()
    SetState(value *string)()
    SetStreetAddress(value *string)()
    SetSurname(value *string)()
    SetTeamwork(value UserTeamworkable)()
    SetTodo(value Todoable)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
    SetUsageLocation(value *string)()
    SetUserPrincipalName(value *string)()
    SetUserType(value *string)()
}
