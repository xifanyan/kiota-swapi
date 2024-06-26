package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
    i08e471563dc6f380ccdcb6c265b0c002a2e545fd6bb1a0d00200bb47e85dd1ee "github.com/xifanyan/kiota-swapi/jobs/projects/item/collections/item/jobs/item/status"
)

// ProjectsItemCollectionsItemJobsItemStatusRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}\collections\{collectionId}\jobs\{jobId}\status
type ProjectsItemCollectionsItemJobsItemStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProjectsItemCollectionsItemJobsItemStatusRequestBuilderGetQueryParameters this endpoint is almost identical to jobs?ids=ID except that it also allows to wait for a job to reach a final status.
type ProjectsItemCollectionsItemJobsItemStatusRequestBuilderGetQueryParameters struct {
    // An option to wait until the selected job reaches a final status (included PAUSED). The argument is a timeout in milliseconds before the endpoint returns.
    WaitForCompletionTimeout *int64 `uriparametername:"waitForCompletionTimeout"`
}
// ProjectsItemCollectionsItemJobsItemStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsItemCollectionsItemJobsItemStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ProjectsItemCollectionsItemJobsItemStatusRequestBuilderGetQueryParameters
}
// ProjectsItemCollectionsItemJobsItemStatusRequestBuilderPutQueryParameters changes the status of a job
type ProjectsItemCollectionsItemJobsItemStatusRequestBuilderPutQueryParameters struct {
    // The requested status change.
    // Deprecated: This property is deprecated, use StatusAsPutStatusQueryParameterType instead
    Status *string `uriparametername:"status"`
    // The requested status change.
    StatusAsPutStatusQueryParameterType *i08e471563dc6f380ccdcb6c265b0c002a2e545fd6bb1a0d00200bb47e85dd1ee.PutStatusQueryParameterType `uriparametername:"status"`
}
// ProjectsItemCollectionsItemJobsItemStatusRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsItemCollectionsItemJobsItemStatusRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ProjectsItemCollectionsItemJobsItemStatusRequestBuilderPutQueryParameters
}
// NewProjectsItemCollectionsItemJobsItemStatusRequestBuilderInternal instantiates a new ProjectsItemCollectionsItemJobsItemStatusRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsItemStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) {
    m := &ProjectsItemCollectionsItemJobsItemStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}/jobs/{jobId}/status{?waitForCompletionTimeout*}", pathParameters),
    }
    return m
}
// NewProjectsItemCollectionsItemJobsItemStatusRequestBuilder instantiates a new ProjectsItemCollectionsItemJobsItemStatusRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsItemStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsItemCollectionsItemJobsItemStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this endpoint is almost identical to jobs?ids=ID except that it also allows to wait for a job to reach a final status.
// returns a JobDescriptionable when successful
func (m *ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsItemStatusRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.JobDescriptionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateJobDescriptionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.JobDescriptionable), nil
}
// Put changes the status of a job
// returns a JobDescriptionable when successful
func (m *ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) Put(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsItemStatusRequestBuilderPutRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.JobDescriptionable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateJobDescriptionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.JobDescriptionable), nil
}
// ToGetRequestInformation this endpoint is almost identical to jobs?ids=ID except that it also allows to wait for a job to reach a final status.
// returns a *RequestInformation when successful
func (m *ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsItemStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation changes the status of a job
// returns a *RequestInformation when successful
func (m *ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsItemStatusRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}/jobs/{jobId}/status?status={status}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ProjectsItemCollectionsItemJobsItemStatusRequestBuilder when successful
func (m *ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) WithUrl(rawUrl string)(*ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) {
    return NewProjectsItemCollectionsItemJobsItemStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
