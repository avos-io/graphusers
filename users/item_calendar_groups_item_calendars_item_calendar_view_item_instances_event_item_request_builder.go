package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i42049ece93d9a63e1eca3944b9014beb5e382587d98b67237a4e6ba308528453 "github.com/avos-io/graphusers/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Cancel provides operations to call the cancel method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemDismissReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Forward provides operations to call the forward method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Eventable, error) {
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
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
