package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\attachments\$count
type ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
}
// ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderGetQueryParameters
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/attachments/$count{?%24filter}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}