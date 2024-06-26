package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WaitForPendingChangesResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Returns true if all currently scheduled transactions finished within the timeout.
    success *bool
}
// NewWaitForPendingChangesResult instantiates a new WaitForPendingChangesResult and sets the default values.
func NewWaitForPendingChangesResult()(*WaitForPendingChangesResult) {
    m := &WaitForPendingChangesResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWaitForPendingChangesResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWaitForPendingChangesResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWaitForPendingChangesResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WaitForPendingChangesResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WaitForPendingChangesResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["success"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccess(val)
        }
        return nil
    }
    return res
}
// GetSuccess gets the success property value. Returns true if all currently scheduled transactions finished within the timeout.
// returns a *bool when successful
func (m *WaitForPendingChangesResult) GetSuccess()(*bool) {
    return m.success
}
// Serialize serializes information the current object
func (m *WaitForPendingChangesResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("success", m.GetSuccess())
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
func (m *WaitForPendingChangesResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSuccess sets the success property value. Returns true if all currently scheduled transactions finished within the timeout.
func (m *WaitForPendingChangesResult) SetSuccess(value *bool)() {
    m.success = value
}
type WaitForPendingChangesResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSuccess()(*bool)
    SetSuccess(value *bool)()
}
