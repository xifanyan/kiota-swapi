package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FolderFieldResource the list of folder field (taxonomy, smart    filter) resources.
type FolderFieldResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This provides a description of available query parameters ('...&<name>=<value>&...') for the resource.
    availableQueryParameters []QueryParameterable
    // A description of the resource.
    description *string
    // The ID of the resource.
    id *FolderFieldResource_id
}
// NewFolderFieldResource instantiates a new FolderFieldResource and sets the default values.
func NewFolderFieldResource()(*FolderFieldResource) {
    m := &FolderFieldResource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFolderFieldResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFolderFieldResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFolderFieldResource(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FolderFieldResource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableQueryParameters gets the availableQueryParameters property value. This provides a description of available query parameters ('...&<name>=<value>&...') for the resource.
// returns a []QueryParameterable when successful
func (m *FolderFieldResource) GetAvailableQueryParameters()([]QueryParameterable) {
    return m.availableQueryParameters
}
// GetDescription gets the description property value. A description of the resource.
// returns a *string when successful
func (m *FolderFieldResource) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FolderFieldResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableQueryParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateQueryParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]QueryParameterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(QueryParameterable)
                }
            }
            m.SetAvailableQueryParameters(res)
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFolderFieldResource_id)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val.(*FolderFieldResource_id))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the resource.
// returns a *FolderFieldResource_id when successful
func (m *FolderFieldResource) GetId()(*FolderFieldResource_id) {
    return m.id
}
// Serialize serializes information the current object
func (m *FolderFieldResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAvailableQueryParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAvailableQueryParameters()))
        for i, v := range m.GetAvailableQueryParameters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("availableQueryParameters", cast)
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
func (m *FolderFieldResource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableQueryParameters sets the availableQueryParameters property value. This provides a description of available query parameters ('...&<name>=<value>&...') for the resource.
func (m *FolderFieldResource) SetAvailableQueryParameters(value []QueryParameterable)() {
    m.availableQueryParameters = value
}
// SetDescription sets the description property value. A description of the resource.
func (m *FolderFieldResource) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The ID of the resource.
func (m *FolderFieldResource) SetId(value *FolderFieldResource_id)() {
    m.id = value
}
type FolderFieldResourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableQueryParameters()([]QueryParameterable)
    GetDescription()(*string)
    GetId()(*FolderFieldResource_id)
    SetAvailableQueryParameters(value []QueryParameterable)()
    SetDescription(value *string)()
    SetId(value *FolderFieldResource_id)()
}
