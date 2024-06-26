package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ProjectsItemCollectionsItemJobsRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}\collections\{collectionId}\jobs
type ProjectsItemCollectionsItemJobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProjectsItemCollectionsItemJobsRequestBuilderGetQueryParameters connections against the record collection identified by projectId/collectionId and retrieves monitoring information about jobs.
type ProjectsItemCollectionsItemJobsRequestBuilderGetQueryParameters struct {
    // A comma-separated list of job ids for which information is to be retrieved. If absent, all jobs will be in scope.
    Ids *string `uriparametername:"ids"`
    // A boolean argument which controls if statusSelection is interpreted as 'return jobs with one of the specified status (invertStatusSelection=true, the default) or as 'return jobs which have none of the specified job status (invertStatusSelection=false)
    InvertStatusSelection *bool `uriparametername:"invertStatusSelection"`
    // Restricts the returned job descriptions to those with matching job status. The argument is a comma-separated list of job status like RUNNING or FINISHED.
    StatusSelection *string `uriparametername:"statusSelection"`
}
// ProjectsItemCollectionsItemJobsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsItemCollectionsItemJobsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ProjectsItemCollectionsItemJobsRequestBuilderGetQueryParameters
}
// ProjectsItemCollectionsItemJobsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsItemCollectionsItemJobsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByJobId gets an item from the github.com/xifanyan/kiota-swapi.jobs.projects.item.collections.item.jobs.item collection
// returns a *ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder when successful
func (m *ProjectsItemCollectionsItemJobsRequestBuilder) ByJobId(jobId string)(*ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if jobId != "" {
        urlTplParams["jobId"] = jobId
    }
    return NewProjectsItemCollectionsItemJobsWithJobItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewProjectsItemCollectionsItemJobsRequestBuilderInternal instantiates a new ProjectsItemCollectionsItemJobsRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsRequestBuilder) {
    m := &ProjectsItemCollectionsItemJobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}/jobs{?ids*,invertStatusSelection*,statusSelection*}", pathParameters),
    }
    return m
}
// NewProjectsItemCollectionsItemJobsRequestBuilder instantiates a new ProjectsItemCollectionsItemJobsRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsItemCollectionsItemJobsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get connections against the record collection identified by projectId/collectionId and retrieves monitoring information about jobs.
// returns a JobDescriptionListable when successful
func (m *ProjectsItemCollectionsItemJobsRequestBuilder) Get(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.JobDescriptionListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateJobDescriptionListFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.JobDescriptionListable), nil
}
// Post submits a job identified by a name and key-value pairs. The job name must be registered inside of the server. Available jobs can be retrieved using the endpoint /registeredJobTypes.
// returns a *string when successful
func (m *ProjectsItemCollectionsItemJobsRequestBuilder) Post(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.JobSubmissionRequestable, requestConfiguration *ProjectsItemCollectionsItemJobsRequestBuilderPostRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*string), nil
}
// ToGetRequestInformation connections against the record collection identified by projectId/collectionId and retrieves monitoring information about jobs.
// returns a *RequestInformation when successful
func (m *ProjectsItemCollectionsItemJobsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation submits a job identified by a name and key-value pairs. The job name must be registered inside of the server. Available jobs can be retrieved using the endpoint /registeredJobTypes.
// returns a *RequestInformation when successful
func (m *ProjectsItemCollectionsItemJobsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.JobSubmissionRequestable, requestConfiguration *ProjectsItemCollectionsItemJobsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}/jobs", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProjectsItemCollectionsItemJobsRequestBuilder when successful
func (m *ProjectsItemCollectionsItemJobsRequestBuilder) WithUrl(rawUrl string)(*ProjectsItemCollectionsItemJobsRequestBuilder) {
    return NewProjectsItemCollectionsItemJobsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
