package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecordData 0-n records to be created.
type RecordData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The fieldData property
    fieldData []FieldDataable
    // A unique identifier for the record. It resembles the value of 'uniqueField' when fetching records and must be provided at insertion time.
    uniqueId *string
}
// NewRecordData instantiates a new RecordData and sets the default values.
func NewRecordData()(*RecordData) {
    m := &RecordData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecordDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecordDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecordData(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RecordData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldData gets the fieldData property value. The fieldData property
// returns a []FieldDataable when successful
func (m *RecordData) GetFieldData()([]FieldDataable) {
    return m.fieldData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecordData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fieldData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFieldDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FieldDataable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FieldDataable)
                }
            }
            m.SetFieldData(res)
        }
        return nil
    }
    res["uniqueId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueId(val)
        }
        return nil
    }
    return res
}
// GetUniqueId gets the uniqueId property value. A unique identifier for the record. It resembles the value of 'uniqueField' when fetching records and must be provided at insertion time.
// returns a *string when successful
func (m *RecordData) GetUniqueId()(*string) {
    return m.uniqueId
}
// Serialize serializes information the current object
func (m *RecordData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFieldData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFieldData()))
        for i, v := range m.GetFieldData() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("fieldData", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uniqueId", m.GetUniqueId())
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
func (m *RecordData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFieldData sets the fieldData property value. The fieldData property
func (m *RecordData) SetFieldData(value []FieldDataable)() {
    m.fieldData = value
}
// SetUniqueId sets the uniqueId property value. A unique identifier for the record. It resembles the value of 'uniqueField' when fetching records and must be provided at insertion time.
func (m *RecordData) SetUniqueId(value *string)() {
    m.uniqueId = value
}
type RecordDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFieldData()([]FieldDataable)
    GetUniqueId()(*string)
    SetFieldData(value []FieldDataable)()
    SetUniqueId(value *string)()
}
