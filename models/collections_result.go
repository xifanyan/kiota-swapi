package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CollectionsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of results found by the search. With paging this may be more than available in the result.
    numberResults *int64
    // The list of collections.
    results []Collectionable
    // The status of the response.
    status StatusObjectable
}
// NewCollectionsResult instantiates a new CollectionsResult and sets the default values.
func NewCollectionsResult()(*CollectionsResult) {
    m := &CollectionsResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCollectionsResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCollectionsResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollectionsResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CollectionsResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CollectionsResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Collectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Collectionable)
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
// GetNumberResults gets the numberResults property value. The number of results found by the search. With paging this may be more than available in the result.
// returns a *int64 when successful
func (m *CollectionsResult) GetNumberResults()(*int64) {
    return m.numberResults
}
// GetResults gets the results property value. The list of collections.
// returns a []Collectionable when successful
func (m *CollectionsResult) GetResults()([]Collectionable) {
    return m.results
}
// GetStatus gets the status property value. The status of the response.
// returns a StatusObjectable when successful
func (m *CollectionsResult) GetStatus()(StatusObjectable) {
    return m.status
}
// Serialize serializes information the current object
func (m *CollectionsResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CollectionsResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumberResults sets the numberResults property value. The number of results found by the search. With paging this may be more than available in the result.
func (m *CollectionsResult) SetNumberResults(value *int64)() {
    m.numberResults = value
}
// SetResults sets the results property value. The list of collections.
func (m *CollectionsResult) SetResults(value []Collectionable)() {
    m.results = value
}
// SetStatus sets the status property value. The status of the response.
func (m *CollectionsResult) SetStatus(value StatusObjectable)() {
    m.status = value
}
type CollectionsResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNumberResults()(*int64)
    GetResults()([]Collectionable)
    GetStatus()(StatusObjectable)
    SetNumberResults(value *int64)()
    SetResults(value []Collectionable)()
    SetStatus(value StatusObjectable)()
}
