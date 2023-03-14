package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemCalendarsItemEventsEventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
type ItemCalendarsItemEventsEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type ItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters
}
// ItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept provides operations to call the accept method.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Accept()(*ItemCalendarsItemEventsItemAcceptRequestBuilder) {
    return NewItemCalendarsItemEventsItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Attachments()(*ItemCalendarsItemEventsItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemEventsItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) AttachmentsById(id string)(*ItemCalendarsItemEventsItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemCalendarsItemEventsItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Calendar()(*ItemCalendarsItemEventsItemCalendarRequestBuilder) {
    return NewItemCalendarsItemEventsItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Cancel provides operations to call the cancel method.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Cancel()(*ItemCalendarsItemEventsItemCancelRequestBuilder) {
    return NewItemCalendarsItemEventsItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemCalendarsItemEventsEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsEventItemRequestBuilder) {
    m := &ItemCalendarsItemEventsEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarsItemEventsEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemEventsEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Decline()(*ItemCalendarsItemEventsItemDeclineRequestBuilder) {
    return NewItemCalendarsItemEventsItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property events for users
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DismissReminder provides operations to call the dismissReminder method.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) DismissReminder()(*ItemCalendarsItemEventsItemDismissReminderRequestBuilder) {
    return NewItemCalendarsItemEventsItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Extensions()(*ItemCalendarsItemEventsItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemEventsItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) ExtensionsById(id string)(*ItemCalendarsItemEventsItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemCalendarsItemEventsItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Forward provides operations to call the forward method.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Forward()(*ItemCalendarsItemEventsItemForwardRequestBuilder) {
    return NewItemCalendarsItemEventsItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Eventable), nil
}
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Instances()(*ItemCalendarsItemEventsItemInstancesRequestBuilder) {
    return NewItemCalendarsItemEventsItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) InstancesById(id string)(*ItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewItemCalendarsItemEventsItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarsItemEventsItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemEventsItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarsItemEventsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemEventsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property events in users
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) Patch(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Eventable, requestConfiguration *ItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Eventable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
        "5XX": i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Eventable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarsItemEventsItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemEventsItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarsItemEventsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemEventsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarsItemEventsItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemEventsItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarsItemEventsItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemEventsItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property events for users
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property events in users
func (m *ItemCalendarsItemEventsEventItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Eventable, requestConfiguration *ItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
