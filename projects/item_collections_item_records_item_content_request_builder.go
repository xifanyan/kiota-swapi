package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437 "github.com/xifanyan/kiota-swapi/models"
)

// ItemCollectionsItemRecordsItemContentRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\records\{recordId}\content
type ItemCollectionsItemRecordsItemContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemRecordsItemContentRequestBuilderGetQueryParameters returns selected properties (possibly including paged body content) for the record in question.
type ItemCollectionsItemRecordsItemContentRequestBuilderGetQueryParameters struct {
    // Indicate whether to retrieve content for the record, default is false. This parameter is required in order to display highlighting results. It fetches fields configured as 'Content Fields' in the data model, applies highlighting, paging, and summarization as needed, and returns the resulting XML.
    Body *bool `uriparametername:"body"`
    // An URI-encoded, comma-separated list of fields to be retrieved for each document. Valid field names are typically those which can be inspected using the fields endpoint (i.e. those configured in the data model). In addition, search results support to retrieve dynamic properties as field name: 'rm_is_best_bet' returns true if and only if 'Best bets boosting' is active and a match is a best bet.
    Fields *string `uriparametername:"fields"`
    // An URI-encoded, comma-separated list of fields which will be fetched in XML form for which highlighting has been applied. The resulting field content is in XML form with <recomDescriptiveWord> tags for highlighted words. The response object will contain one field for each item in this list, and the result items will be have the suffix 'Xml' to indicate that the highlighted value is in XML form (example 'titleXml'). Note that fields specified in this context do not necessarily need to be known in the data model. Note furthermore that highlighting will only match for indexed fields.
    FieldsHighlighted *string `uriparametername:"fieldsHighlighted"`
    // An URI-encoded, comma-separated list of folder fields (taxonomies) to be retrieved for each document. Note that folder fields can be returned by 'fields' as well in which case an array of folder ids is returned. Use 'folderFields' if you need the folder and its display name.
    FolderFields *string `uriparametername:"folderFields"`
    // An argument to 'body': A comma-separated list of folder field names for which folder-specific highlighting is to be applied.
    HighlightFolderFieldList *string `uriparametername:"highlightFolderFieldList"`
    // An argument to 'body': Allows to return specific highlighted positions within the document. Valid choices are 'first', 'previous', 'next', 'last', 'firstUser', 'nextUser'. Default is to use no hit navigation. Note that 'first', 'last', and 'firstUser' ignore the page number and jump to the first resp. last hit. The choice 'next' jumps to the first page containing a hit which has a page number bigger than the 'page' parameter. If 'page' is the last page, it continues searching at the first page.
    HighlightHitNavigation *string `uriparametername:"highlightHitNavigation"`
    // This is a part of search-term based highlighting, see 'highlightSearchTermQuery'. It has the same syntax as 'joinRestriction' for searches.
    HighlightSearchTermJoinRestriction *string `uriparametername:"highlightSearchTermJoinRestriction"`
    // Specify a language to be used when interpreting the query. The default is to use the result of the query language detection. The argument is a two-letter ISO 639 language code like 'en'. The language is used in order to define stemming rules, for example.
    HighlightSearchTermLanguage *string `uriparametername:"highlightSearchTermLanguage"`
    // A query expression which should be used to highlight hits in the requested document. The query expression resembles the value of 'query' for search operations, but it is only used for search term highlighting and concept highlighting. Note that the record id in question does not necessarily need to match the query in question.
    HighlightSearchTermQuery *string `uriparametername:"highlightSearchTermQuery"`
    // An argument to 'body': A comma-separated list of terms to highlight.
    HighlightUserTerms *string `uriparametername:"highlightUserTerms"`
    // Define which page to retrieve (counting starts at 1). Default is 1. A page number which is too large will automatically be replaced by the highest page number.
    Page *int32 `uriparametername:"page"`
    // An argument to 'body': indicates whether to provide a textual summary for the record, default is false in this context. The summary shows text portions around matching highlighted words in the document.
    Summarize *bool `uriparametername:"summarize"`
}
// ItemCollectionsItemRecordsItemContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsItemContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemRecordsItemContentRequestBuilderGetQueryParameters
}
// ItemCollectionsItemRecordsItemContentRequestBuilderPostQueryParameters returns selected properties (possibly including paged body content) for the record in question.
type ItemCollectionsItemRecordsItemContentRequestBuilderPostQueryParameters struct {
    // Indicate whether to retrieve content for the record, default is false. This parameter is required in order to display highlighting results. It fetches fields configured as 'Content Fields' in the data model, applies highlighting, paging, and summarization as needed, and returns the resulting XML.
    Body *bool `uriparametername:"body"`
    // An URI-encoded, comma-separated list of fields to be retrieved for each document. Valid field names are typically those which can be inspected using the fields endpoint (i.e. those configured in the data model). In addition, search results support to retrieve dynamic properties as field name: 'rm_is_best_bet' returns true if and only if 'Best bets boosting' is active and a match is a best bet.
    Fields *string `uriparametername:"fields"`
    // An URI-encoded, comma-separated list of fields which will be fetched in XML form for which highlighting has been applied. The resulting field content is in XML form with <recomDescriptiveWord> tags for highlighted words. The response object will contain one field for each item in this list, and the result items will be have the suffix 'Xml' to indicate that the highlighted value is in XML form (example 'titleXml'). Note that fields specified in this context do not necessarily need to be known in the data model. Note furthermore that highlighting will only match for indexed fields.
    FieldsHighlighted *string `uriparametername:"fieldsHighlighted"`
    // An URI-encoded, comma-separated list of folder fields (taxonomies) to be retrieved for each document. Note that folder fields can be returned by 'fields' as well in which case an array of folder ids is returned. Use 'folderFields' if you need the folder and its display name.
    FolderFields *string `uriparametername:"folderFields"`
    // An argument to 'body': A comma-separated list of folder field names for which folder-specific highlighting is to be applied.
    HighlightFolderFieldList *string `uriparametername:"highlightFolderFieldList"`
    // An argument to 'body': Allows to return specific highlighted positions within the document. Valid choices are 'first', 'previous', 'next', 'last', 'firstUser', 'nextUser'. Default is to use no hit navigation. Note that 'first', 'last', and 'firstUser' ignore the page number and jump to the first resp. last hit. The choice 'next' jumps to the first page containing a hit which has a page number bigger than the 'page' parameter. If 'page' is the last page, it continues searching at the first page.
    HighlightHitNavigation *string `uriparametername:"highlightHitNavigation"`
    // This is a part of search-term based highlighting, see 'highlightSearchTermQuery'. It has the same syntax as 'joinRestriction' for searches.
    HighlightSearchTermJoinRestriction *string `uriparametername:"highlightSearchTermJoinRestriction"`
    // Specify a language to be used when interpreting the query. The default is to use the result of the query language detection. The argument is a two-letter ISO 639 language code like 'en'. The language is used in order to define stemming rules, for example.
    HighlightSearchTermLanguage *string `uriparametername:"highlightSearchTermLanguage"`
    // A query expression which should be used to highlight hits in the requested document. The query expression resembles the value of 'query' for search operations, but it is only used for search term highlighting and concept highlighting. Note that the record id in question does not necessarily need to match the query in question.
    HighlightSearchTermQuery *string `uriparametername:"highlightSearchTermQuery"`
    // An argument to 'body': A comma-separated list of terms to highlight.
    HighlightUserTerms *string `uriparametername:"highlightUserTerms"`
    // Define which page to retrieve (counting starts at 1). Default is 1. A page number which is too large will automatically be replaced by the highest page number.
    Page *int32 `uriparametername:"page"`
    // An argument to 'body': indicates whether to provide a textual summary for the record, default is false in this context. The summary shows text portions around matching highlighted words in the document.
    Summarize *bool `uriparametername:"summarize"`
}
// ItemCollectionsItemRecordsItemContentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsItemContentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemRecordsItemContentRequestBuilderPostQueryParameters
}
// ItemCollectionsItemRecordsItemContentRequestBuilderPutQueryParameters allows to change properties of documents, for example by tagging them into folders or by modifying textual content.
type ItemCollectionsItemRecordsItemContentRequestBuilderPutQueryParameters struct {
    // Specifies whether the call should block until all indexes are updated. The default value 'false' is to return immediately which means that the display of values is always correct, but search indices may be outdated until all indices have been updated eventually (that means: queries will not respect the change until the index is updated). The change is persisted as soon as the API call returns, irrespective of 'blockUntilComplete'.
    BlockUntilComplete *bool `uriparametername:"blockUntilComplete"`
}
// ItemCollectionsItemRecordsItemContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemRecordsItemContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCollectionsItemRecordsItemContentRequestBuilderPutQueryParameters
}
// NewItemCollectionsItemRecordsItemContentRequestBuilderInternal instantiates a new ItemCollectionsItemRecordsItemContentRequestBuilder and sets the default values.
func NewItemCollectionsItemRecordsItemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemRecordsItemContentRequestBuilder) {
    m := &ItemCollectionsItemRecordsItemContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/records/{recordId}/content{?body*,fields*,fieldsHighlighted*,folderFields*,highlightFolderFieldList*,highlightHitNavigation*,highlightSearchTermJoinRestriction*,highlightSearchTermLanguage*,highlightSearchTermQuery*,highlightUserTerms*,page*,summarize*}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemRecordsItemContentRequestBuilder instantiates a new ItemCollectionsItemRecordsItemContentRequestBuilder and sets the default values.
func NewItemCollectionsItemRecordsItemContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemRecordsItemContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemRecordsItemContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns selected properties (possibly including paged body content) for the record in question.
// returns a Recordable when successful
func (m *ItemCollectionsItemRecordsItemContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsItemContentRequestBuilderGetRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.Recordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateRecordFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.Recordable), nil
}
// Post returns selected properties (possibly including paged body content) for the record in question.
// returns a Recordable when successful
func (m *ItemCollectionsItemRecordsItemContentRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsItemContentRequestBuilderPostRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.Recordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateRecordFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.Recordable), nil
}
// Put allows to change properties of documents, for example by tagging them into folders or by modifying textual content.
// returns a ChangeResultable when successful
func (m *ItemCollectionsItemRecordsItemContentRequestBuilder) Put(ctx context.Context, body []ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ChangeRequestable, requestConfiguration *ItemCollectionsItemRecordsItemContentRequestBuilderPutRequestConfiguration)(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ChangeResultable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.CreateChangeResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ChangeResultable), nil
}
// ToGetRequestInformation returns selected properties (possibly including paged body content) for the record in question.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsItemContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsItemContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation returns selected properties (possibly including paged body content) for the record in question.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsItemContentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemRecordsItemContentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation allows to change properties of documents, for example by tagging them into folders or by modifying textual content.
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemRecordsItemContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []ibec152b551cf495fa0e5e0ada537cd83bf312a9b36f6f1c37c65e6a69145c437.ChangeRequestable, requestConfiguration *ItemCollectionsItemRecordsItemContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/projects/{projectId}/collections/{collectionId}/records/{recordId}/content?blockUntilComplete={blockUntilComplete}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(body))
    for i, v := range body {
        if v != nil {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
    }
    err := requestInfo.SetContentFromParsableCollection(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", cast)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemRecordsItemContentRequestBuilder when successful
func (m *ItemCollectionsItemRecordsItemContentRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemRecordsItemContentRequestBuilder) {
    return NewItemCollectionsItemRecordsItemContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
