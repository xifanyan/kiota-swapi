package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections
type ItemCollectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCollectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCollectionId gets an item from the github.com/xifanyan/kiota-swapi.projects.item.collections.item collection
// returns a *ItemCollectionsWithCollectionItemRequestBuilder when successful
func (m *ItemCollectionsRequestBuilder) ByCollectionId(collectionId string)(*ItemCollectionsWithCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if collectionId != "" {
        urlTplParams["collectionId"] = collectionId
    }
    return NewItemCollectionsWithCollectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCollectionsRequestBuilderInternal instantiates a new ItemCollectionsRequestBuilder and sets the default values.
func NewItemCollectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsRequestBuilder) {
    m := &ItemCollectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections", pathParameters),
    }
    return m
}
// NewItemCollectionsRequestBuilder instantiates a new ItemCollectionsRequestBuilder and sets the default values.
func NewItemCollectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get decisiv Search offers multiple domains (e.g. Documents, Matter, People, etc) belonging to a project. With this operation, the set of available domains is returned, e.g. for display as tabs in the UI. For any project that is not a non-merging meta-engine, the collection 'default' will be returned as the only collection.
// returns a CollectionsResultable when successful
func (m *ItemCollectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CollectionsResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateCollectionsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CollectionsResultable), nil
}
// Post decisiv Search offers multiple domains (e.g. Documents, Matter, People, etc) belonging to a project. With this operation, the set of available domains is returned, e.g. for display as tabs in the UI. For any project that is not a non-merging meta-engine, the collection 'default' will be returned as the only collection.
// returns a CollectionsResultable when successful
func (m *ItemCollectionsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CollectionsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateCollectionsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CollectionsResultable), nil
}
// ToGetRequestInformation decisiv Search offers multiple domains (e.g. Documents, Matter, People, etc) belonging to a project. With this operation, the set of available domains is returned, e.g. for display as tabs in the UI. For any project that is not a non-merging meta-engine, the collection 'default' will be returned as the only collection.
// returns a *RequestInformation when successful
func (m *ItemCollectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation decisiv Search offers multiple domains (e.g. Documents, Matter, People, etc) belonging to a project. With this operation, the set of available domains is returned, e.g. for display as tabs in the UI. For any project that is not a non-merging meta-engine, the collection 'default' will be returned as the only collection.
// returns a *RequestInformation when successful
func (m *ItemCollectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsRequestBuilder when successful
func (m *ItemCollectionsRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsRequestBuilder) {
    return NewItemCollectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
