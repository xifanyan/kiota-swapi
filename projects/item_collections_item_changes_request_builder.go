package projects

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCollectionsItemChangesRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\changes
type ItemCollectionsItemChangesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCollectionsItemChangesRequestBuilderInternal instantiates a new ItemCollectionsItemChangesRequestBuilder and sets the default values.
func NewItemCollectionsItemChangesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemChangesRequestBuilder) {
    m := &ItemCollectionsItemChangesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/changes", pathParameters),
    }
    return m
}
// NewItemCollectionsItemChangesRequestBuilder instantiates a new ItemCollectionsItemChangesRequestBuilder and sets the default values.
func NewItemCollectionsItemChangesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemChangesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemChangesRequestBuilderInternal(urlParams, requestAdapter)
}
// Queue the queue property
// returns a *ItemCollectionsItemChangesQueueRequestBuilder when successful
func (m *ItemCollectionsItemChangesRequestBuilder) Queue()(*ItemCollectionsItemChangesQueueRequestBuilder) {
    return NewItemCollectionsItemChangesQueueRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
