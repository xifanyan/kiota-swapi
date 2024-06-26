package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemFiltersItemValuesRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\filters\{fieldId}\values
type ItemCollectionsItemFiltersItemValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemFiltersItemValuesRequestBuilderGetQueryParameters returns a list of folders (categories) in the given folder field (taxonomy, smart filter). Each resulting folder entry has a count of associated documents. The counts are based on the set of documents found for the given query and restrictions
type ItemCollectionsItemFiltersItemValuesRequestBuilderGetQueryParameters struct {
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Define a limit on the number of results returned. Default is 20, maximally allowed value is 1000.
    Limit *int32 `uriparametername:"limit"`
    // Define a relative position from which to start fetching results. Can be positive or zero (default).
    Offset *int32 `uriparametername:"offset"`
    // Provide an order criterion for the results. Needs to be one of 'count', 'relevance', 'name', 'name:asc', 'name:desc'. The latter three correspond to the display name.
    Order *string `uriparametername:"order"`
    // Optional prefix to constrain the folders returned. This is equivalent to restrictFoldersByQuery=rm_display_name=VALUE* where VALUE is the prefix of choice.
    Prefix *string `uriparametername:"prefix"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
    // Optional query to constrain the folders returned. Examples are rm_folder_id=uniquevalue or rm_display_name=displayname
    RestrictFoldersByQuery *string `uriparametername:"restrictFoldersByQuery"`
    // Defines if folders with count=0 should be returned. Default is false.
    ReturnEmptyFolders *bool `uriparametername:"returnEmptyFolders"`
}
// ItemCollectionsItemFiltersItemValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemFiltersItemValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemFiltersItemValuesRequestBuilderGetQueryParameters
}
// ItemCollectionsItemFiltersItemValuesRequestBuilderPostQueryParameters returns a list of folders (categories) in the given folder field (taxonomy, smart filter). Each resulting folder entry has a count of associated documents. The counts are based on the set of documents found for the given query and restrictions
type ItemCollectionsItemFiltersItemValuesRequestBuilderPostQueryParameters struct {
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Define a limit on the number of results returned. Default is 20, maximally allowed value is 1000.
    Limit *int32 `uriparametername:"limit"`
    // Define a relative position from which to start fetching results. Can be positive or zero (default).
    Offset *int32 `uriparametername:"offset"`
    // Provide an order criterion for the results. Needs to be one of 'count', 'relevance', 'name', 'name:asc', 'name:desc'. The latter three correspond to the display name.
    Order *string `uriparametername:"order"`
    // Optional prefix to constrain the folders returned. This is equivalent to restrictFoldersByQuery=rm_display_name=VALUE* where VALUE is the prefix of choice.
    Prefix *string `uriparametername:"prefix"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
    // Optional query to constrain the folders returned. Examples are rm_folder_id=uniquevalue or rm_display_name=displayname
    RestrictFoldersByQuery *string `uriparametername:"restrictFoldersByQuery"`
    // Defines if folders with count=0 should be returned. Default is false.
    ReturnEmptyFolders *bool `uriparametername:"returnEmptyFolders"`
}
// ItemCollectionsItemFiltersItemValuesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemFiltersItemValuesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemFiltersItemValuesRequestBuilderPostQueryParameters
}
// NewItemCollectionsItemFiltersItemValuesRequestBuilderInternal instantiates a new ItemCollectionsItemFiltersItemValuesRequestBuilder and sets the default values.
func NewItemCollectionsItemFiltersItemValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemFiltersItemValuesRequestBuilder) {
    m := &ItemCollectionsItemFiltersItemValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/filters/{fieldId}/values{?joinRestriction*,language*,limit*,offset*,order*,prefix*,query*,restrictFoldersByQuery*,returnEmptyFolders*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemFiltersItemValuesRequestBuilder instantiates a new ItemCollectionsItemFiltersItemValuesRequestBuilder and sets the default values.
func NewItemCollectionsItemFiltersItemValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemFiltersItemValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemFiltersItemValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of folders (categories) in the given folder field (taxonomy, smart filter). Each resulting folder entry has a count of associated documents. The counts are based on the set of documents found for the given query and restrictions
// returns a FolderValuesResultable when successful
func (m *ItemCollectionsItemFiltersItemValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersItemValuesRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.FolderValuesResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateFolderValuesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.FolderValuesResultable), nil
}
// Post returns a list of folders (categories) in the given folder field (taxonomy, smart filter). Each resulting folder entry has a count of associated documents. The counts are based on the set of documents found for the given query and restrictions
// returns a FolderValuesResultable when successful
func (m *ItemCollectionsItemFiltersItemValuesRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersItemValuesRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.FolderValuesResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateFolderValuesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.FolderValuesResultable), nil
}
// ToGetRequestInformation returns a list of folders (categories) in the given folder field (taxonomy, smart filter). Each resulting folder entry has a count of associated documents. The counts are based on the set of documents found for the given query and restrictions
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemFiltersItemValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersItemValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation returns a list of folders (categories) in the given folder field (taxonomy, smart filter). Each resulting folder entry has a count of associated documents. The counts are based on the set of documents found for the given query and restrictions
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemFiltersItemValuesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersItemValuesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCollectionsItemFiltersItemValuesRequestBuilder when successful
func (m *ItemCollectionsItemFiltersItemValuesRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemFiltersItemValuesRequestBuilder) {
    return NewItemCollectionsItemFiltersItemValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
