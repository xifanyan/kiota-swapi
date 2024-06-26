package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SearchResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of results found by the search. With paging this may be more than available in the result.
    numberResults *int64
    // The list of records.
    results []Recordable
    // A list of matching spelling suggestions (if configured for the target project). The return value 'null' means that such suggestions have not been requested.
    spellingSuggestions SpellingSuggestionResultable
    // A list of matching sponsored links (if configured for the target project).
    sponsoredLinks []SponsoredLinkable
    // The status of the response.
    status StatusObjectable
}
// NewSearchResult instantiates a new SearchResult and sets the default values.
func NewSearchResult()(*SearchResult) {
    m := &SearchResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSearchResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SearchResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SearchResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["numberResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberResults(val)
        }
        return nil
    }
    res["results"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Recordable)
                }
            }
            m.SetResults(res)
        }
        return nil
    }
    res["spellingSuggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSpellingSuggestionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpellingSuggestions(val.(SpellingSuggestionResultable))
        }
        return nil
    }
    res["sponsoredLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSponsoredLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SponsoredLinkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SponsoredLinkable)
                }
            }
            m.SetSponsoredLinks(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatusObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(StatusObjectable))
        }
        return nil
    }
    return res
}
// GetNumberResults gets the numberResults property value. The number of results found by the search. With paging this may be more than available in the result.
// returns a *int64 when successful
func (m *SearchResult) GetNumberResults()(*int64) {
    return m.numberResults
}
// GetResults gets the results property value. The list of records.
// returns a []Recordable when successful
func (m *SearchResult) GetResults()([]Recordable) {
    return m.results
}
// GetSpellingSuggestions gets the spellingSuggestions property value. A list of matching spelling suggestions (if configured for the target project). The return value 'null' means that such suggestions have not been requested.
// returns a SpellingSuggestionResultable when successful
func (m *SearchResult) GetSpellingSuggestions()(SpellingSuggestionResultable) {
    return m.spellingSuggestions
}
// GetSponsoredLinks gets the sponsoredLinks property value. A list of matching sponsored links (if configured for the target project).
// returns a []SponsoredLinkable when successful
func (m *SearchResult) GetSponsoredLinks()([]SponsoredLinkable) {
    return m.sponsoredLinks
}
// GetStatus gets the status property value. The status of the response.
// returns a StatusObjectable when successful
func (m *SearchResult) GetStatus()(StatusObjectable) {
    return m.status
}
// Serialize serializes information the current object
func (m *SearchResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("numberResults", m.GetNumberResults())
        if err != nil {
            return err
        }
    }
    if m.GetResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResults()))
        for i, v := range m.GetResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("results", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("spellingSuggestions", m.GetSpellingSuggestions())
        if err != nil {
            return err
        }
    }
    if m.GetSponsoredLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSponsoredLinks()))
        for i, v := range m.GetSponsoredLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("sponsoredLinks", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumberResults sets the numberResults property value. The number of results found by the search. With paging this may be more than available in the result.
func (m *SearchResult) SetNumberResults(value *int64)() {
    m.numberResults = value
}
// SetResults sets the results property value. The list of records.
func (m *SearchResult) SetResults(value []Recordable)() {
    m.results = value
}
// SetSpellingSuggestions sets the spellingSuggestions property value. A list of matching spelling suggestions (if configured for the target project). The return value 'null' means that such suggestions have not been requested.
func (m *SearchResult) SetSpellingSuggestions(value SpellingSuggestionResultable)() {
    m.spellingSuggestions = value
}
// SetSponsoredLinks sets the sponsoredLinks property value. A list of matching sponsored links (if configured for the target project).
func (m *SearchResult) SetSponsoredLinks(value []SponsoredLinkable)() {
    m.sponsoredLinks = value
}
// SetStatus sets the status property value. The status of the response.
func (m *SearchResult) SetStatus(value StatusObjectable)() {
    m.status = value
}
type SearchResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNumberResults()(*int64)
    GetResults()([]Recordable)
    GetSpellingSuggestions()(SpellingSuggestionResultable)
    GetSponsoredLinks()([]SponsoredLinkable)
    GetStatus()(StatusObjectable)
    SetNumberResults(value *int64)()
    SetResults(value []Recordable)()
    SetSpellingSuggestions(value SpellingSuggestionResultable)()
    SetSponsoredLinks(value []SponsoredLinkable)()
    SetStatus(value StatusObjectable)()
}
