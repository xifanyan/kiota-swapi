package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ProjectsRequestBuilder builds and executes requests for operations under \projects
type ProjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProjectsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ProjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByProjectId gets an item from the github.com/xifanyan/kiota-swapi.projects.item collection
// returns a *WithProjectItemRequestBuilder when successful
func (m *ProjectsRequestBuilder) ByProjectId(projectId string)(*WithProjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if projectId != "" {
        urlTplParams["projectId"] = projectId
    }
    return NewWithProjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewProjectsRequestBuilderInternal instantiates a new ProjectsRequestBuilder and sets the default values.
func NewProjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsRequestBuilder) {
    m := &ProjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects", pathParameters),
    }
    return m
}
// NewProjectsRequestBuilder instantiates a new ProjectsRequestBuilder and sets the default values.
func NewProjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get discover available projects, i.e. all projects that are present in an installation.
// returns a ProjectsResultable when successful
func (m *ProjectsRequestBuilder) Get(ctx context.Context, requestConfiguration *ProjectsRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ProjectsResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateProjectsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ProjectsResultable), nil
}
// Post discover available projects, i.e. all projects that are present in an installation.
// returns a ProjectsResultable when successful
func (m *ProjectsRequestBuilder) Post(ctx context.Context, requestConfiguration *ProjectsRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ProjectsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateProjectsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ProjectsResultable), nil
}
// ToGetRequestInformation discover available projects, i.e. all projects that are present in an installation.
// returns a *RequestInformation when successful
func (m *ProjectsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProjectsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation discover available projects, i.e. all projects that are present in an installation.
// returns a *RequestInformation when successful
func (m *ProjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ProjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProjectsRequestBuilder when successful
func (m *ProjectsRequestBuilder) WithUrl(rawUrl string)(*ProjectsRequestBuilder) {
    return NewProjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
