package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}\collections\{collectionId}\jobs\{jobId}\resultFile
type ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProjectsItemCollectionsItemJobsItemResultFileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsItemCollectionsItemJobsItemResultFileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewProjectsItemCollectionsItemJobsItemResultFileRequestBuilderInternal instantiates a new ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsItemResultFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder) {
    m := &ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}/jobs/{jobId}/resultFile", pathParameters),
    }
    return m
}
// NewProjectsItemCollectionsItemJobsItemResultFileRequestBuilder instantiates a new ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsItemResultFileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsItemCollectionsItemJobsItemResultFileRequestBuilderInternal(urlParams, requestAdapter)
}
// Get for jobs with JobDescription.hasJobResultFile, this endpoint returns the job result file.
// returns a []byte when successful
func (m *ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder) Get(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsItemResultFileRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation for jobs with JobDescription.hasJobResultFile, this endpoint returns the job result file.
// returns a *RequestInformation when successful
func (m *ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsItemResultFileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder when successful
func (m *ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder) WithUrl(rawUrl string)(*ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder) {
    return NewProjectsItemCollectionsItemJobsItemResultFileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
