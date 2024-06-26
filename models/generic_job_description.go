package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenericJobDescription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The arguments property
    arguments []JobArgumentable
    // The description property
    description *string
    // The registeredName property
    registeredName *string
}
// NewGenericJobDescription instantiates a new GenericJobDescription and sets the default values.
func NewGenericJobDescription()(*GenericJobDescription) {
    m := &GenericJobDescription{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGenericJobDescriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenericJobDescriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGenericJobDescription(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GenericJobDescription) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArguments gets the arguments property value. The arguments property
// returns a []JobArgumentable when successful
func (m *GenericJobDescription) GetArguments()([]JobArgumentable) {
    return m.arguments
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *GenericJobDescription) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenericJobDescription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["arguments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobArgumentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobArgumentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JobArgumentable)
                }
            }
            m.SetArguments(res)
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
    res["registeredName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredName(val)
        }
        return nil
    }
    return res
}
// GetRegisteredName gets the registeredName property value. The registeredName property
// returns a *string when successful
func (m *GenericJobDescription) GetRegisteredName()(*string) {
    return m.registeredName
}
// Serialize serializes information the current object
func (m *GenericJobDescription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetArguments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetArguments()))
        for i, v := range m.GetArguments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("arguments", cast)
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
        err := writer.WriteStringValue("registeredName", m.GetRegisteredName())
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
func (m *GenericJobDescription) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArguments sets the arguments property value. The arguments property
func (m *GenericJobDescription) SetArguments(value []JobArgumentable)() {
    m.arguments = value
}
// SetDescription sets the description property value. The description property
func (m *GenericJobDescription) SetDescription(value *string)() {
    m.description = value
}
// SetRegisteredName sets the registeredName property value. The registeredName property
func (m *GenericJobDescription) SetRegisteredName(value *string)() {
    m.registeredName = value
}
type GenericJobDescriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArguments()([]JobArgumentable)
    GetDescription()(*string)
    GetRegisteredName()(*string)
    SetArguments(value []JobArgumentable)()
    SetDescription(value *string)()
    SetRegisteredName(value *string)()
}
