package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemSearchTokenRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\searchToken
type ItemCollectionsItemSearchTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemSearchTokenRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemSearchTokenRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCollectionsItemSearchTokenRequestBuilderGetQueryParameters returns a referencable search token for the search.
type ItemCollectionsItemSearchTokenRequestBuilderGetQueryParameters struct {
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Provide an optional order criteria. In most cases, search tokens need no ordering as they are simply used to define a set of documents. Format - <field1>[:asc,:desc],field2>[:asc,:desc],...,<fieldN>[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.
    Order *string `uriparametername:"order"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
}
// ItemCollectionsItemSearchTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemSearchTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemSearchTokenRequestBuilderGetQueryParameters
}
// ItemCollectionsItemSearchTokenRequestBuilderPostQueryParameters returns a referencable search token for the search.
type ItemCollectionsItemSearchTokenRequestBuilderPostQueryParameters struct {
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Provide an optional order criteria. In most cases, search tokens need no ordering as they are simply used to define a set of documents. Format - <field1>[:asc,:desc],field2>[:asc,:desc],...,<fieldN>[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.
    Order *string `uriparametername:"order"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
}
// ItemCollectionsItemSearchTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemSearchTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemSearchTokenRequestBuilderPostQueryParameters
}
// ItemCollectionsItemSearchTokenRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemSearchTokenRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCollectionsItemSearchTokenRequestBuilderInternal instantiates a new ItemCollectionsItemSearchTokenRequestBuilder and sets the default values.
func NewItemCollectionsItemSearchTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemSearchTokenRequestBuilder) {
    m := &ItemCollectionsItemSearchTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/searchToken{?joinRestriction*,language*,order*,query*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemSearchTokenRequestBuilder instantiates a new ItemCollectionsItemSearchTokenRequestBuilder and sets the default values.
func NewItemCollectionsItemSearchTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemSearchTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemSearchTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a search token.
// returns a []byte when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) Delete(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenable, requestConfiguration *ItemCollectionsItemSearchTokenRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
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
// Get returns a referencable search token for the search.
// returns a SearchResultTokenResponseable when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemSearchTokenRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateSearchResultTokenResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenResponseable), nil
}
// Post returns a referencable search token for the search.
// returns a SearchResultTokenResponseable when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemSearchTokenRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateSearchResultTokenResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenResponseable), nil
}
// Put renews a search token's life time.
// returns a SearchResultTokenResponseable when successful
// returns a SearchResultTokenResponse error when the service returns a 410 status code
func (m *ItemCollectionsItemSearchTokenRequestBuilder) Put(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenable, requestConfiguration *ItemCollectionsItemSearchTokenRequestBuilderPutRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "410": ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateSearchResultTokenResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateSearchResultTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenResponseable), nil
}
// SortOrderSnapshot the sortOrderSnapshot property
// returns a *ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) SortOrderSnapshot()(*ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder) {
    return NewItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a search token.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenable, requestConfiguration *ItemCollectionsItemSearchTokenRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/projects/{projectId}/collections/{collectionId}/searchToken", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToGetRequestInformation returns a referencable search token for the search.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemSearchTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation returns a referencable search token for the search.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemSearchTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation renews a search token's life time.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) ToPutRequestInformation(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenable, requestConfiguration *ItemCollectionsItemSearchTokenRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/projects/{projectId}/collections/{collectionId}/searchToken", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemSearchTokenRequestBuilder when successful
func (m *ItemCollectionsItemSearchTokenRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemSearchTokenRequestBuilder) {
    return NewItemCollectionsItemSearchTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
