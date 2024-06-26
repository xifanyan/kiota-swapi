package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobDescriptionList a list of JobDescriptions
type JobDescriptionList struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The list of jobs.
    jobs []JobDescriptionable
}
// NewJobDescriptionList instantiates a new JobDescriptionList and sets the default values.
func NewJobDescriptionList()(*JobDescriptionList) {
    m := &JobDescriptionList{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobDescriptionListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateJobDescriptionListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobDescriptionList(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *JobDescriptionList) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *JobDescriptionList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["jobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobDescriptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobDescriptionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JobDescriptionable)
                }
            }
            m.SetJobs(res)
        }
        return nil
    }
    return res
}
// GetJobs gets the jobs property value. The list of jobs.
// returns a []JobDescriptionable when successful
func (m *JobDescriptionList) GetJobs()([]JobDescriptionable) {
    return m.jobs
}
// Serialize serializes information the current object
func (m *JobDescriptionList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("jobs", cast)
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
func (m *JobDescriptionList) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetJobs sets the jobs property value. The list of jobs.
func (m *JobDescriptionList) SetJobs(value []JobDescriptionable)() {
    m.jobs = value
}
type JobDescriptionListable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetJobs()([]JobDescriptionable)
    SetJobs(value []JobDescriptionable)()
}
