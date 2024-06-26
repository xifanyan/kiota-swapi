package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemMeasuresRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\measures
type ItemCollectionsItemMeasuresRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemMeasuresRequestBuilderPostQueryParameters returns the measure values. For zero-dimensional measures, it just computes a single value as aggregate over the entire search result. For one-dimensional measures, the result resembles a sorted table in which measure dimensions (like folders) are associated with the measure. For two-dimensional measures, the result is a matrix (a cube) which returns measure value associated with a pair of values. The most simple measure is the count of documents which fall into one of the buckets.
type ItemCollectionsItemMeasuresRequestBuilderPostQueryParameters struct {
    // Specify a restriction using a joined collection. The argument is expected to be of the form <targetFieldID>:<query>. Example: doc_author_id:office="New York" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The <targetFieldID> id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The <query> is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=<targetFieldID>:<query> and the query syntax <targetFieldID>find(<query>) are equivalent. 3. For meta engine projects, <targetFieldID> is valid independent of the collectionId (i.e. you can use collectionId=people and a <targetFieldID> which is in the people collection).
    JoinRestriction *string `uriparametername:"joinRestriction"`
    // Specify a language to be used when interpreting the query, default is to use language recognition.
    Language *string `uriparametername:"language"`
    // Defines the aggregate type. The default is to compute counts (i.e. a MeasureTypeParameter with typeName 'count'). All other arguments need a JSON representation of MeasureTypeParameter (url encoded json).
    MeasureType *string `uriparametername:"measureType"`
    // The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '"An explicit phrase"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '<', '=', '>', '<=', '>=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. 
    Query *string `uriparametername:"query"`
}
// ItemCollectionsItemMeasuresRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemMeasuresRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemMeasuresRequestBuilderPostQueryParameters
}
// NewItemCollectionsItemMeasuresRequestBuilderInternal instantiates a new ItemCollectionsItemMeasuresRequestBuilder and sets the default values.
func NewItemCollectionsItemMeasuresRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemMeasuresRequestBuilder) {
    m := &ItemCollectionsItemMeasuresRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/measures{?joinRestriction*,language*,measureType*,query*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemMeasuresRequestBuilder instantiates a new ItemCollectionsItemMeasuresRequestBuilder and sets the default values.
func NewItemCollectionsItemMeasuresRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemMeasuresRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemMeasuresRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns the measure values. For zero-dimensional measures, it just computes a single value as aggregate over the entire search result. For one-dimensional measures, the result resembles a sorted table in which measure dimensions (like folders) are associated with the measure. For two-dimensional measures, the result is a matrix (a cube) which returns measure value associated with a pair of values. The most simple measure is the count of documents which fall into one of the buckets.
// returns a MeasureCubeable when successful
func (m *ItemCollectionsItemMeasuresRequestBuilder) Post(ctx context.Context, body []ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.DimensionRequestable, requestConfiguration *ItemCollectionsItemMeasuresRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.MeasureCubeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateMeasureCubeFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.MeasureCubeable), nil
}
// ToPostRequestInformation returns the measure values. For zero-dimensional measures, it just computes a single value as aggregate over the entire search result. For one-dimensional measures, the result resembles a sorted table in which measure dimensions (like folders) are associated with the measure. For two-dimensional measures, the result is a matrix (a cube) which returns measure value associated with a pair of values. The most simple measure is the count of documents which fall into one of the buckets.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemMeasuresRequestBuilder) ToPostRequestInformation(ctx context.Context, body []ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.DimensionRequestable, requestConfiguration *ItemCollectionsItemMeasuresRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemCollectionsItemMeasuresRequestBuilder when successful
func (m *ItemCollectionsItemMeasuresRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemMeasuresRequestBuilder) {
    return NewItemCollectionsItemMeasuresRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
