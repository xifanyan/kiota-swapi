package jobs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}\collections\{collectionId}\jobs\{jobId}
type ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewProjectsItemCollectionsItemJobsWithJobItemRequestBuilderInternal instantiates a new ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsWithJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder) {
    m := &ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}/jobs/{jobId}", pathParameters),
    }
    return m
}
// NewProjectsItemCollectionsItemJobsWithJobItemRequestBuilder instantiates a new ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemJobsWithJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsItemCollectionsItemJobsWithJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ErrorFile the errorFile property
// returns a *ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder when successful
func (m *ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder) ErrorFile()(*ProjectsItemCollectionsItemJobsItemErrorFileRequestBuilder) {
    return NewProjectsItemCollectionsItemJobsItemErrorFileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResultFile the resultFile property
// returns a *ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder when successful
func (m *ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder) ResultFile()(*ProjectsItemCollectionsItemJobsItemResultFileRequestBuilder) {
    return NewProjectsItemCollectionsItemJobsItemResultFileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Status the status property
// returns a *ProjectsItemCollectionsItemJobsItemStatusRequestBuilder when successful
func (m *ProjectsItemCollectionsItemJobsWithJobItemRequestBuilder) Status()(*ProjectsItemCollectionsItemJobsItemStatusRequestBuilder) {
    return NewProjectsItemCollectionsItemJobsItemStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
