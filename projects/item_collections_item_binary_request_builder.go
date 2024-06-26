package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCollectionsItemBinaryRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\binary
type ItemCollectionsItemBinaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemBinaryRequestBuilderGetQueryParameters returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
type ItemCollectionsItemBinaryRequestBuilderGetQueryParameters struct {
    // A single field name storing binary data.
    Field *string `uriparametername:"field"`
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Provide order criteria. Format - <field1>[:asc,:desc],field2>[:asc,:desc],...,<fieldN>[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.
    Order *string `uriparametername:"order"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
    // Define the stream to select (counting starts at 1). This is useful if and only if the query returns more than one result. Default is 1.
    SelectedIndex *int32 `uriparametername:"selectedIndex"`
}
// ItemCollectionsItemBinaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemBinaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemBinaryRequestBuilderGetQueryParameters
}
// ItemCollectionsItemBinaryRequestBuilderPostQueryParameters returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
type ItemCollectionsItemBinaryRequestBuilderPostQueryParameters struct {
    // A single field name storing binary data.
    Field *string `uriparametername:"field"`
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Provide order criteria. Format - <field1>[:asc,:desc],field2>[:asc,:desc],...,<fieldN>[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.
    Order *string `uriparametername:"order"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
    // Define the stream to select (counting starts at 1). This is useful if and only if the query returns more than one result. Default is 1.
    SelectedIndex *int32 `uriparametername:"selectedIndex"`
}
// ItemCollectionsItemBinaryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemBinaryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemBinaryRequestBuilderPostQueryParameters
}
// ByRecordId gets an item from the github.com/xifanyan/kiota-swapi.projects.item.collections.item.binary.item collection
// returns a *ItemCollectionsItemBinaryWithRecordItemRequestBuilder when successful
func (m *ItemCollectionsItemBinaryRequestBuilder) ByRecordId(recordId string)(*ItemCollectionsItemBinaryWithRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if recordId != "" {
        urlTplParams["recordId"] = recordId
    }
    return NewItemCollectionsItemBinaryWithRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCollectionsItemBinaryRequestBuilderInternal instantiates a new ItemCollectionsItemBinaryRequestBuilder and sets the default values.
func NewItemCollectionsItemBinaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemBinaryRequestBuilder) {
    m := &ItemCollectionsItemBinaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/binary{?field*,joinRestriction*,language*,order*,query*,selectedIndex*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemBinaryRequestBuilder instantiates a new ItemCollectionsItemBinaryRequestBuilder and sets the default values.
func NewItemCollectionsItemBinaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemBinaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemBinaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.
// returns a []byte when successful
func (m *ItemCollectionsItemBinaryRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemBinaryRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *ItemCollectionsItemBinaryRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemBinaryRequestBuilderPostRequestConfiguration)([]byte, error) {
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
func (m *ItemCollectionsItemBinaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemBinaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemCollectionsItemBinaryRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemBinaryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCollectionsItemBinaryRequestBuilder when successful
func (m *ItemCollectionsItemBinaryRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemBinaryRequestBuilder) {
    return NewItemCollectionsItemBinaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
