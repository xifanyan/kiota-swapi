package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCollectionsItemBinaryItemContentRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\binary\{recordId}\content
type ItemCollectionsItemBinaryItemContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemBinaryItemContentRequestBuilderGetQueryParameters returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
type ItemCollectionsItemBinaryItemContentRequestBuilderGetQueryParameters struct {
    // A single field name storing binary data.
    Field *string `uriparametername:"field"`
}
// ItemCollectionsItemBinaryItemContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemBinaryItemContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemBinaryItemContentRequestBuilderGetQueryParameters
}
// ItemCollectionsItemBinaryItemContentRequestBuilderPostQueryParameters returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
type ItemCollectionsItemBinaryItemContentRequestBuilderPostQueryParameters struct {
    // A single field name storing binary data.
    Field *string `uriparametername:"field"`
}
// ItemCollectionsItemBinaryItemContentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemBinaryItemContentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemBinaryItemContentRequestBuilderPostQueryParameters
}
// NewItemCollectionsItemBinaryItemContentRequestBuilderInternal instantiates a new ItemCollectionsItemBinaryItemContentRequestBuilder and sets the default values.
func NewItemCollectionsItemBinaryItemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemBinaryItemContentRequestBuilder) {
    m := &ItemCollectionsItemBinaryItemContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/binary/{recordId}/content{?field*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemBinaryItemContentRequestBuilder instantiates a new ItemCollectionsItemBinaryItemContentRequestBuilder and sets the default values.
func NewItemCollectionsItemBinaryItemContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemBinaryItemContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemBinaryItemContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
// returns a []byte when successful
func (m *ItemCollectionsItemBinaryItemContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemBinaryItemContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Post returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
// returns a []byte when successful
func (m *ItemCollectionsItemBinaryItemContentRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemBinaryItemContentRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemBinaryItemContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemBinaryItemContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream")
    return requestInfo, nil
}
// ToPostRequestInformation returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemBinaryItemContentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemBinaryItemContentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemBinaryItemContentRequestBuilder when successful
func (m *ItemCollectionsItemBinaryItemContentRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemBinaryItemContentRequestBuilder) {
    return NewItemCollectionsItemBinaryItemContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
