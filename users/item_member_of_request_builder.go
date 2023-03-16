package users

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMemberOfRequestBuilder builds and executes requests for operations under \users\{user-id}\memberOf
type ItemMemberOfRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// NewItemMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewItemMemberOfRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *ItemMemberOfRequestBuilder {
	m := &ItemMemberOfRequestBuilder{}
	m.urlTemplate = "{+baseurl}/users/{user%2Did}/memberOf"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemMemberOfRequestBuilder instantiates a new MemberOfRequestBuilder and sets the default values.
func NewItemMemberOfRequestBuilder(
	rawUrl string,
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter,
) *ItemMemberOfRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *ItemMemberOfRequestBuilder) Count() *ItemMemberOfCountRequestBuilder {
	return NewItemMemberOfCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GraphGroup casts the previous resource to group.
func (m *ItemMemberOfRequestBuilder) GraphGroup() *ItemMemberOfGraphGroupRequestBuilder {
	return NewItemMemberOfGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
