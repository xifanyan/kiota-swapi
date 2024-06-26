package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemChangesQueueRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\changes\queue
type ItemCollectionsItemChangesQueueRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemChangesQueueRequestBuilderGetQueryParameters change requests submitted with 'blockUntilComplete=false' are queued. Their value is immediately visible when fetching it, but searches are still based on their old state until the queue has been processed. This method allows to wait for queued change requests; more precisely: for all change requests which are queued (or in progress) at the time when this wait request is submitted.
type ItemCollectionsItemChangesQueueRequestBuilderGetQueryParameters struct {
    // Defines that only scheduled high priority change requests should be considered. 'High Priority' means: change requests involving single documents.
    OnlyHighPriorityChanges *bool `uriparametername:"onlyHighPriorityChanges"`
    // Defines an optional timeout in milliseconds. Default is 60000. Use a negative value to wait indefinitely.
    TimeoutMillis *int64 `uriparametername:"timeoutMillis"`
}
// ItemCollectionsItemChangesQueueRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemChangesQueueRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemChangesQueueRequestBuilderGetQueryParameters
}
// ItemCollectionsItemChangesQueueRequestBuilderPostQueryParameters change requests submitted with 'blockUntilComplete=false' are queued. Their value is immediately visible when fetching it, but searches are still based on their old state until the queue has been processed. This method allows to wait for queued change requests; more precisely: for all change requests which are queued (or in progress) at the time when this wait request is submitted.
type ItemCollectionsItemChangesQueueRequestBuilderPostQueryParameters struct {
    // Defines that only scheduled high priority change requests should be considered. 'High Priority' means: change requests involving single documents.
    OnlyHighPriorityChanges *bool `uriparametername:"onlyHighPriorityChanges"`
    // Defines an optional timeout in milliseconds. Default is 60000. Use a negative value to wait indefinitely.
    TimeoutMillis *int64 `uriparametername:"timeoutMillis"`
}
// ItemCollectionsItemChangesQueueRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemChangesQueueRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemChangesQueueRequestBuilderPostQueryParameters
}
// NewItemCollectionsItemChangesQueueRequestBuilderInternal instantiates a new ItemCollectionsItemChangesQueueRequestBuilder and sets the default values.
func NewItemCollectionsItemChangesQueueRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemChangesQueueRequestBuilder) {
    m := &ItemCollectionsItemChangesQueueRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/changes/queue{?onlyHighPriorityChanges*,timeoutMillis*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemChangesQueueRequestBuilder instantiates a new ItemCollectionsItemChangesQueueRequestBuilder and sets the default values.
func NewItemCollectionsItemChangesQueueRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemChangesQueueRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemChangesQueueRequestBuilderInternal(urlParams, requestAdapter)
}
// Get change requests submitted with 'blockUntilComplete=false' are queued. Their value is immediately visible when fetching it, but searches are still based on their old state until the queue has been processed. This method allows to wait for queued change requests; more precisely: for all change requests which are queued (or in progress) at the time when this wait request is submitted.
// returns a WaitForPendingChangesResultable when successful
func (m *ItemCollectionsItemChangesQueueRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemChangesQueueRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.WaitForPendingChangesResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateWaitForPendingChangesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.WaitForPendingChangesResultable), nil
}
// Post change requests submitted with 'blockUntilComplete=false' are queued. Their value is immediately visible when fetching it, but searches are still based on their old state until the queue has been processed. This method allows to wait for queued change requests; more precisely: for all change requests which are queued (or in progress) at the time when this wait request is submitted.
// returns a WaitForPendingChangesResultable when successful
func (m *ItemCollectionsItemChangesQueueRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemChangesQueueRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.WaitForPendingChangesResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateWaitForPendingChangesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.WaitForPendingChangesResultable), nil
}
// ToGetRequestInformation change requests submitted with 'blockUntilComplete=false' are queued. Their value is immediately visible when fetching it, but searches are still based on their old state until the queue has been processed. This method allows to wait for queued change requests; more precisely: for all change requests which are queued (or in progress) at the time when this wait request is submitted.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemChangesQueueRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemChangesQueueRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation change requests submitted with 'blockUntilComplete=false' are queued. Their value is immediately visible when fetching it, but searches are still based on their old state until the queue has been processed. This method allows to wait for queued change requests; more precisely: for all change requests which are queued (or in progress) at the time when this wait request is submitted.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemChangesQueueRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemChangesQueueRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemChangesQueueRequestBuilder when successful
func (m *ItemCollectionsItemChangesQueueRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemChangesQueueRequestBuilder) {
    return NewItemCollectionsItemChangesQueueRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
