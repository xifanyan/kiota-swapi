package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FolderValuesResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of results found by the search. With paging this may be more than available in the result. NOTE: this number will be 2147483647 (max int) if paging through folder results of a merging meta engine. This is due to a performance optimization when merging results. In such a case, you will need to page forward until the system has received all results (in which case the final page will contain the correct numberResults).
    numberResults *int64
    // The list of folders (categories).
    results []FolderRecordable
    // The status of the response.
    status StatusObjectable
}
// NewFolderValuesResult instantiates a new FolderValuesResult and sets the default values.
func NewFolderValuesResult()(*FolderValuesResult) {
    m := &FolderValuesResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFolderValuesResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFolderValuesResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFolderValuesResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FolderValuesResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FolderValuesResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["numberResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberResults(val)
        }
        return nil
    }
    res["results"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFolderRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FolderRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FolderRecordable)
                }
            }
            m.SetResults(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatusObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(StatusObjectable))
        }
        return nil
    }
    return res
}
// GetNumberResults gets the numberResults property value. The number of results found by the search. With paging this may be more than available in the result. NOTE: this number will be 2147483647 (max int) if paging through folder results of a merging meta engine. This is due to a performance optimization when merging results. In such a case, you will need to page forward until the system has received all results (in which case the final page will contain the correct numberResults).
// returns a *int64 when successful
func (m *FolderValuesResult) GetNumberResults()(*int64) {
    return m.numberResults
}
// GetResults gets the results property value. The list of folders (categories).
// returns a []FolderRecordable when successful
func (m *FolderValuesResult) GetResults()([]FolderRecordable) {
    return m.results
}
// GetStatus gets the status property value. The status of the response.
// returns a StatusObjectable when successful
func (m *FolderValuesResult) GetStatus()(StatusObjectable) {
    return m.status
}
// Serialize serializes information the current object
func (m *FolderValuesResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("numberResults", m.GetNumberResults())
        if err != nil {
            return err
        }
    }
    if m.GetResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResults()))
        for i, v := range m.GetResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("results", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("status", m.GetStatus())
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
func (m *FolderValuesResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumberResults sets the numberResults property value. The number of results found by the search. With paging this may be more than available in the result. NOTE: this number will be 2147483647 (max int) if paging through folder results of a merging meta engine. This is due to a performance optimization when merging results. In such a case, you will need to page forward until the system has received all results (in which case the final page will contain the correct numberResults).
func (m *FolderValuesResult) SetNumberResults(value *int64)() {
    m.numberResults = value
}
// SetResults sets the results property value. The list of folders (categories).
func (m *FolderValuesResult) SetResults(value []FolderRecordable)() {
    m.results = value
}
// SetStatus sets the status property value. The status of the response.
func (m *FolderValuesResult) SetStatus(value StatusObjectable)() {
    m.status = value
}
type FolderValuesResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNumberResults()(*int64)
    GetResults()([]FolderRecordable)
    GetStatus()(StatusObjectable)
    SetNumberResults(value *int64)()
    SetResults(value []FolderRecordable)()
    SetStatus(value StatusObjectable)()
}
