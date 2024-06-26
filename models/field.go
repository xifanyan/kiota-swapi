package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Field a list of field-value pairs, if any requested.
type Field struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The field ID.
    id *string
    // The value for the given document in the given field or null if there is no such value. Date values are formatted as string according to ISO_8601, i.e. in the date format 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX' in UTC timezone. Folder ids are formatted as comma-separated list without any escape characters (see valueObject for a type-safe access).
    value *string
    // This is either null or redundant: it provides the very same data as in 'value', but as typed object. This property is populated if and only if the result is not a plain string. Examples where this property is populated and contains useful values are: integer fields, date fields (which return the number of milliseconds since January 1, 1970, 00:00:00 GMT, formatted as string) or an array of folder values (if they have been requested as 'fields'). Note that 'value' will always be populated with a string-representation of the value. Use 'valueObject' in order to have type-safe access (especially if you retrieved a list of folder ids).
    valueObject Field_valueObjectable
}
// NewField instantiates a new Field and sets the default values.
func NewField()(*Field) {
    m := &Field{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFieldFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFieldFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewField(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Field) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Field) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["valueObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateField_valueObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueObject(val.(Field_valueObjectable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The field ID.
// returns a *string when successful
func (m *Field) GetId()(*string) {
    return m.id
}
// GetValue gets the value property value. The value for the given document in the given field or null if there is no such value. Date values are formatted as string according to ISO_8601, i.e. in the date format 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX' in UTC timezone. Folder ids are formatted as comma-separated list without any escape characters (see valueObject for a type-safe access).
// returns a *string when successful
func (m *Field) GetValue()(*string) {
    return m.value
}
// GetValueObject gets the valueObject property value. This is either null or redundant: it provides the very same data as in 'value', but as typed object. This property is populated if and only if the result is not a plain string. Examples where this property is populated and contains useful values are: integer fields, date fields (which return the number of milliseconds since January 1, 1970, 00:00:00 GMT, formatted as string) or an array of folder values (if they have been requested as 'fields'). Note that 'value' will always be populated with a string-representation of the value. Use 'valueObject' in order to have type-safe access (especially if you retrieved a list of folder ids).
// returns a Field_valueObjectable when successful
func (m *Field) GetValueObject()(Field_valueObjectable) {
    return m.valueObject
}
// Serialize serializes information the current object
func (m *Field) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
    {
        err := writer.WriteObjectValue("valueObject", m.GetValueObject())
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
func (m *Field) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The field ID.
func (m *Field) SetId(value *string)() {
    m.id = value
}
// SetValue sets the value property value. The value for the given document in the given field or null if there is no such value. Date values are formatted as string according to ISO_8601, i.e. in the date format 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX' in UTC timezone. Folder ids are formatted as comma-separated list without any escape characters (see valueObject for a type-safe access).
func (m *Field) SetValue(value *string)() {
    m.value = value
}
// SetValueObject sets the valueObject property value. This is either null or redundant: it provides the very same data as in 'value', but as typed object. This property is populated if and only if the result is not a plain string. Examples where this property is populated and contains useful values are: integer fields, date fields (which return the number of milliseconds since January 1, 1970, 00:00:00 GMT, formatted as string) or an array of folder values (if they have been requested as 'fields'). Note that 'value' will always be populated with a string-representation of the value. Use 'valueObject' in order to have type-safe access (especially if you retrieved a list of folder ids).
func (m *Field) SetValueObject(value Field_valueObjectable)() {
    m.valueObject = value
}
type Fieldable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetValue()(*string)
    GetValueObject()(Field_valueObjectable)
    SetId(value *string)()
    SetValue(value *string)()
    SetValueObject(value Field_valueObjectable)()
}
