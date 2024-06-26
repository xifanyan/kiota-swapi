package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CachedSearchDescription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The trace is which was used when creating the search. It starts with the value of the request header SWA-MDC-TOKEN which was used when creating the search. If the search was created without this header, it is auto-generated.
    creationTraceId *string
}
// NewCachedSearchDescription instantiates a new CachedSearchDescription and sets the default values.
func NewCachedSearchDescription()(*CachedSearchDescription) {
    m := &CachedSearchDescription{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCachedSearchDescriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCachedSearchDescriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCachedSearchDescription(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CachedSearchDescription) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreationTraceId gets the creationTraceId property value. The trace is which was used when creating the search. It starts with the value of the request header SWA-MDC-TOKEN which was used when creating the search. If the search was created without this header, it is auto-generated.
// returns a *string when successful
func (m *CachedSearchDescription) GetCreationTraceId()(*string) {
    return m.creationTraceId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CachedSearchDescription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["creationTraceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationTraceId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CachedSearchDescription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("creationTraceId", m.GetCreationTraceId())
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
func (m *CachedSearchDescription) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreationTraceId sets the creationTraceId property value. The trace is which was used when creating the search. It starts with the value of the request header SWA-MDC-TOKEN which was used when creating the search. If the search was created without this header, it is auto-generated.
func (m *CachedSearchDescription) SetCreationTraceId(value *string)() {
    m.creationTraceId = value
}
type CachedSearchDescriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreationTraceId()(*string)
    SetCreationTraceId(value *string)()
}
