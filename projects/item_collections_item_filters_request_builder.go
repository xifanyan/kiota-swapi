package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemFiltersRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\filters
type ItemCollectionsItemFiltersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemFiltersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemFiltersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCollectionsItemFiltersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemFiltersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByFieldId gets an item from the github.com/xifanyan/kiota-swapi.projects.item.collections.item.filters.item collection
// returns a *ItemCollectionsItemFiltersWithFieldItemRequestBuilder when successful
func (m *ItemCollectionsItemFiltersRequestBuilder) ByFieldId(fieldId string)(*ItemCollectionsItemFiltersWithFieldItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if fieldId != "" {
        urlTplParams["fieldId"] = fieldId
    }
    return NewItemCollectionsItemFiltersWithFieldItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCollectionsItemFiltersRequestBuilderInternal instantiates a new ItemCollectionsItemFiltersRequestBuilder and sets the default values.
func NewItemCollectionsItemFiltersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemFiltersRequestBuilder) {
    m := &ItemCollectionsItemFiltersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/filters", pathParameters),
    }
    return m
}
// NewItemCollectionsItemFiltersRequestBuilder instantiates a new ItemCollectionsItemFiltersRequestBuilder and sets the default values.
func NewItemCollectionsItemFiltersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemFiltersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemFiltersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the list of available folder fields (taxonomies, smart filters) and their properties. Folder fields can be used for filtering a search result by restricting to documents having a given folder in a given folder field. Also values can be accumulated across documents having a certain folder in a given folder field.
// returns a FieldsResultable when successful
func (m *ItemCollectionsItemFiltersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.FieldsResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateFieldsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.FieldsResultable), nil
}
// Post returns the list of available folder fields (taxonomies, smart filters) and their properties. Folder fields can be used for filtering a search result by restricting to documents having a given folder in a given folder field. Also values can be accumulated across documents having a certain folder in a given folder field.
// returns a FieldsResultable when successful
func (m *ItemCollectionsItemFiltersRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.FieldsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateFieldsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.FieldsResultable), nil
}
// ToGetRequestInformation returns the list of available folder fields (taxonomies, smart filters) and their properties. Folder fields can be used for filtering a search result by restricting to documents having a given folder in a given folder field. Also values can be accumulated across documents having a certain folder in a given folder field.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemFiltersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation returns the list of available folder fields (taxonomies, smart filters) and their properties. Folder fields can be used for filtering a search result by restricting to documents having a given folder in a given folder field. Also values can be accumulated across documents having a certain folder in a given folder field.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemFiltersRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemFiltersRequestBuilder when successful
func (m *ItemCollectionsItemFiltersRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemFiltersRequestBuilder) {
    return NewItemCollectionsItemFiltersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
