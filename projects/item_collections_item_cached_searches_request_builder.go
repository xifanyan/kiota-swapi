package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemCachedSearchesRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\cachedSearches
type ItemCollectionsItemCachedSearchesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemCachedSearchesRequestBuilderDeleteQueryParameters drops cached searches.
type ItemCollectionsItemCachedSearchesRequestBuilderDeleteQueryParameters struct {
    // A comma-separated list of creation trace ids. Only these associated caches will be cleared. Omitting this value will close all available searches for the specified collection. Such values can be retrieved using the GET endpoint.
    CreationTraceIds *string `uriparametername:"creationTraceIds"`
}
// ItemCollectionsItemCachedSearchesRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemCachedSearchesRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemCachedSearchesRequestBuilderDeleteQueryParameters
}
// ItemCollectionsItemCachedSearchesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemCachedSearchesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCollectionsItemCachedSearchesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemCachedSearchesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCollectionsItemCachedSearchesRequestBuilderInternal instantiates a new ItemCollectionsItemCachedSearchesRequestBuilder and sets the default values.
func NewItemCollectionsItemCachedSearchesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemCachedSearchesRequestBuilder) {
    m := &ItemCollectionsItemCachedSearchesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/cachedSearches", pathParameters),
    }
    return m
}
// NewItemCollectionsItemCachedSearchesRequestBuilder instantiates a new ItemCollectionsItemCachedSearchesRequestBuilder and sets the default values.
func NewItemCollectionsItemCachedSearchesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemCachedSearchesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemCachedSearchesRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete drops cached searches.
// returns a []CachedSearchDescriptionable when successful
func (m *ItemCollectionsItemCachedSearchesRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCollectionsItemCachedSearchesRequestBuilderDeleteRequestConfiguration)([]ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateCachedSearchDescriptionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable)
        }
    }
    return val, nil
}
// Get returns a list of cached searches. Search web api automatically caches search results such that paging or following actions are fast without re-executing the same query. This method allows to inspect the available caches for the current session and for the selected project / collection. The result value resembles a trace id for each cached search; more precisely: the trace id which created the cached search result. Trace ids can be set using the request header SWA-MDC-TOKEN; they typically serve as log file analysis utility - in case of searches, they also identify the search initiator context.
// returns a []CachedSearchDescriptionable when successful
func (m *ItemCollectionsItemCachedSearchesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemCachedSearchesRequestBuilderGetRequestConfiguration)([]ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateCachedSearchDescriptionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable)
        }
    }
    return val, nil
}
// Post returns a list of cached searches. Search web api automatically caches search results such that paging or following actions are fast without re-executing the same query. This method allows to inspect the available caches for the current session and for the selected project / collection. The result value resembles a trace id for each cached search; more precisely: the trace id which created the cached search result. Trace ids can be set using the request header SWA-MDC-TOKEN; they typically serve as log file analysis utility - in case of searches, they also identify the search initiator context.
// returns a []CachedSearchDescriptionable when successful
func (m *ItemCollectionsItemCachedSearchesRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemCachedSearchesRequestBuilderPostRequestConfiguration)([]ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateCachedSearchDescriptionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CachedSearchDescriptionable)
        }
    }
    return val, nil
}
// ToDeleteRequestInformation drops cached searches.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemCachedSearchesRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemCachedSearchesRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/projects/{projectId}/collections/{collectionId}/cachedSearches{?creationTraceIds*}", m.BaseRequestBuilder.PathParameters)
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
// ToGetRequestInformation returns a list of cached searches. Search web api automatically caches search results such that paging or following actions are fast without re-executing the same query. This method allows to inspect the available caches for the current session and for the selected project / collection. The result value resembles a trace id for each cached search; more precisely: the trace id which created the cached search result. Trace ids can be set using the request header SWA-MDC-TOKEN; they typically serve as log file analysis utility - in case of searches, they also identify the search initiator context.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemCachedSearchesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemCachedSearchesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation returns a list of cached searches. Search web api automatically caches search results such that paging or following actions are fast without re-executing the same query. This method allows to inspect the available caches for the current session and for the selected project / collection. The result value resembles a trace id for each cached search; more precisely: the trace id which created the cached search result. Trace ids can be set using the request header SWA-MDC-TOKEN; they typically serve as log file analysis utility - in case of searches, they also identify the search initiator context.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemCachedSearchesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemCachedSearchesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemCachedSearchesRequestBuilder when successful
func (m *ItemCollectionsItemCachedSearchesRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemCachedSearchesRequestBuilder) {
    return NewItemCollectionsItemCachedSearchesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
