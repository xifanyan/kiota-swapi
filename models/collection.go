package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Collection the list of collections.
type Collection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The associated display name. Can be the same as id.
    displayName *string
    // The sub collection ID, as defined in the meta engine configuration in 'Meta engine sources: Server definition'. The ID resembles the configuration column 'Identifier' and is to be used where the API requires a 'collectionId'. Use 'default' for single engines and merging meta engines.
    id *string
}
// NewCollection instantiates a new Collection and sets the default values.
func NewCollection()(*Collection) {
    m := &Collection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollection(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Collection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The associated display name. Can be the same as id.
// returns a *string when successful
func (m *Collection) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Collection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
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
    return res
}
// GetId gets the id property value. The sub collection ID, as defined in the meta engine configuration in 'Meta engine sources: Server definition'. The ID resembles the configuration column 'Identifier' and is to be used where the API requires a 'collectionId'. Use 'default' for single engines and merging meta engines.
// returns a *string when successful
func (m *Collection) GetId()(*string) {
    return m.id
}
// Serialize serializes information the current object
func (m *Collection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *Collection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The associated display name. Can be the same as id.
func (m *Collection) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The sub collection ID, as defined in the meta engine configuration in 'Meta engine sources: Server definition'. The ID resembles the configuration column 'Identifier' and is to be used where the API requires a 'collectionId'. Use 'default' for single engines and merging meta engines.
func (m *Collection) SetId(value *string)() {
    m.id = value
}
type Collectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    SetDisplayName(value *string)()
    SetId(value *string)()
}
