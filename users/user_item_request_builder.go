package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// UserItemRequestBuilder provides operations to manage the collection of user entities.
type UserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserItemRequestBuilderGetQueryParameters retrieve the properties and relationships of user object.
type UserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserItemRequestBuilderGetQueryParameters
}
// UserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Activities()(*ItemActivitiesRequestBuilder) {
    return NewItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AgreementAcceptances provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) AgreementAcceptances()(*ItemAgreementAcceptancesRequestBuilder) {
    return NewItemAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) AppRoleAssignments()(*ItemAppRoleAssignmentsRequestBuilder) {
    return NewItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignLicense provides operations to call the assignLicense method.
func (m *UserItemRequestBuilder) AssignLicense()(*ItemAssignLicenseRequestBuilder) {
    return NewItemAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Authentication provides operations to manage the authentication property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Authentication()(*ItemAuthenticationRequestBuilder) {
    return NewItemAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Calendar()(*ItemCalendarRequestBuilder) {
    return NewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarGroups provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CalendarGroups()(*ItemCalendarGroupsRequestBuilder) {
    return NewItemCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Calendars provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Calendars()(*ItemCalendarsRequestBuilder) {
    return NewItemCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CalendarView()(*ItemCalendarViewRequestBuilder) {
    return NewItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChangePassword provides operations to call the changePassword method.
func (m *UserItemRequestBuilder) ChangePassword()(*ItemChangePasswordRequestBuilder) {
    return NewItemChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Chats provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Chats()(*ItemChatsRequestBuilder) {
    return NewItemChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *UserItemRequestBuilder) CheckMemberGroups()(*ItemCheckMemberGroupsRequestBuilder) {
    return NewItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *UserItemRequestBuilder) CheckMemberObjects()(*ItemCheckMemberObjectsRequestBuilder) {
    return NewItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    m := &UserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ContactFolders()(*ItemContactFoldersRequestBuilder) {
    return NewItemContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Contacts provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Contacts()(*ItemContactsRequestBuilder) {
    return NewItemContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CreatedObjects()(*ItemCreatedObjectsRequestBuilder) {
    return NewItemCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-delete?view=graph-rest-1.0
func (m *UserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceManagementTroubleshootingEvents provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) DeviceManagementTroubleshootingEvents()(*ItemDeviceManagementTroubleshootingEventsRequestBuilder) {
    return NewItemDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectReports provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) DirectReports()(*ItemDirectReportsRequestBuilder) {
    return NewItemDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Drive()(*ItemDriveRequestBuilder) {
    return NewItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Drives()(*ItemDrivesRequestBuilder) {
    return NewItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Events provides operations to manage the events property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserItemRequestBuilder) ExportDeviceAndAppManagementData()(*ItemExportDeviceAndAppManagementDataRequestBuilder) {
    return NewItemExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserItemRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ItemExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return NewItemExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top)
}
// ExportPersonalData provides operations to call the exportPersonalData method.
func (m *UserItemRequestBuilder) ExportPersonalData()(*ItemExportPersonalDataRequestBuilder) {
    return NewItemExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Extensions()(*ItemExtensionsRequestBuilder) {
    return NewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FindMeetingTimes provides operations to call the findMeetingTimes method.
func (m *UserItemRequestBuilder) FindMeetingTimes()(*ItemFindMeetingTimesRequestBuilder) {
    return NewItemFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FollowedSites provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) FollowedSites()(*ItemFollowedSitesRequestBuilder) {
    return NewItemFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve the properties and relationships of user object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-get?view=graph-rest-1.0
func (m *UserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable), nil
}
// GetMailTips provides operations to call the getMailTips method.
func (m *UserItemRequestBuilder) GetMailTips()(*ItemGetMailTipsRequestBuilder) {
    return NewItemGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserItemRequestBuilder) GetManagedAppDiagnosticStatuses()(*ItemGetManagedAppDiagnosticStatusesRequestBuilder) {
    return NewItemGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserItemRequestBuilder) GetManagedAppPolicies()(*ItemGetManagedAppPoliciesRequestBuilder) {
    return NewItemGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserItemRequestBuilder) GetManagedDevicesWithAppFailures()(*ItemGetManagedDevicesWithAppFailuresRequestBuilder) {
    return NewItemGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *UserItemRequestBuilder) GetMemberGroups()(*ItemGetMemberGroupsRequestBuilder) {
    return NewItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *UserItemRequestBuilder) GetMemberObjects()(*ItemGetMemberObjectsRequestBuilder) {
    return NewItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InferenceClassification provides operations to manage the inferenceClassification property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) InferenceClassification()(*ItemInferenceClassificationRequestBuilder) {
    return NewItemInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Insights provides operations to manage the insights property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Insights()(*ItemInsightsRequestBuilder) {
    return NewItemInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// JoinedTeams provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) JoinedTeams()(*ItemJoinedTeamsRequestBuilder) {
    return NewItemJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) LicenseDetails()(*ItemLicenseDetailsRequestBuilder) {
    return NewItemLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MailFolders provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) MailFolders()(*ItemMailFoldersRequestBuilder) {
    return NewItemMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ManagedAppRegistrations()(*ItemManagedAppRegistrationsRequestBuilder) {
    return NewItemManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ManagedDevices()(*ItemManagedDevicesRequestBuilder) {
    return NewItemManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Manager provides operations to manage the manager property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Manager()(*ItemManagerRequestBuilder) {
    return NewItemManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) MemberOf()(*ItemMemberOfRequestBuilder) {
    return NewItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Messages()(*ItemMessagesRequestBuilder) {
    return NewItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Oauth2PermissionGrants()(*ItemOauth2PermissionGrantsRequestBuilder) {
    return NewItemOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Onenote()(*ItemOnenoteRequestBuilder) {
    return NewItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OnlineMeetings provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OnlineMeetings()(*ItemOnlineMeetingsRequestBuilder) {
    return NewItemOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Outlook provides operations to manage the outlook property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Outlook()(*ItemOutlookRequestBuilder) {
    return NewItemOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedDevices provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OwnedDevices()(*ItemOwnedDevicesRequestBuilder) {
    return NewItemOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OwnedObjects()(*ItemOwnedObjectsRequestBuilder) {
    return NewItemOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-update?view=graph-rest-1.0
func (m *UserItemRequestBuilder) Patch(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable, requestConfiguration *UserItemRequestBuilderPatchRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable), nil
}
// People provides operations to manage the people property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) People()(*ItemPeopleRequestBuilder) {
    return NewItemPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Photo()(*ItemPhotoRequestBuilder) {
    return NewItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Photos provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Photos()(*ItemPhotosRequestBuilder) {
    return NewItemPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Planner provides operations to manage the planner property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Planner()(*ItemPlannerRequestBuilder) {
    return NewItemPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Presence provides operations to manage the presence property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Presence()(*ItemPresenceRequestBuilder) {
    return NewItemPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RegisteredDevices provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) RegisteredDevices()(*ItemRegisteredDevicesRequestBuilder) {
    return NewItemRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserItemRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime)
}
// RemoveAllDevicesFromManagement provides operations to call the removeAllDevicesFromManagement method.
func (m *UserItemRequestBuilder) RemoveAllDevicesFromManagement()(*ItemRemoveAllDevicesFromManagementRequestBuilder) {
    return NewItemRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReprocessLicenseAssignment provides operations to call the reprocessLicenseAssignment method.
func (m *UserItemRequestBuilder) ReprocessLicenseAssignment()(*ItemReprocessLicenseAssignmentRequestBuilder) {
    return NewItemReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Restore provides operations to call the restore method.
func (m *UserItemRequestBuilder) Restore()(*ItemRestoreRequestBuilder) {
    return NewItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RevokeSignInSessions provides operations to call the revokeSignInSessions method.
func (m *UserItemRequestBuilder) RevokeSignInSessions()(*ItemRevokeSignInSessionsRequestBuilder) {
    return NewItemRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ScopedRoleMemberOf provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ScopedRoleMemberOf()(*ItemScopedRoleMemberOfRequestBuilder) {
    return NewItemScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SendMail provides operations to call the sendMail method.
func (m *UserItemRequestBuilder) SendMail()(*ItemSendMailRequestBuilder) {
    return NewItemSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Teamwork provides operations to manage the teamwork property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Teamwork()(*ItemTeamworkRequestBuilder) {
    return NewItemTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UserItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Todo provides operations to manage the todo property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Todo()(*ItemTodoRequestBuilder) {
    return NewItemTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation retrieve the properties and relationships of user object.
func (m *UserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UserItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Userable, requestConfiguration *UserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) TransitiveMemberOf()(*ItemTransitiveMemberOfRequestBuilder) {
    return NewItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TranslateExchangeIds provides operations to call the translateExchangeIds method.
func (m *UserItemRequestBuilder) TranslateExchangeIds()(*ItemTranslateExchangeIdsRequestBuilder) {
    return NewItemTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WipeManagedAppRegistrationsByDeviceTag provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
func (m *UserItemRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return NewItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
