package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemRecordsWithRecordItemRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\records\{recordId}
type ItemCollectionsItemRecordsWithRecordItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemRecordsWithRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsWithRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCollectionsItemRecordsWithRecordItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsWithRecordItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCollectionsItemRecordsWithRecordItemRequestBuilderInternal instantiates a new ItemCollectionsItemRecordsWithRecordItemRequestBuilder and sets the default values.
func NewItemCollectionsItemRecordsWithRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemRecordsWithRecordItemRequestBuilder) {
    m := &ItemCollectionsItemRecordsWithRecordItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/records/{recordId}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemRecordsWithRecordItemRequestBuilder instantiates a new ItemCollectionsItemRecordsWithRecordItemRequestBuilder and sets the default values.
func NewItemCollectionsItemRecordsWithRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemRecordsWithRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemRecordsWithRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
// returns a *ItemCollectionsItemRecordsItemContentRequestBuilder when successful
func (m *ItemCollectionsItemRecordsWithRecordItemRequestBuilder) Content()(*ItemCollectionsItemRecordsItemContentRequestBuilder) {
    return NewItemCollectionsItemRecordsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get discover available resources for a record ID (document ID).
// returns a RecordResourcesResultable when successful
func (m *ItemCollectionsItemRecordsWithRecordItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsWithRecordItemRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.RecordResourcesResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateRecordResourcesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.RecordResourcesResultable), nil
}
// Post discover available resources for a record ID (document ID).
// returns a RecordResourcesResultable when successful
func (m *ItemCollectionsItemRecordsWithRecordItemRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsWithRecordItemRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.RecordResourcesResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateRecordResourcesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.RecordResourcesResultable), nil
}
// ToGetRequestInformation discover available resources for a record ID (document ID).
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsWithRecordItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsWithRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation discover available resources for a record ID (document ID).
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsWithRecordItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsWithRecordItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemRecordsWithRecordItemRequestBuilder when successful
func (m *ItemCollectionsItemRecordsWithRecordItemRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemRecordsWithRecordItemRequestBuilder) {
    return NewItemCollectionsItemRecordsWithRecordItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
