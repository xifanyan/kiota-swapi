package jobs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProjectsItemCollectionsRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}\collections
type ProjectsItemCollectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCollectionId gets an item from the github.com/xifanyan/kiota-swapi.jobs.projects.item.collections.item collection
// returns a *ProjectsItemCollectionsWithCollectionItemRequestBuilder when successful
func (m *ProjectsItemCollectionsRequestBuilder) ByCollectionId(collectionId string)(*ProjectsItemCollectionsWithCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if collectionId != "" {
        urlTplParams["collectionId"] = collectionId
    }
    return NewProjectsItemCollectionsWithCollectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewProjectsItemCollectionsRequestBuilderInternal instantiates a new ProjectsItemCollectionsRequestBuilder and sets the default values.
func NewProjectsItemCollectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsRequestBuilder) {
    m := &ProjectsItemCollectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}/collections", pathParameters),
    }
    return m
}
// NewProjectsItemCollectionsRequestBuilder instantiates a new ProjectsItemCollectionsRequestBuilder and sets the default values.
func NewProjectsItemCollectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsItemCollectionsRequestBuilderInternal(urlParams, requestAdapter)
}
