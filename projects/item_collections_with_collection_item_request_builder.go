package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsWithCollectionItemRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}
type ItemCollectionsWithCollectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsWithCollectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsWithCollectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCollectionsWithCollectionItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsWithCollectionItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Binary the binary property
// returns a *ItemCollectionsItemBinaryRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Binary()(*ItemCollectionsItemBinaryRequestBuilder) {
    return NewItemCollectionsItemBinaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CachedSearches the cachedSearches property
// returns a *ItemCollectionsItemCachedSearchesRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) CachedSearches()(*ItemCollectionsItemCachedSearchesRequestBuilder) {
    return NewItemCollectionsItemCachedSearchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Changes the changes property
// returns a *ItemCollectionsItemChangesRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Changes()(*ItemCollectionsItemChangesRequestBuilder) {
    return NewItemCollectionsItemChangesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCollectionsWithCollectionItemRequestBuilderInternal instantiates a new ItemCollectionsWithCollectionItemRequestBuilder and sets the default values.
func NewItemCollectionsWithCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsWithCollectionItemRequestBuilder) {
    m := &ItemCollectionsWithCollectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}", pathParameters),
    }
    return m
}
// NewItemCollectionsWithCollectionItemRequestBuilder instantiates a new ItemCollectionsWithCollectionItemRequestBuilder and sets the default values.
func NewItemCollectionsWithCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsWithCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsWithCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Fields the fields property
// returns a *ItemCollectionsItemFieldsRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Fields()(*ItemCollectionsItemFieldsRequestBuilder) {
    return NewItemCollectionsItemFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Filters the filters property
// returns a *ItemCollectionsItemFiltersRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Filters()(*ItemCollectionsItemFiltersRequestBuilder) {
    return NewItemCollectionsItemFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get discover available collection-level resources. Examples are 'records' which describes the actual searchable objects in the domain, 'filters' which describes the folder fields (taxonomies, smart filters), and 'fields' which describs the available fields.
// returns a CollectionResourcesResultable when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsWithCollectionItemRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CollectionResourcesResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateCollectionResourcesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CollectionResourcesResultable), nil
}
// Measures the measures property
// returns a *ItemCollectionsItemMeasuresRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Measures()(*ItemCollectionsItemMeasuresRequestBuilder) {
    return NewItemCollectionsItemMeasuresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post discover available collection-level resources. Examples are 'records' which describes the actual searchable objects in the domain, 'filters' which describes the folder fields (taxonomies, smart filters), and 'fields' which describs the available fields.
// returns a CollectionResourcesResultable when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsWithCollectionItemRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CollectionResourcesResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateCollectionResourcesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CollectionResourcesResultable), nil
}
// Records the records property
// returns a *ItemCollectionsItemRecordsRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) Records()(*ItemCollectionsItemRecordsRequestBuilder) {
    return NewItemCollectionsItemRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchToken the searchToken property
// returns a *ItemCollectionsItemSearchTokenRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) SearchToken()(*ItemCollectionsItemSearchTokenRequestBuilder) {
    return NewItemCollectionsItemSearchTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation discover available collection-level resources. Examples are 'records' which describes the actual searchable objects in the domain, 'filters' which describes the folder fields (taxonomies, smart filters), and 'fields' which describs the available fields.
// returns a *RequestInformation when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsWithCollectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation discover available collection-level resources. Examples are 'records' which describes the actual searchable objects in the domain, 'filters' which describes the folder fields (taxonomies, smart filters), and 'fields' which describs the available fields.
// returns a *RequestInformation when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsWithCollectionItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsWithCollectionItemRequestBuilder when successful
func (m *ItemCollectionsWithCollectionItemRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsWithCollectionItemRequestBuilder) {
    return NewItemCollectionsWithCollectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
