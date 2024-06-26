package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemRecordsRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\records
type ItemCollectionsItemRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemRecordsRequestBuilderGetQueryParameters returns the list of records as specified by the parameters.
type ItemCollectionsItemRecordsRequestBuilderGetQueryParameters struct {
    // Indicate whether to retrieve summarized content for the record, default is false. This argument retrieves only parts of the content, namely the part which is relevant for the query or the first sentences of the document.
    Body *bool `uriparametername:"body"`
    // An URI-encoded, comma-separated list of fields to be retrieved for each document. Valid field names are typically those which can be inspected using the fields endpoint (i.e. those configured in the data model). In addition, search results support to retrieve dynamic properties as field name: 'rm_is_best_bet' returns true if and only if 'Best bets boosting' is active and a match is a best bet.
    Fields *string `uriparametername:"fields"`
    // An URI-encoded, comma-separated list of folder fields (taxonomies) to be retrieved for each document. Note that folder fields can be returned by 'fields' as well in which case an array of folder ids is returned. Use 'folderFields' if you need the folder and its display name.
    FolderFields *string `uriparametername:"folderFields"`
    // An argument to 'body': Indicate whether to provide highlighting information (via XML tags of the form <recomDescriptiveWord>, default is true
    Highlight *bool `uriparametername:"highlight"`
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Define a page size / limit on the number of results returned. Default is 20. The maximally allowed value is 1000 for all standard paging use-cases. However, if the 'accept' header is set to 'application/x-ndjson', the result will be returned in streaming mode which allows to specify page=1, limit=-1 in order to return all documents (see response documentation for details).
    Limit *int32 `uriparametername:"limit"`
    // Provide order criteria. Format - <field1>[:asc,:desc],field2>[:asc,:desc],...,<fieldN>[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.
    Order *string `uriparametername:"order"`
    // Define which page to retrieve (counting starts at 1). Default is 1.
    Page *int32 `uriparametername:"page"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
    // Computes and returns any configured spelling suggestions for the provided query. Note that this requires spelling suggestions to be configured in the project's data model. Default is to not compute spelling suggestions.
    SpellingSuggestions *string `uriparametername:"spellingSuggestions"`
    // Computes and returns any configured sponsored links for the provided query. Note that this requires sponsored links to be configured in the project's data model. Default is to not compute sponsored links.
    SponsoredLinks *string `uriparametername:"sponsoredLinks"`
}
// ItemCollectionsItemRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemRecordsRequestBuilderGetQueryParameters
}
// ItemCollectionsItemRecordsRequestBuilderPostQueryParameters returns the list of records as specified by the parameters.
type ItemCollectionsItemRecordsRequestBuilderPostQueryParameters struct {
    // Indicate whether to retrieve summarized content for the record, default is false. This argument retrieves only parts of the content, namely the part which is relevant for the query or the first sentences of the document.
    Body *bool `uriparametername:"body"`
    // An URI-encoded, comma-separated list of fields to be retrieved for each document. Valid field names are typically those which can be inspected using the fields endpoint (i.e. those configured in the data model). In addition, search results support to retrieve dynamic properties as field name: 'rm_is_best_bet' returns true if and only if 'Best bets boosting' is active and a match is a best bet.
    Fields *string `uriparametername:"fields"`
    // An URI-encoded, comma-separated list of folder fields (taxonomies) to be retrieved for each document. Note that folder fields can be returned by 'fields' as well in which case an array of folder ids is returned. Use 'folderFields' if you need the folder and its display name.
    FolderFields *string `uriparametername:"folderFields"`
    // An argument to 'body': Indicate whether to provide highlighting information (via XML tags of the form <recomDescriptiveWord>, default is true
    Highlight *bool `uriparametername:"highlight"`
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Define a page size / limit on the number of results returned. Default is 20. The maximally allowed value is 1000 for all standard paging use-cases. However, if the 'accept' header is set to 'application/x-ndjson', the result will be returned in streaming mode which allows to specify page=1, limit=-1 in order to return all documents (see response documentation for details).
    Limit *int32 `uriparametername:"limit"`
    // Provide order criteria. Format - <field1>[:asc,:desc],field2>[:asc,:desc],...,<fieldN>[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.
    Order *string `uriparametername:"order"`
    // Define which page to retrieve (counting starts at 1). Default is 1.
    Page *int32 `uriparametername:"page"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
    // Computes and returns any configured spelling suggestions for the provided query. Note that this requires spelling suggestions to be configured in the project's data model. Default is to not compute spelling suggestions.
    SpellingSuggestions *string `uriparametername:"spellingSuggestions"`
    // Computes and returns any configured sponsored links for the provided query. Note that this requires sponsored links to be configured in the project's data model. Default is to not compute sponsored links.
    SponsoredLinks *string `uriparametername:"sponsoredLinks"`
}
// ItemCollectionsItemRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemRecordsRequestBuilderPostQueryParameters
}
// ItemCollectionsItemRecordsRequestBuilderPutQueryParameters allows to change properties of documents, for example by tagging them into folders or by modifying textual content.
type ItemCollectionsItemRecordsRequestBuilderPutQueryParameters struct {
    // Specifies whether the call should block until all indexes are updated. The default value 'false' is to return immediately which means that the display of values is always correct, but search indices may be outdated until all indices have been updated eventually (that means: queries will not respect the change until the index is updated). The change is persisted as soon as the API call returns, irrespective of 'blockUntilComplete'.
    BlockUntilComplete *bool `uriparametername:"blockUntilComplete"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
}
// ItemCollectionsItemRecordsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemRecordsRequestBuilderPutQueryParameters
}
// ByRecordId gets an item from the github.com/xifanyan/kiota-swapi.projects.item.collections.item.records.item collection
// returns a *ItemCollectionsItemRecordsWithRecordItemRequestBuilder when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) ByRecordId(recordId string)(*ItemCollectionsItemRecordsWithRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if recordId != "" {
        urlTplParams["recordId"] = recordId
    }
    return NewItemCollectionsItemRecordsWithRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCollectionsItemRecordsRequestBuilderInternal instantiates a new ItemCollectionsItemRecordsRequestBuilder and sets the default values.
func NewItemCollectionsItemRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemRecordsRequestBuilder) {
    m := &ItemCollectionsItemRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/records{?body*,fields*,folderFields*,highlight*,joinRestriction*,language*,limit*,order*,page*,query*,spellingSuggestions*,sponsoredLinks*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemRecordsRequestBuilder instantiates a new ItemCollectionsItemRecordsRequestBuilder and sets the default values.
func NewItemCollectionsItemRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the list of records as specified by the parameters.
// returns a SearchResultable when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateSearchResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultable), nil
}
// InsertRemoveTransaction the insertRemoveTransaction property
// returns a *ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) InsertRemoveTransaction()(*ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder) {
    return NewItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post returns the list of records as specified by the parameters.
// returns a SearchResultable when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateSearchResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultable), nil
}
// Put allows to change properties of documents, for example by tagging them into folders or by modifying textual content.
// returns a ChangeResultable when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) Put(ctx context.Context, body []ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ChangeRequestable, requestConfiguration *ItemCollectionsItemRecordsRequestBuilderPutRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ChangeResultable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateChangeResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ChangeResultable), nil
}
// ToGetRequestInformation returns the list of records as specified by the parameters.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation returns the list of records as specified by the parameters.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation allows to change properties of documents, for example by tagging them into folders or by modifying textual content.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) ToPutRequestInformation(ctx context.Context, body []ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ChangeRequestable, requestConfiguration *ItemCollectionsItemRecordsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/projects/{projectId}/collections/{collectionId}/records?blockUntilComplete={blockUntilComplete}{&language*,query*}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(body))
    for i, v := range body {
        if v != nil {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
    }
    err := requestInfo.SetContentFromParsableCollection(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", cast)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemRecordsRequestBuilder when successful
func (m *ItemCollectionsItemRecordsRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemRecordsRequestBuilder) {
    return NewItemCollectionsItemRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
