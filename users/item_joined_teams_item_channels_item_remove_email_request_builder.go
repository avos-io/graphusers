package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder provides operations to call the removeEmail method.
type ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilderInternal instantiates a new RemoveEmailRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder) {
    m := &ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/channels/{channel%2Did}/removeEmail";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder instantiates a new RemoveEmailRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove the email address of a channel. You can remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/channel-removeemail?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation remove the email address of a channel. You can remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
func (m *ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemChannelsItemRemoveEmailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
