package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemTeamworkInstalledAppsItemChatRequestBuilder provides operations to manage the chat property of the microsoft.graph.userScopeTeamsAppInstallation entity.
type ItemTeamworkInstalledAppsItemChatRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamworkInstalledAppsItemChatRequestBuilderGetQueryParameters retrieve the chat of the specified user and Teams app.
type ItemTeamworkInstalledAppsItemChatRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamworkInstalledAppsItemChatRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamworkInstalledAppsItemChatRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamworkInstalledAppsItemChatRequestBuilderGetQueryParameters
}
// NewItemTeamworkInstalledAppsItemChatRequestBuilderInternal instantiates a new ChatRequestBuilder and sets the default values.
func NewItemTeamworkInstalledAppsItemChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamworkInstalledAppsItemChatRequestBuilder) {
    m := &ItemTeamworkInstalledAppsItemChatRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/teamwork/installedApps/{userScopeTeamsAppInstallation%2Did}/chat{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamworkInstalledAppsItemChatRequestBuilder instantiates a new ChatRequestBuilder and sets the default values.
func NewItemTeamworkInstalledAppsItemChatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamworkInstalledAppsItemChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamworkInstalledAppsItemChatRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve the chat of the specified user and Teams app.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/userscopeteamsappinstallation-get-chat?view=graph-rest-1.0
func (m *ItemTeamworkInstalledAppsItemChatRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamworkInstalledAppsItemChatRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Chatable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateChatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Chatable), nil
}
// ToGetRequestInformation retrieve the chat of the specified user and Teams app.
func (m *ItemTeamworkInstalledAppsItemChatRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamworkInstalledAppsItemChatRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
