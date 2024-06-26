package projects

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCollectionsItemBinaryWithRecordItemRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\binary\{recordId}
type ItemCollectionsItemBinaryWithRecordItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCollectionsItemBinaryWithRecordItemRequestBuilderInternal instantiates a new ItemCollectionsItemBinaryWithRecordItemRequestBuilder and sets the default values.
func NewItemCollectionsItemBinaryWithRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemBinaryWithRecordItemRequestBuilder) {
    m := &ItemCollectionsItemBinaryWithRecordItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/binary/{recordId}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemBinaryWithRecordItemRequestBuilder instantiates a new ItemCollectionsItemBinaryWithRecordItemRequestBuilder and sets the default values.
func NewItemCollectionsItemBinaryWithRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemBinaryWithRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemBinaryWithRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
// returns a *ItemCollectionsItemBinaryItemContentRequestBuilder when successful
func (m *ItemCollectionsItemBinaryWithRecordItemRequestBuilder) Content()(*ItemCollectionsItemBinaryItemContentRequestBuilder) {
    return NewItemCollectionsItemBinaryItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
