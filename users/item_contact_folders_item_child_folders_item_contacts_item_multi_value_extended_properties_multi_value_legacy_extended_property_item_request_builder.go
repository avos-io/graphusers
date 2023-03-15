package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.contact entity.
type ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters the collection of multi-value extended properties defined for the contact. Read-only. Nullable.
type ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters
}
// ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal instantiates a new MultiValueLegacyExtendedPropertyItemRequestBuilder and sets the default values.
func NewItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    m := &ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/childFolders/{contactFolder%2Did1}/contacts/{contact%2Did}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder instantiates a new MultiValueLegacyExtendedPropertyItemRequestBuilder and sets the default values.
func NewItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property multiValueExtendedProperties for users
func (m *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MultiValueLegacyExtendedPropertyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MultiValueLegacyExtendedPropertyable), nil
}
// Patch update the navigation property multiValueExtendedProperties in users
func (m *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) Patch(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MultiValueLegacyExtendedPropertyable, requestConfiguration *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MultiValueLegacyExtendedPropertyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MultiValueLegacyExtendedPropertyable), nil
}
// ToDeleteRequestInformation delete navigation property multiValueExtendedProperties for users
func (m *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property multiValueExtendedProperties in users
func (m *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.MultiValueLegacyExtendedPropertyable, requestConfiguration *ItemContactFoldersItemChildFoldersItemContactsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
