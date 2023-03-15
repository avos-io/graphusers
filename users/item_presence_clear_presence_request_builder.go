package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemPresenceClearPresenceRequestBuilder provides operations to call the clearPresence method.
type ItemPresenceClearPresenceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemPresenceClearPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPresenceClearPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPresenceClearPresenceRequestBuilderInternal instantiates a new ClearPresenceRequestBuilder and sets the default values.
func NewItemPresenceClearPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceClearPresenceRequestBuilder) {
    m := &ItemPresenceClearPresenceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/presence/clearPresence";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemPresenceClearPresenceRequestBuilder instantiates a new ClearPresenceRequestBuilder and sets the default values.
func NewItemPresenceClearPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceClearPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPresenceClearPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the application's presence session for a user. If it is the user's only presence session, the user's presence will change to `Offline/Offline`. For details about presences sessions, see presence: setPresence.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/presence-clearpresence?view=graph-rest-1.0
func (m *ItemPresenceClearPresenceRequestBuilder) Post(ctx context.Context, body ItemPresenceClearPresencePostRequestBodyable, requestConfiguration *ItemPresenceClearPresenceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation clear the application's presence session for a user. If it is the user's only presence session, the user's presence will change to `Offline/Offline`. For details about presences sessions, see presence: setPresence.
func (m *ItemPresenceClearPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPresenceClearPresencePostRequestBodyable, requestConfiguration *ItemPresenceClearPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
