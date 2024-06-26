package jobs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProjectsItemCollectionsWithCollectionItemRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}\collections\{collectionId}
type ProjectsItemCollectionsWithCollectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewProjectsItemCollectionsWithCollectionItemRequestBuilderInternal instantiates a new ProjectsItemCollectionsWithCollectionItemRequestBuilder and sets the default values.
func NewProjectsItemCollectionsWithCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsWithCollectionItemRequestBuilder) {
    m := &ProjectsItemCollectionsWithCollectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}", pathParameters),
    }
    return m
}
// NewProjectsItemCollectionsWithCollectionItemRequestBuilder instantiates a new ProjectsItemCollectionsWithCollectionItemRequestBuilder and sets the default values.
func NewProjectsItemCollectionsWithCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsWithCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsItemCollectionsWithCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Jobs the jobs property
// returns a *ProjectsItemCollectionsItemJobsRequestBuilder when successful
func (m *ProjectsItemCollectionsWithCollectionItemRequestBuilder) Jobs()(*ProjectsItemCollectionsItemJobsRequestBuilder) {
    return NewProjectsItemCollectionsItemJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RegisteredJobTypes the registeredJobTypes property
// returns a *ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder when successful
func (m *ProjectsItemCollectionsWithCollectionItemRequestBuilder) RegisteredJobTypes()(*ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder) {
    return NewProjectsItemCollectionsItemRegisteredJobTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
