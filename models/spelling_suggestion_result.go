package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SpellingSuggestionResult a list of matching spelling suggestions (if configured for the target project). The return value 'null' means that such suggestions have not been requested.
type SpellingSuggestionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The suggestedWords property
    suggestedWords []SpellingSuggestionable
    // The total improvement in percent if all suggestions are applied (between 0 and 1). Note that this number is unrelated to the individual improvements of suggested words as it also takes those words into account for which no suggestion has been returned.
    totalImprovementPercent *float32
}
// NewSpellingSuggestionResult instantiates a new SpellingSuggestionResult and sets the default values.
func NewSpellingSuggestionResult()(*SpellingSuggestionResult) {
    m := &SpellingSuggestionResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSpellingSuggestionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSpellingSuggestionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSpellingSuggestionResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SpellingSuggestionResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SpellingSuggestionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["suggestedWords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSpellingSuggestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SpellingSuggestionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SpellingSuggestionable)
                }
            }
            m.SetSuggestedWords(res)
        }
        return nil
    }
    res["totalImprovementPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalImprovementPercent(val)
        }
        return nil
    }
    return res
}
// GetSuggestedWords gets the suggestedWords property value. The suggestedWords property
// returns a []SpellingSuggestionable when successful
func (m *SpellingSuggestionResult) GetSuggestedWords()([]SpellingSuggestionable) {
    return m.suggestedWords
}
// GetTotalImprovementPercent gets the totalImprovementPercent property value. The total improvement in percent if all suggestions are applied (between 0 and 1). Note that this number is unrelated to the individual improvements of suggested words as it also takes those words into account for which no suggestion has been returned.
// returns a *float32 when successful
func (m *SpellingSuggestionResult) GetTotalImprovementPercent()(*float32) {
    return m.totalImprovementPercent
}
// Serialize serializes information the current object
func (m *SpellingSuggestionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSuggestedWords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSuggestedWords()))
        for i, v := range m.GetSuggestedWords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("suggestedWords", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("totalImprovementPercent", m.GetTotalImprovementPercent())
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
func (m *SpellingSuggestionResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSuggestedWords sets the suggestedWords property value. The suggestedWords property
func (m *SpellingSuggestionResult) SetSuggestedWords(value []SpellingSuggestionable)() {
    m.suggestedWords = value
}
// SetTotalImprovementPercent sets the totalImprovementPercent property value. The total improvement in percent if all suggestions are applied (between 0 and 1). Note that this number is unrelated to the individual improvements of suggested words as it also takes those words into account for which no suggestion has been returned.
func (m *SpellingSuggestionResult) SetTotalImprovementPercent(value *float32)() {
    m.totalImprovementPercent = value
}
type SpellingSuggestionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSuggestedWords()([]SpellingSuggestionable)
    GetTotalImprovementPercent()(*float32)
    SetSuggestedWords(value []SpellingSuggestionable)()
    SetTotalImprovementPercent(value *float32)()
}
