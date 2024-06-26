package jobs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProjectsWithProjectItemRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}
type ProjectsWithProjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Collections the collections property
// returns a *ProjectsItemCollectionsRequestBuilder when successful
func (m *ProjectsWithProjectItemRequestBuilder) Collections()(*ProjectsItemCollectionsRequestBuilder) {
    return NewProjectsItemCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewProjectsWithProjectItemRequestBuilderInternal instantiates a new ProjectsWithProjectItemRequestBuilder and sets the default values.
func NewProjectsWithProjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsWithProjectItemRequestBuilder) {
    m := &ProjectsWithProjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}", pathParameters),
    }
    return m
}
// NewProjectsWithProjectItemRequestBuilder instantiates a new ProjectsWithProjectItemRequestBuilder and sets the default values.
func NewProjectsWithProjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsWithProjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsWithProjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
