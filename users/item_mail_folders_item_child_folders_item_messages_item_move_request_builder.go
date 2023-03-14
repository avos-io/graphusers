package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder provides operations to call the move method.
type ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal instantiates a new MoveRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/move";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder instantiates a new MoveRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post move a message to another folder within the specified user's mailbox. This creates a new copy of the message in the destination folder and removes the original message.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/message-move?view=graph-rest-1.0
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) Post(ctx context.Context, body ItemMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderPostRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Messageable), nil
}
// ToPostRequestInformation move a message to another folder within the specified user's mailbox. This creates a new copy of the message in the destination folder and removes the original message.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
