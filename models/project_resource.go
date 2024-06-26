package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProjectResource the list of project resources.
type ProjectResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A description of the resource.
    description *string
    // The ID of the resource.
    id *ProjectResource_id
}
// NewProjectResource instantiates a new ProjectResource and sets the default values.
func NewProjectResource()(*ProjectResource) {
    m := &ProjectResource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProjectResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProjectResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProjectResource(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProjectResource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. A description of the resource.
// returns a *string when successful
func (m *ProjectResource) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProjectResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProjectResource_id)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val.(*ProjectResource_id))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the resource.
// returns a *ProjectResource_id when successful
func (m *ProjectResource) GetId()(*ProjectResource_id) {
    return m.id
}
// Serialize serializes information the current object
func (m *ProjectResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetId() != nil {
        cast := (*m.GetId()).String()
        err := writer.WriteStringValue("id", &cast)
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
func (m *ProjectResource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. A description of the resource.
func (m *ProjectResource) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The ID of the resource.
func (m *ProjectResource) SetId(value *ProjectResource_id)() {
    m.id = value
}
type ProjectResourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetId()(*ProjectResource_id)
    SetDescription(value *string)()
    SetId(value *ProjectResource_id)()
}
