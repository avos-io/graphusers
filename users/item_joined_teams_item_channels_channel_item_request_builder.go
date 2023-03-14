package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemJoinedTeamsItemChannelsChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ItemJoinedTeamsItemChannelsChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemJoinedTeamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedTeamsItemChannelsChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type ItemJoinedTeamsItemChannelsChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedTeamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedTeamsItemChannelsChannelItemRequestBuilderGetQueryParameters
}
// ItemJoinedTeamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) CompleteMigration()(*ItemJoinedTeamsItemChannelsItemCompleteMigrationRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemJoinedTeamsItemChannelsChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemChannelsChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) {
    m := &ItemJoinedTeamsItemChannelsChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/channels/{channel%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemJoinedTeamsItemChannelsChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemChannelsChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemChannelsChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property channels for users
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*ItemJoinedTeamsItemChannelsItemDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) FilesFolder()(*ItemJoinedTeamsItemChannelsItemFilesFolderRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the collection of channels and messages associated with the team.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Channelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Channelable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) Members()(*ItemJoinedTeamsItemChannelsItemMembersRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) MembersById(id string)(*ItemJoinedTeamsItemChannelsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemJoinedTeamsItemChannelsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) Messages()(*ItemJoinedTeamsItemChannelsItemMessagesRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) MessagesById(id string)(*ItemJoinedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewItemJoinedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property channels in users
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) Patch(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Channelable, requestConfiguration *ItemJoinedTeamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Channelable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Channelable), nil
}
// ProvisionEmail provides operations to call the provisionEmail method.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) ProvisionEmail()(*ItemJoinedTeamsItemChannelsItemProvisionEmailRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RemoveEmail provides operations to call the removeEmail method.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) RemoveEmail()(*ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) SharedWithTeams()(*ItemJoinedTeamsItemChannelsItemSharedWithTeamsRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SharedWithTeamsById provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) SharedWithTeamsById(id string)(*ItemJoinedTeamsItemChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return NewItemJoinedTeamsItemChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) Tabs()(*ItemJoinedTeamsItemChannelsItemTabsRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) TabsById(id string)(*ItemJoinedTeamsItemChannelsItemTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewItemJoinedTeamsItemChannelsItemTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property channels for users
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the collection of channels and messages associated with the team.
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property channels in users
func (m *ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Channelable, requestConfiguration *ItemJoinedTeamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
