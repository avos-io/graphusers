package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GroupsRequestBuilder builds and executes requests for operations under \groups
type GroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewGroupsRequestBuilderInternal instantiates a new GroupsRequestBuilder and sets the default values.
func NewGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsRequestBuilder) {
    m := &GroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewGroupsRequestBuilder instantiates a new GroupsRequestBuilder and sets the default values.
func NewGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *GroupsRequestBuilder) Count()(*CountRequestBuilder) {
    return NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delta provides operations to call the delta method.
func (m *GroupsRequestBuilder) Delta()(*DeltaRequestBuilder) {
    return NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetAvailableExtensionProperties provides operations to call the getAvailableExtensionProperties method.
func (m *GroupsRequestBuilder) GetAvailableExtensionProperties()(*GetAvailableExtensionPropertiesRequestBuilder) {
    return NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetByIds provides operations to call the getByIds method.
func (m *GroupsRequestBuilder) GetByIds()(*GetByIdsRequestBuilder) {
    return NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *GroupsRequestBuilder) ValidateProperties()(*ValidatePropertiesRequestBuilder) {
    return NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
