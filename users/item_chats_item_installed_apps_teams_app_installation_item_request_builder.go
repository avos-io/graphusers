package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder provides operations to manage the installedApps property of the microsoft.graph.chat entity.
type ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderGetQueryParameters a collection of all the apps in the chat. Nullable.
type ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderGetQueryParameters
}
// ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal instantiates a new TeamsAppInstallationItemRequestBuilder and sets the default values.
func NewItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    m := &ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/installedApps/{teamsAppInstallation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder instantiates a new TeamsAppInstallationItemRequestBuilder and sets the default values.
func NewItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property installedApps for users
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of all the apps in the chat. Nullable.
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamsAppInstallationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateTeamsAppInstallationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamsAppInstallationable), nil
}
// Patch update the navigation property installedApps in users
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) Patch(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamsAppInstallationable, requestConfiguration *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamsAppInstallationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateTeamsAppInstallationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamsAppInstallationable), nil
}
// TeamsApp provides operations to manage the teamsApp property of the microsoft.graph.teamsAppInstallation entity.
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) TeamsApp()(*ItemChatsItemInstalledAppsItemTeamsAppRequestBuilder) {
    return NewItemChatsItemInstalledAppsItemTeamsAppRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TeamsAppDefinition provides operations to manage the teamsAppDefinition property of the microsoft.graph.teamsAppInstallation entity.
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) TeamsAppDefinition()(*ItemChatsItemInstalledAppsItemTeamsAppDefinitionRequestBuilder) {
    return NewItemChatsItemInstalledAppsItemTeamsAppDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property installedApps for users
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation a collection of all the apps in the chat. Nullable.
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property installedApps in users
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.TeamsAppInstallationable, requestConfiguration *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Upgrade provides operations to call the upgrade method.
func (m *ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) Upgrade()(*ItemChatsItemInstalledAppsItemUpgradeRequestBuilder) {
    return NewItemChatsItemInstalledAppsItemUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
