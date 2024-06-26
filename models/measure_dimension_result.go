package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MeasureDimensionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Documents with any value in this dimension. Negative if this value is unsupported.
    documentsWithAnyValue *int64
    // Documents without any value in this dimension. Negative if this value is unsupported.
    documentsWithNoValue *int64
    // the associated field name
    fieldName *string
    // The dimension members requested by the request.
    members []MeasureDimensionMemberable
    // The size of this dimension. For one-dimensional counts, this resembles all available results and can be larger than the elements in 'members' (due to paging). For two-dimensional counts, this value is the same as the size of 'members'.
    size *int64
}
// NewMeasureDimensionResult instantiates a new MeasureDimensionResult and sets the default values.
func NewMeasureDimensionResult()(*MeasureDimensionResult) {
    m := &MeasureDimensionResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMeasureDimensionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMeasureDimensionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeasureDimensionResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MeasureDimensionResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDocumentsWithAnyValue gets the documentsWithAnyValue property value. Documents with any value in this dimension. Negative if this value is unsupported.
// returns a *int64 when successful
func (m *MeasureDimensionResult) GetDocumentsWithAnyValue()(*int64) {
    return m.documentsWithAnyValue
}
// GetDocumentsWithNoValue gets the documentsWithNoValue property value. Documents without any value in this dimension. Negative if this value is unsupported.
// returns a *int64 when successful
func (m *MeasureDimensionResult) GetDocumentsWithNoValue()(*int64) {
    return m.documentsWithNoValue
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MeasureDimensionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["documentsWithAnyValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentsWithAnyValue(val)
        }
        return nil
    }
    res["documentsWithNoValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentsWithNoValue(val)
        }
        return nil
    }
    res["fieldName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFieldName(val)
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeasureDimensionMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeasureDimensionMemberable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MeasureDimensionMemberable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetFieldName gets the fieldName property value. the associated field name
// returns a *string when successful
func (m *MeasureDimensionResult) GetFieldName()(*string) {
    return m.fieldName
}
// GetMembers gets the members property value. The dimension members requested by the request.
// returns a []MeasureDimensionMemberable when successful
func (m *MeasureDimensionResult) GetMembers()([]MeasureDimensionMemberable) {
    return m.members
}
// GetSize gets the size property value. The size of this dimension. For one-dimensional counts, this resembles all available results and can be larger than the elements in 'members' (due to paging). For two-dimensional counts, this value is the same as the size of 'members'.
// returns a *int64 when successful
func (m *MeasureDimensionResult) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *MeasureDimensionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("documentsWithAnyValue", m.GetDocumentsWithAnyValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("documentsWithNoValue", m.GetDocumentsWithNoValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fieldName", m.GetFieldName())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("size", m.GetSize())
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
func (m *MeasureDimensionResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDocumentsWithAnyValue sets the documentsWithAnyValue property value. Documents with any value in this dimension. Negative if this value is unsupported.
func (m *MeasureDimensionResult) SetDocumentsWithAnyValue(value *int64)() {
    m.documentsWithAnyValue = value
}
// SetDocumentsWithNoValue sets the documentsWithNoValue property value. Documents without any value in this dimension. Negative if this value is unsupported.
func (m *MeasureDimensionResult) SetDocumentsWithNoValue(value *int64)() {
    m.documentsWithNoValue = value
}
// SetFieldName sets the fieldName property value. the associated field name
func (m *MeasureDimensionResult) SetFieldName(value *string)() {
    m.fieldName = value
}
// SetMembers sets the members property value. The dimension members requested by the request.
func (m *MeasureDimensionResult) SetMembers(value []MeasureDimensionMemberable)() {
    m.members = value
}
// SetSize sets the size property value. The size of this dimension. For one-dimensional counts, this resembles all available results and can be larger than the elements in 'members' (due to paging). For two-dimensional counts, this value is the same as the size of 'members'.
func (m *MeasureDimensionResult) SetSize(value *int64)() {
    m.size = value
}
type MeasureDimensionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDocumentsWithAnyValue()(*int64)
    GetDocumentsWithNoValue()(*int64)
    GetFieldName()(*string)
    GetMembers()([]MeasureDimensionMemberable)
    GetSize()(*int64)
    SetDocumentsWithAnyValue(value *int64)()
    SetDocumentsWithNoValue(value *int64)()
    SetFieldName(value *string)()
    SetMembers(value []MeasureDimensionMemberable)()
    SetSize(value *int64)()
}
