package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GroupItemRequestBuilder builds and executes requests for operations under \groups\{group-id}
type GroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AcceptedSenders provides operations to manage the acceptedSenders property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) AcceptedSenders()(*ItemAcceptedSendersRequestBuilder) {
    return NewItemAcceptedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AddFavorite provides operations to call the addFavorite method.
func (m *GroupItemRequestBuilder) AddFavorite()(*ItemAddFavoriteRequestBuilder) {
    return NewItemAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) AppRoleAssignments()(*ItemAppRoleAssignmentsRequestBuilder) {
    return NewItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignLicense provides operations to call the assignLicense method.
func (m *GroupItemRequestBuilder) AssignLicense()(*ItemAssignLicenseRequestBuilder) {
    return NewItemAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Calendar()(*ItemCalendarRequestBuilder) {
    return NewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) CalendarView()(*ItemCalendarViewRequestBuilder) {
    return NewItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CheckGrantedPermissionsForApp provides operations to call the checkGrantedPermissionsForApp method.
func (m *GroupItemRequestBuilder) CheckGrantedPermissionsForApp()(*ItemCheckGrantedPermissionsForAppRequestBuilder) {
    return NewItemCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *GroupItemRequestBuilder) CheckMemberGroups()(*ItemCheckMemberGroupsRequestBuilder) {
    return NewItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *GroupItemRequestBuilder) CheckMemberObjects()(*ItemCheckMemberObjectsRequestBuilder) {
    return NewItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    m := &GroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Conversations provides operations to manage the conversations property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Conversations()(*ItemConversationsRequestBuilder) {
    return NewItemConversationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CreatedOnBehalfOf provides operations to manage the createdOnBehalfOf property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) CreatedOnBehalfOf()(*ItemCreatedOnBehalfOfRequestBuilder) {
    return NewItemCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Drive()(*ItemDriveRequestBuilder) {
    return NewItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Drives()(*ItemDrivesRequestBuilder) {
    return NewItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Events provides operations to manage the events property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Extensions()(*ItemExtensionsRequestBuilder) {
    return NewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *GroupItemRequestBuilder) GetMemberGroups()(*ItemGetMemberGroupsRequestBuilder) {
    return NewItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *GroupItemRequestBuilder) GetMemberObjects()(*ItemGetMemberObjectsRequestBuilder) {
    return NewItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupLifecyclePolicies provides operations to manage the groupLifecyclePolicies property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) GroupLifecyclePolicies()(*ItemGroupLifecyclePoliciesRequestBuilder) {
    return NewItemGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) MemberOf()(*ItemMemberOfRequestBuilder) {
    return NewItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersWithLicenseErrors provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) MembersWithLicenseErrors()(*ItemMembersWithLicenseErrorsRequestBuilder) {
    return NewItemMembersWithLicenseErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Onenote()(*ItemOnenoteRequestBuilder) {
    return NewItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Owners()(*ItemOwnersRequestBuilder) {
    return NewItemOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) PermissionGrants()(*ItemPermissionGrantsRequestBuilder) {
    return NewItemPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Photo()(*ItemPhotoRequestBuilder) {
    return NewItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Photos provides operations to manage the photos property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Photos()(*ItemPhotosRequestBuilder) {
    return NewItemPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Planner provides operations to manage the planner property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Planner()(*ItemPlannerRequestBuilder) {
    return NewItemPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RejectedSenders provides operations to manage the rejectedSenders property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) RejectedSenders()(*ItemRejectedSendersRequestBuilder) {
    return NewItemRejectedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RemoveFavorite provides operations to call the removeFavorite method.
func (m *GroupItemRequestBuilder) RemoveFavorite()(*ItemRemoveFavoriteRequestBuilder) {
    return NewItemRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Renew provides operations to call the renew method.
func (m *GroupItemRequestBuilder) Renew()(*ItemRenewRequestBuilder) {
    return NewItemRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResetUnseenCount provides operations to call the resetUnseenCount method.
func (m *GroupItemRequestBuilder) ResetUnseenCount()(*ItemResetUnseenCountRequestBuilder) {
    return NewItemResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Restore provides operations to call the restore method.
func (m *GroupItemRequestBuilder) Restore()(*ItemRestoreRequestBuilder) {
    return NewItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Sites()(*ItemSitesRequestBuilder) {
    return NewItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubscribeByMail provides operations to call the subscribeByMail method.
func (m *GroupItemRequestBuilder) SubscribeByMail()(*ItemSubscribeByMailRequestBuilder) {
    return NewItemSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Team provides operations to manage the team property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Team()(*ItemTeamRequestBuilder) {
    return NewItemTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Threads provides operations to manage the threads property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Threads()(*ItemThreadsRequestBuilder) {
    return NewItemThreadsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) TransitiveMemberOf()(*ItemTransitiveMemberOfRequestBuilder) {
    return NewItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TransitiveMembers provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) TransitiveMembers()(*ItemTransitiveMembersRequestBuilder) {
    return NewItemTransitiveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UnsubscribeByMail provides operations to call the unsubscribeByMail method.
func (m *GroupItemRequestBuilder) UnsubscribeByMail()(*ItemUnsubscribeByMailRequestBuilder) {
    return NewItemUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *GroupItemRequestBuilder) ValidateProperties()(*ItemValidatePropertiesRequestBuilder) {
    return NewItemValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
