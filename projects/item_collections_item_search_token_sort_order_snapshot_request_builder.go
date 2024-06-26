package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\searchToken\sortOrderSnapshot
type ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderPostQueryParameters computes a sort order snapshot of the specified search result token and stores it for the token. A 'sort order snapshot' is an atomically created sorted list of documents of the token, sorted according to its ordering. Its purpose is to allow a consistent sort order even if the sort keys are subject to concurrent modification. Keep in mind that mindserver always uses partial sort when requesting a page from a search result, i.e. any modifications to sort keys between two page fetches can and will affect the resulting sort order! The sort order snapshot allows to create a presorted consistent result.    Sort order snapshots can and need to be dereferenced explicitly. The suggested workflow is as follows:  <ol>  <li>search result token A is based on an ordering which is known to be subject to concurrent  modification,</li>  <li>you explicitly create a sort order snapshot for token A,</li>  <li>search result token B is based on a field-based-search expression including SEARCH_IN_SEARCHRESULT_RANGE referencing token A using a range of your choice. example:  SEARCH_IN_SEARCHRESULT_RANGE=A;0;10 . In this context, the results returned from token A use the sort order snapshot.</li>  </ol>  Thus, a sort order snapshots allows to access the ordering by means of this field-based-search expression.  It does _not_ allow to access the scores of A.   The sorted list of documents is stored in memory, next to the token's state. It is cleared if the token is closed or the engine gets restarted.    There can be at most one sort order snapshot per token; any previously existing sort order snap shot  will be overwritten.
type ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderPostQueryParameters struct {
    // an optional limit describing how many documents should become part of the sort order snapshot. A negative value (or a value which is larger than the search result) include all documents. All other values restrict the sort order snapshot to the N top-ranked documents of the last search. Consequently, usage of the sort order snapshot will return at most these N top ranked documents. Default is to use the entire search result.
    TopN *int64 `uriparametername:"topN"`
}
// ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderPostQueryParameters
}
// NewItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderInternal instantiates a new ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder and sets the default values.
func NewItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder) {
    m := &ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/searchToken/sortOrderSnapshot{?topN*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder instantiates a new ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder and sets the default values.
func NewItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderInternal(urlParams, requestAdapter)
}
// Post computes a sort order snapshot of the specified search result token and stores it for the token. A 'sort order snapshot' is an atomically created sorted list of documents of the token, sorted according to its ordering. Its purpose is to allow a consistent sort order even if the sort keys are subject to concurrent modification. Keep in mind that mindserver always uses partial sort when requesting a page from a search result, i.e. any modifications to sort keys between two page fetches can and will affect the resulting sort order! The sort order snapshot allows to create a presorted consistent result.    Sort order snapshots can and need to be dereferenced explicitly. The suggested workflow is as follows:  <ol>  <li>search result token A is based on an ordering which is known to be subject to concurrent  modification,</li>  <li>you explicitly create a sort order snapshot for token A,</li>  <li>search result token B is based on a field-based-search expression including SEARCH_IN_SEARCHRESULT_RANGE referencing token A using a range of your choice. example:  SEARCH_IN_SEARCHRESULT_RANGE=A;0;10 . In this context, the results returned from token A use the sort order snapshot.</li>  </ol>  Thus, a sort order snapshots allows to access the ordering by means of this field-based-search expression.  It does _not_ allow to access the scores of A.   The sorted list of documents is stored in memory, next to the token's state. It is cleared if the token is closed or the engine gets restarted.    There can be at most one sort order snapshot per token; any previously existing sort order snap shot  will be overwritten.
// returns a []byte when successful
func (m *ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder) Post(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenable, requestConfiguration *ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation computes a sort order snapshot of the specified search result token and stores it for the token. A 'sort order snapshot' is an atomically created sorted list of documents of the token, sorted according to its ordering. Its purpose is to allow a consistent sort order even if the sort keys are subject to concurrent modification. Keep in mind that mindserver always uses partial sort when requesting a page from a search result, i.e. any modifications to sort keys between two page fetches can and will affect the resulting sort order! The sort order snapshot allows to create a presorted consistent result.    Sort order snapshots can and need to be dereferenced explicitly. The suggested workflow is as follows:  <ol>  <li>search result token A is based on an ordering which is known to be subject to concurrent  modification,</li>  <li>you explicitly create a sort order snapshot for token A,</li>  <li>search result token B is based on a field-based-search expression including SEARCH_IN_SEARCHRESULT_RANGE referencing token A using a range of your choice. example:  SEARCH_IN_SEARCHRESULT_RANGE=A;0;10 . In this context, the results returned from token A use the sort order snapshot.</li>  </ol>  Thus, a sort order snapshots allows to access the ordering by means of this field-based-search expression.  It does _not_ allow to access the scores of A.   The sorted list of documents is stored in memory, next to the token's state. It is cleared if the token is closed or the engine gets restarted.    There can be at most one sort order snapshot per token; any previously existing sort order snap shot  will be overwritten.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder) ToPostRequestInformation(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.SearchResultTokenable, requestConfiguration *ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder when successful
func (m *ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder) {
    return NewItemCollectionsItemSearchTokenSortOrderSnapshotRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
