package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}\collections\{collectionId}\jobs\{jobId}\errorFile
type ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewProjectsItemCollectionsItemJobsItemErrorFileRequestBuilderInternal instantiates a new ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsItemErrorFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder) {
    m := &ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}/jobs/{jobId}/errorFile", pathParameters),
    }
    return m
}
// NewProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder instantiates a new ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsItemCollectionsItemJobsItemErrorFileRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this endpoint is suitable for jobs with errors (or which ended in a FATAL state). It returns the complete list of errors encountered so far. The endpoint returns values if and only if JobDescription.hasErrorFile=true
// returns a []byte when successful
func (m *ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder) Get(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation this endpoint is suitable for jobs with errors (or which ended in a FATAL state). It returns the complete list of errors encountered so far. The endpoint returns values if and only if JobDescription.hasErrorFile=true
// returns a *RequestInformation when successful
func (m *ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder when successful
func (m *ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder) WithUrl(rawUrl string)(*ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder) {
    return NewProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
