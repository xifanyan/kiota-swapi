package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type JobArgument struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The defaultValue property
    defaultValue *string
    // The description property
    description *string
    // The name property
    name *string
    // The type property
    typeEscaped *string
}
// NewJobArgument instantiates a new JobArgument and sets the default values.
func NewJobArgument()(*JobArgument) {
    m := &JobArgument{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobArgumentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateJobArgumentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobArgument(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *JobArgument) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultValue gets the defaultValue property value. The defaultValue property
// returns a *string when successful
func (m *JobArgument) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *JobArgument) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *JobArgument) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *JobArgument) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *JobArgument) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *JobArgument) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *JobArgument) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultValue sets the defaultValue property value. The defaultValue property
func (m *JobArgument) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetDescription sets the description property value. The description property
func (m *JobArgument) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name property
func (m *JobArgument) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *JobArgument) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type JobArgumentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*string)
    GetDescription()(*string)
    GetName()(*string)
    GetTypeEscaped()(*string)
    SetDefaultValue(value *string)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetTypeEscaped(value *string)()
}
