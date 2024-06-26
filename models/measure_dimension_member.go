package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeasureDimensionMember the dimension members requested by the request.
type MeasureDimensionMember struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A list of field-value pairs, if any requested.
    fields []Fieldable
    // The member's identifier.
    identifier *string
}
// NewMeasureDimensionMember instantiates a new MeasureDimensionMember and sets the default values.
func NewMeasureDimensionMember()(*MeasureDimensionMember) {
    m := &MeasureDimensionMember{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMeasureDimensionMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMeasureDimensionMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeasureDimensionMember(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MeasureDimensionMember) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MeasureDimensionMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFieldFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fieldable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Fieldable)
                }
            }
            m.SetFields(res)
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. A list of field-value pairs, if any requested.
// returns a []Fieldable when successful
func (m *MeasureDimensionMember) GetFields()([]Fieldable) {
    return m.fields
}
// GetIdentifier gets the identifier property value. The member's identifier.
// returns a *string when successful
func (m *MeasureDimensionMember) GetIdentifier()(*string) {
    return m.identifier
}
// Serialize serializes information the current object
func (m *MeasureDimensionMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFields() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFields()))
        for i, v := range m.GetFields() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("fields", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
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
func (m *MeasureDimensionMember) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFields sets the fields property value. A list of field-value pairs, if any requested.
func (m *MeasureDimensionMember) SetFields(value []Fieldable)() {
    m.fields = value
}
// SetIdentifier sets the identifier property value. The member's identifier.
func (m *MeasureDimensionMember) SetIdentifier(value *string)() {
    m.identifier = value
}
type MeasureDimensionMemberable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFields()([]Fieldable)
    GetIdentifier()(*string)
    SetFields(value []Fieldable)()
    SetIdentifier(value *string)()
}
