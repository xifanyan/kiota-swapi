package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobSubmissionRequest specifies the job in question.
type JobSubmissionRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A longer description for the job in question. The description will appear in the /jobs endpoint and in other job monitoring interfaces.
    description *string
    // A display name for the job in question. The display name will appear in the /jobs endpoint and in other job monitoring interfaces.
    displayName *string
    // The job as it is registered in the server
    jobNameRegisteredInBackend *string
    // Job parameters.
    jobParameters JobSubmissionRequest_jobParametersable
    // The priority. The higher the more resources it will receive compared to other jobs. In this context, '10' is a suitable priority for standard low priority background processes. High priority background processes should receive a priority of 30. All jobs with more priority than 100 will be treated as on-the-fly jobs for which other jobs must be aborted and rescheduled actively. Default is 10.
    priority *int32
}
// NewJobSubmissionRequest instantiates a new JobSubmissionRequest and sets the default values.
func NewJobSubmissionRequest()(*JobSubmissionRequest) {
    m := &JobSubmissionRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobSubmissionRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateJobSubmissionRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobSubmissionRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *JobSubmissionRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. A longer description for the job in question. The description will appear in the /jobs endpoint and in other job monitoring interfaces.
// returns a *string when successful
func (m *JobSubmissionRequest) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. A display name for the job in question. The display name will appear in the /jobs endpoint and in other job monitoring interfaces.
// returns a *string when successful
func (m *JobSubmissionRequest) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *JobSubmissionRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["jobNameRegisteredInBackend"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobNameRegisteredInBackend(val)
        }
        return nil
    }
    res["jobParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJobSubmissionRequest_jobParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobParameters(val.(JobSubmissionRequest_jobParametersable))
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetJobNameRegisteredInBackend gets the jobNameRegisteredInBackend property value. The job as it is registered in the server
// returns a *string when successful
func (m *JobSubmissionRequest) GetJobNameRegisteredInBackend()(*string) {
    return m.jobNameRegisteredInBackend
}
// GetJobParameters gets the jobParameters property value. Job parameters.
// returns a JobSubmissionRequest_jobParametersable when successful
func (m *JobSubmissionRequest) GetJobParameters()(JobSubmissionRequest_jobParametersable) {
    return m.jobParameters
}
// GetPriority gets the priority property value. The priority. The higher the more resources it will receive compared to other jobs. In this context, '10' is a suitable priority for standard low priority background processes. High priority background processes should receive a priority of 30. All jobs with more priority than 100 will be treated as on-the-fly jobs for which other jobs must be aborted and rescheduled actively. Default is 10.
// returns a *int32 when successful
func (m *JobSubmissionRequest) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *JobSubmissionRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jobNameRegisteredInBackend", m.GetJobNameRegisteredInBackend())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("jobParameters", m.GetJobParameters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
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
func (m *JobSubmissionRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. A longer description for the job in question. The description will appear in the /jobs endpoint and in other job monitoring interfaces.
func (m *JobSubmissionRequest) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. A display name for the job in question. The display name will appear in the /jobs endpoint and in other job monitoring interfaces.
func (m *JobSubmissionRequest) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetJobNameRegisteredInBackend sets the jobNameRegisteredInBackend property value. The job as it is registered in the server
func (m *JobSubmissionRequest) SetJobNameRegisteredInBackend(value *string)() {
    m.jobNameRegisteredInBackend = value
}
// SetJobParameters sets the jobParameters property value. Job parameters.
func (m *JobSubmissionRequest) SetJobParameters(value JobSubmissionRequest_jobParametersable)() {
    m.jobParameters = value
}
// SetPriority sets the priority property value. The priority. The higher the more resources it will receive compared to other jobs. In this context, '10' is a suitable priority for standard low priority background processes. High priority background processes should receive a priority of 30. All jobs with more priority than 100 will be treated as on-the-fly jobs for which other jobs must be aborted and rescheduled actively. Default is 10.
func (m *JobSubmissionRequest) SetPriority(value *int32)() {
    m.priority = value
}
type JobSubmissionRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetJobNameRegisteredInBackend()(*string)
    GetJobParameters()(JobSubmissionRequest_jobParametersable)
    GetPriority()(*int32)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetJobNameRegisteredInBackend(value *string)()
    SetJobParameters(value JobSubmissionRequest_jobParametersable)()
    SetPriority(value *int32)()
}
