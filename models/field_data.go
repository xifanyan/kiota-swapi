package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FieldData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The field name. A common use-case is to create new folders. In this case, rm_display_name is the field for the folder's display name. The field rm_prop_parent is the field defining the new record's parent (it resembles the parent's unique id).
    fieldName *string
    // The value to assign to this field. Only one of 'value' or 'valueList' can be populated. If this field data is of type binary, the 'value' must be the index of the binary artifact, i.e. the index within the multipart upload form.
    value *string
    // A list of values to assign to this field (for multi-value fields only). Only one of 'value' or 'valueList' can be populated.
    valueList []string
}
// NewFieldData instantiates a new FieldData and sets the default values.
func NewFieldData()(*FieldData) {
    m := &FieldData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFieldDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFieldDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFieldData(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FieldData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FieldData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["valueList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetValueList(res)
        }
        return nil
    }
    return res
}
// GetFieldName gets the fieldName property value. The field name. A common use-case is to create new folders. In this case, rm_display_name is the field for the folder's display name. The field rm_prop_parent is the field defining the new record's parent (it resembles the parent's unique id).
// returns a *string when successful
func (m *FieldData) GetFieldName()(*string) {
    return m.fieldName
}
// GetValue gets the value property value. The value to assign to this field. Only one of 'value' or 'valueList' can be populated. If this field data is of type binary, the 'value' must be the index of the binary artifact, i.e. the index within the multipart upload form.
// returns a *string when successful
func (m *FieldData) GetValue()(*string) {
    return m.value
}
// GetValueList gets the valueList property value. A list of values to assign to this field (for multi-value fields only). Only one of 'value' or 'valueList' can be populated.
// returns a []string when successful
func (m *FieldData) GetValueList()([]string) {
    return m.valueList
}
// Serialize serializes information the current object
func (m *FieldData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fieldName", m.GetFieldName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    if m.GetValueList() != nil {
        err := writer.WriteCollectionOfStringValues("valueList", m.GetValueList())
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
func (m *FieldData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFieldName sets the fieldName property value. The field name. A common use-case is to create new folders. In this case, rm_display_name is the field for the folder's display name. The field rm_prop_parent is the field defining the new record's parent (it resembles the parent's unique id).
func (m *FieldData) SetFieldName(value *string)() {
    m.fieldName = value
}
// SetValue sets the value property value. The value to assign to this field. Only one of 'value' or 'valueList' can be populated. If this field data is of type binary, the 'value' must be the index of the binary artifact, i.e. the index within the multipart upload form.
func (m *FieldData) SetValue(value *string)() {
    m.value = value
}
// SetValueList sets the valueList property value. A list of values to assign to this field (for multi-value fields only). Only one of 'value' or 'valueList' can be populated.
func (m *FieldData) SetValueList(value []string)() {
    m.valueList = value
}
type FieldDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFieldName()(*string)
    GetValue()(*string)
    GetValueList()([]string)
    SetFieldName(value *string)()
    SetValue(value *string)()
    SetValueList(value []string)()
}
