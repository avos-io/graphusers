package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder provides operations to count the resources in the collection.
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderGetQueryParameters get the number of the resource
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderGetQueryParameters
}
// NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder) {
    m := &ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/$count{?%24search,%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
