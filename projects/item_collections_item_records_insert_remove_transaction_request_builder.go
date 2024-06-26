package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\records\insertRemoveTransaction
type ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderPostQueryParameters creates new records in the specified collection. Note that this end point is intended to insert folders (=categories) into folder collections (=taxonomies). To this end, the projectId should be something like singleMindServer.NAME (or metaMindServer.NAME) and the collectionId should resemble the folder collection name (=taxonomy id). 
type ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderPostQueryParameters struct {
    Binaries []string `uriparametername:"binaries"`
    Request *string `uriparametername:"request"`
}
// ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderPostQueryParameters
}
// NewItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderInternal instantiates a new ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder and sets the default values.
func NewItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder) {
    m := &ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/records/insertRemoveTransaction?binaries={binaries}&request={request}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder instantiates a new ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder and sets the default values.
func NewItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates new records in the specified collection. Note that this end point is intended to insert folders (=categories) into folder collections (=taxonomies). To this end, the projectId should be something like singleMindServer.NAME (or metaMindServer.NAME) and the collectionId should resemble the folder collection name (=taxonomy id). 
// returns a InsertRemoveResultable when successful
func (m *ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder) Post(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.InsertRemoveRequestable, requestConfiguration *ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.InsertRemoveResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateInsertRemoveResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.InsertRemoveResultable), nil
}
// ToPostRequestInformation creates new records in the specified collection. Note that this end point is intended to insert folders (=categories) into folder collections (=taxonomies). To this end, the projectId should be something like singleMindServer.NAME (or metaMindServer.NAME) and the collectionId should resemble the folder collection name (=taxonomy id). 
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.InsertRemoveRequestable, requestConfiguration *ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
// returns a *ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder when successful
func (m *ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder) {
    return NewItemCollectionsItemRecordsInsertRemoveTransactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
