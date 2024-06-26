package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type JobPhaseDescription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The completionFractionOfWork property
    completionFractionOfWork *float32
    // The currentErrorCount property
    currentErrorCount *int64
    // The description property
    description *string
    // The expectedStepCount property
    expectedStepCount *int64
    // The finishDate property
    finishDate *string
    // The finishedSteps property
    finishedSteps *int64
    // The name property
    name *string
    // The skippedSteps property
    skippedSteps *int64
    // The startDate property
    startDate *string
}
// NewJobPhaseDescription instantiates a new JobPhaseDescription and sets the default values.
func NewJobPhaseDescription()(*JobPhaseDescription) {
    m := &JobPhaseDescription{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobPhaseDescriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateJobPhaseDescriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobPhaseDescription(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *JobPhaseDescription) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompletionFractionOfWork gets the completionFractionOfWork property value. The completionFractionOfWork property
// returns a *float32 when successful
func (m *JobPhaseDescription) GetCompletionFractionOfWork()(*float32) {
    return m.completionFractionOfWork
}
// GetCurrentErrorCount gets the currentErrorCount property value. The currentErrorCount property
// returns a *int64 when successful
func (m *JobPhaseDescription) GetCurrentErrorCount()(*int64) {
    return m.currentErrorCount
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *JobPhaseDescription) GetDescription()(*string) {
    return m.description
}
// GetExpectedStepCount gets the expectedStepCount property value. The expectedStepCount property
// returns a *int64 when successful
func (m *JobPhaseDescription) GetExpectedStepCount()(*int64) {
    return m.expectedStepCount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *JobPhaseDescription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completionFractionOfWork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionFractionOfWork(val)
        }
        return nil
    }
    res["currentErrorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentErrorCount(val)
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
    res["expectedStepCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedStepCount(val)
        }
        return nil
    }
    res["finishDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinishDate(val)
        }
        return nil
    }
    res["finishedSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinishedSteps(val)
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
    res["skippedSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkippedSteps(val)
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    return res
}
// GetFinishDate gets the finishDate property value. The finishDate property
// returns a *string when successful
func (m *JobPhaseDescription) GetFinishDate()(*string) {
    return m.finishDate
}
// GetFinishedSteps gets the finishedSteps property value. The finishedSteps property
// returns a *int64 when successful
func (m *JobPhaseDescription) GetFinishedSteps()(*int64) {
    return m.finishedSteps
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *JobPhaseDescription) GetName()(*string) {
    return m.name
}
// GetSkippedSteps gets the skippedSteps property value. The skippedSteps property
// returns a *int64 when successful
func (m *JobPhaseDescription) GetSkippedSteps()(*int64) {
    return m.skippedSteps
}
// GetStartDate gets the startDate property value. The startDate property
// returns a *string when successful
func (m *JobPhaseDescription) GetStartDate()(*string) {
    return m.startDate
}
// Serialize serializes information the current object
func (m *JobPhaseDescription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat32Value("completionFractionOfWork", m.GetCompletionFractionOfWork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("currentErrorCount", m.GetCurrentErrorCount())
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
        err := writer.WriteInt64Value("expectedStepCount", m.GetExpectedStepCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("finishDate", m.GetFinishDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("finishedSteps", m.GetFinishedSteps())
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
        err := writer.WriteInt64Value("skippedSteps", m.GetSkippedSteps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("startDate", m.GetStartDate())
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
func (m *JobPhaseDescription) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompletionFractionOfWork sets the completionFractionOfWork property value. The completionFractionOfWork property
func (m *JobPhaseDescription) SetCompletionFractionOfWork(value *float32)() {
    m.completionFractionOfWork = value
}
// SetCurrentErrorCount sets the currentErrorCount property value. The currentErrorCount property
func (m *JobPhaseDescription) SetCurrentErrorCount(value *int64)() {
    m.currentErrorCount = value
}
// SetDescription sets the description property value. The description property
func (m *JobPhaseDescription) SetDescription(value *string)() {
    m.description = value
}
// SetExpectedStepCount sets the expectedStepCount property value. The expectedStepCount property
func (m *JobPhaseDescription) SetExpectedStepCount(value *int64)() {
    m.expectedStepCount = value
}
// SetFinishDate sets the finishDate property value. The finishDate property
func (m *JobPhaseDescription) SetFinishDate(value *string)() {
    m.finishDate = value
}
// SetFinishedSteps sets the finishedSteps property value. The finishedSteps property
func (m *JobPhaseDescription) SetFinishedSteps(value *int64)() {
    m.finishedSteps = value
}
// SetName sets the name property value. The name property
func (m *JobPhaseDescription) SetName(value *string)() {
    m.name = value
}
// SetSkippedSteps sets the skippedSteps property value. The skippedSteps property
func (m *JobPhaseDescription) SetSkippedSteps(value *int64)() {
    m.skippedSteps = value
}
// SetStartDate sets the startDate property value. The startDate property
func (m *JobPhaseDescription) SetStartDate(value *string)() {
    m.startDate = value
}
type JobPhaseDescriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletionFractionOfWork()(*float32)
    GetCurrentErrorCount()(*int64)
    GetDescription()(*string)
    GetExpectedStepCount()(*int64)
    GetFinishDate()(*string)
    GetFinishedSteps()(*int64)
    GetName()(*string)
    GetSkippedSteps()(*int64)
    GetStartDate()(*string)
    SetCompletionFractionOfWork(value *float32)()
    SetCurrentErrorCount(value *int64)()
    SetDescription(value *string)()
    SetExpectedStepCount(value *int64)()
    SetFinishDate(value *string)()
    SetFinishedSteps(value *int64)()
    SetName(value *string)()
    SetSkippedSteps(value *int64)()
    SetStartDate(value *string)()
}
