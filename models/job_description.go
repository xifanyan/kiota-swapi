package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobDescription a job description.
type JobDescription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cancelDate property
    cancelDate *string
    // The collectionName property
    collectionName *string
    // The completionPercentage property
    completionPercentage *float32
    // The creationDate property
    creationDate *string
    // The currentPhaseIndex property
    currentPhaseIndex *int32
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The existsErrorFile property
    existsErrorFile *bool
    // The existsJobResultFile property
    existsJobResultFile *bool
    // The finishDate property
    finishDate *string
    // The id property
    id *string
    // The lastPauseDate property
    lastPauseDate *string
    // The owner property
    owner *string
    // The phases property
    phases []JobPhaseDescriptionable
    // The priority property
    priority *int32
    // The startDate property
    startDate *string
    // The status property
    status *JobDescription_status
    // The topLevel property
    topLevel *bool
    // The totalErrorCount property
    totalErrorCount *int64
    // The totalPauseTimeMs property
    totalPauseTimeMs *int64
    // The totalRuntimeMs property
    totalRuntimeMs *int64
    // The type property
    typeEscaped *string
}
// NewJobDescription instantiates a new JobDescription and sets the default values.
func NewJobDescription()(*JobDescription) {
    m := &JobDescription{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobDescriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateJobDescriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobDescription(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *JobDescription) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCancelDate gets the cancelDate property value. The cancelDate property
// returns a *string when successful
func (m *JobDescription) GetCancelDate()(*string) {
    return m.cancelDate
}
// GetCollectionName gets the collectionName property value. The collectionName property
// returns a *string when successful
func (m *JobDescription) GetCollectionName()(*string) {
    return m.collectionName
}
// GetCompletionPercentage gets the completionPercentage property value. The completionPercentage property
// returns a *float32 when successful
func (m *JobDescription) GetCompletionPercentage()(*float32) {
    return m.completionPercentage
}
// GetCreationDate gets the creationDate property value. The creationDate property
// returns a *string when successful
func (m *JobDescription) GetCreationDate()(*string) {
    return m.creationDate
}
// GetCurrentPhaseIndex gets the currentPhaseIndex property value. The currentPhaseIndex property
// returns a *int32 when successful
func (m *JobDescription) GetCurrentPhaseIndex()(*int32) {
    return m.currentPhaseIndex
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *JobDescription) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *JobDescription) GetDisplayName()(*string) {
    return m.displayName
}
// GetExistsErrorFile gets the existsErrorFile property value. The existsErrorFile property
// returns a *bool when successful
func (m *JobDescription) GetExistsErrorFile()(*bool) {
    return m.existsErrorFile
}
// GetExistsJobResultFile gets the existsJobResultFile property value. The existsJobResultFile property
// returns a *bool when successful
func (m *JobDescription) GetExistsJobResultFile()(*bool) {
    return m.existsJobResultFile
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *JobDescription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cancelDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancelDate(val)
        }
        return nil
    }
    res["collectionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollectionName(val)
        }
        return nil
    }
    res["completionPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionPercentage(val)
        }
        return nil
    }
    res["creationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDate(val)
        }
        return nil
    }
    res["currentPhaseIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPhaseIndex(val)
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
    res["existsErrorFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExistsErrorFile(val)
        }
        return nil
    }
    res["existsJobResultFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExistsJobResultFile(val)
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
    res["lastPauseDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastPauseDate(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["phases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobPhaseDescriptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobPhaseDescriptionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JobPhaseDescriptionable)
                }
            }
            m.SetPhases(res)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseJobDescription_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*JobDescription_status))
        }
        return nil
    }
    res["topLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopLevel(val)
        }
        return nil
    }
    res["totalErrorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalErrorCount(val)
        }
        return nil
    }
    res["totalPauseTimeMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPauseTimeMs(val)
        }
        return nil
    }
    res["totalRuntimeMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRuntimeMs(val)
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
// GetFinishDate gets the finishDate property value. The finishDate property
// returns a *string when successful
func (m *JobDescription) GetFinishDate()(*string) {
    return m.finishDate
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *JobDescription) GetId()(*string) {
    return m.id
}
// GetLastPauseDate gets the lastPauseDate property value. The lastPauseDate property
// returns a *string when successful
func (m *JobDescription) GetLastPauseDate()(*string) {
    return m.lastPauseDate
}
// GetOwner gets the owner property value. The owner property
// returns a *string when successful
func (m *JobDescription) GetOwner()(*string) {
    return m.owner
}
// GetPhases gets the phases property value. The phases property
// returns a []JobPhaseDescriptionable when successful
func (m *JobDescription) GetPhases()([]JobPhaseDescriptionable) {
    return m.phases
}
// GetPriority gets the priority property value. The priority property
// returns a *int32 when successful
func (m *JobDescription) GetPriority()(*int32) {
    return m.priority
}
// GetStartDate gets the startDate property value. The startDate property
// returns a *string when successful
func (m *JobDescription) GetStartDate()(*string) {
    return m.startDate
}
// GetStatus gets the status property value. The status property
// returns a *JobDescription_status when successful
func (m *JobDescription) GetStatus()(*JobDescription_status) {
    return m.status
}
// GetTopLevel gets the topLevel property value. The topLevel property
// returns a *bool when successful
func (m *JobDescription) GetTopLevel()(*bool) {
    return m.topLevel
}
// GetTotalErrorCount gets the totalErrorCount property value. The totalErrorCount property
// returns a *int64 when successful
func (m *JobDescription) GetTotalErrorCount()(*int64) {
    return m.totalErrorCount
}
// GetTotalPauseTimeMs gets the totalPauseTimeMs property value. The totalPauseTimeMs property
// returns a *int64 when successful
func (m *JobDescription) GetTotalPauseTimeMs()(*int64) {
    return m.totalPauseTimeMs
}
// GetTotalRuntimeMs gets the totalRuntimeMs property value. The totalRuntimeMs property
// returns a *int64 when successful
func (m *JobDescription) GetTotalRuntimeMs()(*int64) {
    return m.totalRuntimeMs
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *JobDescription) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *JobDescription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cancelDate", m.GetCancelDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("collectionName", m.GetCollectionName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("completionPercentage", m.GetCompletionPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("creationDate", m.GetCreationDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("currentPhaseIndex", m.GetCurrentPhaseIndex())
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
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("existsErrorFile", m.GetExistsErrorFile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("existsJobResultFile", m.GetExistsJobResultFile())
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
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastPauseDate", m.GetLastPauseDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    if m.GetPhases() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhases()))
        for i, v := range m.GetPhases() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("phases", cast)
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
        err := writer.WriteStringValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("topLevel", m.GetTopLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalErrorCount", m.GetTotalErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalPauseTimeMs", m.GetTotalPauseTimeMs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalRuntimeMs", m.GetTotalRuntimeMs())
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
func (m *JobDescription) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCancelDate sets the cancelDate property value. The cancelDate property
func (m *JobDescription) SetCancelDate(value *string)() {
    m.cancelDate = value
}
// SetCollectionName sets the collectionName property value. The collectionName property
func (m *JobDescription) SetCollectionName(value *string)() {
    m.collectionName = value
}
// SetCompletionPercentage sets the completionPercentage property value. The completionPercentage property
func (m *JobDescription) SetCompletionPercentage(value *float32)() {
    m.completionPercentage = value
}
// SetCreationDate sets the creationDate property value. The creationDate property
func (m *JobDescription) SetCreationDate(value *string)() {
    m.creationDate = value
}
// SetCurrentPhaseIndex sets the currentPhaseIndex property value. The currentPhaseIndex property
func (m *JobDescription) SetCurrentPhaseIndex(value *int32)() {
    m.currentPhaseIndex = value
}
// SetDescription sets the description property value. The description property
func (m *JobDescription) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *JobDescription) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExistsErrorFile sets the existsErrorFile property value. The existsErrorFile property
func (m *JobDescription) SetExistsErrorFile(value *bool)() {
    m.existsErrorFile = value
}
// SetExistsJobResultFile sets the existsJobResultFile property value. The existsJobResultFile property
func (m *JobDescription) SetExistsJobResultFile(value *bool)() {
    m.existsJobResultFile = value
}
// SetFinishDate sets the finishDate property value. The finishDate property
func (m *JobDescription) SetFinishDate(value *string)() {
    m.finishDate = value
}
// SetId sets the id property value. The id property
func (m *JobDescription) SetId(value *string)() {
    m.id = value
}
// SetLastPauseDate sets the lastPauseDate property value. The lastPauseDate property
func (m *JobDescription) SetLastPauseDate(value *string)() {
    m.lastPauseDate = value
}
// SetOwner sets the owner property value. The owner property
func (m *JobDescription) SetOwner(value *string)() {
    m.owner = value
}
// SetPhases sets the phases property value. The phases property
func (m *JobDescription) SetPhases(value []JobPhaseDescriptionable)() {
    m.phases = value
}
// SetPriority sets the priority property value. The priority property
func (m *JobDescription) SetPriority(value *int32)() {
    m.priority = value
}
// SetStartDate sets the startDate property value. The startDate property
func (m *JobDescription) SetStartDate(value *string)() {
    m.startDate = value
}
// SetStatus sets the status property value. The status property
func (m *JobDescription) SetStatus(value *JobDescription_status)() {
    m.status = value
}
// SetTopLevel sets the topLevel property value. The topLevel property
func (m *JobDescription) SetTopLevel(value *bool)() {
    m.topLevel = value
}
// SetTotalErrorCount sets the totalErrorCount property value. The totalErrorCount property
func (m *JobDescription) SetTotalErrorCount(value *int64)() {
    m.totalErrorCount = value
}
// SetTotalPauseTimeMs sets the totalPauseTimeMs property value. The totalPauseTimeMs property
func (m *JobDescription) SetTotalPauseTimeMs(value *int64)() {
    m.totalPauseTimeMs = value
}
// SetTotalRuntimeMs sets the totalRuntimeMs property value. The totalRuntimeMs property
func (m *JobDescription) SetTotalRuntimeMs(value *int64)() {
    m.totalRuntimeMs = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *JobDescription) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type JobDescriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCancelDate()(*string)
    GetCollectionName()(*string)
    GetCompletionPercentage()(*float32)
    GetCreationDate()(*string)
    GetCurrentPhaseIndex()(*int32)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExistsErrorFile()(*bool)
    GetExistsJobResultFile()(*bool)
    GetFinishDate()(*string)
    GetId()(*string)
    GetLastPauseDate()(*string)
    GetOwner()(*string)
    GetPhases()([]JobPhaseDescriptionable)
    GetPriority()(*int32)
    GetStartDate()(*string)
    GetStatus()(*JobDescription_status)
    GetTopLevel()(*bool)
    GetTotalErrorCount()(*int64)
    GetTotalPauseTimeMs()(*int64)
    GetTotalRuntimeMs()(*int64)
    GetTypeEscaped()(*string)
    SetCancelDate(value *string)()
    SetCollectionName(value *string)()
    SetCompletionPercentage(value *float32)()
    SetCreationDate(value *string)()
    SetCurrentPhaseIndex(value *int32)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExistsErrorFile(value *bool)()
    SetExistsJobResultFile(value *bool)()
    SetFinishDate(value *string)()
    SetId(value *string)()
    SetLastPauseDate(value *string)()
    SetOwner(value *string)()
    SetPhases(value []JobPhaseDescriptionable)()
    SetPriority(value *int32)()
    SetStartDate(value *string)()
    SetStatus(value *JobDescription_status)()
    SetTopLevel(value *bool)()
    SetTotalErrorCount(value *int64)()
    SetTotalPauseTimeMs(value *int64)()
    SetTotalRuntimeMs(value *int64)()
    SetTypeEscaped(value *string)()
}
