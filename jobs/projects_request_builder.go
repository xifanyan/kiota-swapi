package jobs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProjectsRequestBuilder builds and executes requests for operations under \jobs\projects
type ProjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByProjectId gets an item from the github.com/xifanyan/kiota-swapi.jobs.projects.item collection
// returns a *ProjectsWithProjectItemRequestBuilder when successful
func (m *ProjectsRequestBuilder) ByProjectId(projectId string)(*ProjectsWithProjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if projectId != "" {
        urlTplParams["projectId"] = projectId
    }
    return NewProjectsWithProjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewProjectsRequestBuilderInternal instantiates a new ProjectsRequestBuilder and sets the default values.
func NewProjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsRequestBuilder) {
    m := &ProjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects", pathParameters),
    }
    return m
}
// NewProjectsRequestBuilder instantiates a new ProjectsRequestBuilder and sets the default values.
func NewProjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsRequestBuilderInternal(urlParams, requestAdapter)
}
