package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersRequestBuilder builds and executes requests for operations under \groups\{group-id}\members
type ItemMembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersRequestBuilder) {
    m := &ItemMembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/members";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemMembersRequestBuilder) Count()(*ItemMembersCountRequestBuilder) {
    return NewItemMembersCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GraphApplication casts the previous resource to application.
func (m *ItemMembersRequestBuilder) GraphApplication()(*ItemMembersGraphApplicationRequestBuilder) {
    return NewItemMembersGraphApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GraphDevice casts the previous resource to device.
func (m *ItemMembersRequestBuilder) GraphDevice()(*ItemMembersGraphDeviceRequestBuilder) {
    return NewItemMembersGraphDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GraphGroup casts the previous resource to group.
func (m *ItemMembersRequestBuilder) GraphGroup()(*ItemMembersGraphGroupRequestBuilder) {
    return NewItemMembersGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
func (m *ItemMembersRequestBuilder) GraphOrgContact()(*ItemMembersGraphOrgContactRequestBuilder) {
    return NewItemMembersGraphOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemMembersRequestBuilder) GraphServicePrincipal()(*ItemMembersGraphServicePrincipalRequestBuilder) {
    return NewItemMembersGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GraphUser casts the previous resource to user.
func (m *ItemMembersRequestBuilder) GraphUser()(*ItemMembersGraphUserRequestBuilder) {
    return NewItemMembersGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ref provides operations to manage the collection of group entities.
func (m *ItemMembersRequestBuilder) Ref()(*ItemMembersRefRequestBuilder) {
    return NewItemMembersRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
