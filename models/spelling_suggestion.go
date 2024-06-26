package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SpellingSuggestion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The improvement in percent achieved by applying the suggestion (between 0 and 1), but limited only to this word.
    improvementPercent *float32
    // Will be true if and only if the originalWord was found in some record.
    isOriginalFound *bool
    // The word is it was found in the original search request.
    originalWord *string
    // The suggested alternative word.
    suggestedWord *string
}
// NewSpellingSuggestion instantiates a new SpellingSuggestion and sets the default values.
func NewSpellingSuggestion()(*SpellingSuggestion) {
    m := &SpellingSuggestion{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSpellingSuggestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSpellingSuggestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSpellingSuggestion(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SpellingSuggestion) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SpellingSuggestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["improvementPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImprovementPercent(val)
        }
        return nil
    }
    res["isOriginalFound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOriginalFound(val)
        }
        return nil
    }
    res["originalWord"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalWord(val)
        }
        return nil
    }
    res["suggestedWord"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestedWord(val)
        }
        return nil
    }
    return res
}
// GetImprovementPercent gets the improvementPercent property value. The improvement in percent achieved by applying the suggestion (between 0 and 1), but limited only to this word.
// returns a *float32 when successful
func (m *SpellingSuggestion) GetImprovementPercent()(*float32) {
    return m.improvementPercent
}
// GetIsOriginalFound gets the isOriginalFound property value. Will be true if and only if the originalWord was found in some record.
// returns a *bool when successful
func (m *SpellingSuggestion) GetIsOriginalFound()(*bool) {
    return m.isOriginalFound
}
// GetOriginalWord gets the originalWord property value. The word is it was found in the original search request.
// returns a *string when successful
func (m *SpellingSuggestion) GetOriginalWord()(*string) {
    return m.originalWord
}
// GetSuggestedWord gets the suggestedWord property value. The suggested alternative word.
// returns a *string when successful
func (m *SpellingSuggestion) GetSuggestedWord()(*string) {
    return m.suggestedWord
}
// Serialize serializes information the current object
func (m *SpellingSuggestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat32Value("improvementPercent", m.GetImprovementPercent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOriginalFound", m.GetIsOriginalFound())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("originalWord", m.GetOriginalWord())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("suggestedWord", m.GetSuggestedWord())
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
func (m *SpellingSuggestion) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImprovementPercent sets the improvementPercent property value. The improvement in percent achieved by applying the suggestion (between 0 and 1), but limited only to this word.
func (m *SpellingSuggestion) SetImprovementPercent(value *float32)() {
    m.improvementPercent = value
}
// SetIsOriginalFound sets the isOriginalFound property value. Will be true if and only if the originalWord was found in some record.
func (m *SpellingSuggestion) SetIsOriginalFound(value *bool)() {
    m.isOriginalFound = value
}
// SetOriginalWord sets the originalWord property value. The word is it was found in the original search request.
func (m *SpellingSuggestion) SetOriginalWord(value *string)() {
    m.originalWord = value
}
// SetSuggestedWord sets the suggestedWord property value. The suggested alternative word.
func (m *SpellingSuggestion) SetSuggestedWord(value *string)() {
    m.suggestedWord = value
}
type SpellingSuggestionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetImprovementPercent()(*float32)
    GetIsOriginalFound()(*bool)
    GetOriginalWord()(*string)
    GetSuggestedWord()(*string)
    SetImprovementPercent(value *float32)()
    SetIsOriginalFound(value *bool)()
    SetOriginalWord(value *string)()
    SetSuggestedWord(value *string)()
}
