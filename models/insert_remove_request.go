package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InsertRemoveRequest the request which defines the data for new records and/or ids of old records which are to be deleted (replaced). 
type InsertRemoveRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // 0-n records to be created.
    newRecords []RecordDataable
    // 0-n record ids which are to be deleted. Specifying a record as both 'delete' and 'new' has re-insert semantics. The specified ids have to be in one of the following formats: 1. they can resemble the 'id' as returned by the records endpoint (aka url-safe string). 2. they can be a folder id (only applicable if you are connected against a folder collection). 3. If connected with a folder collection, you can optionally specify folderId::VALUE here where VALUE is a folder id.
    recordsToDelete []string
    // A query.
    recordsToDeleteByQuery SearchRequestable
}
// NewInsertRemoveRequest instantiates a new InsertRemoveRequest and sets the default values.
func NewInsertRemoveRequest()(*InsertRemoveRequest) {
    m := &InsertRemoveRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInsertRemoveRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInsertRemoveRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInsertRemoveRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InsertRemoveRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InsertRemoveRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecordDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecordDataable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RecordDataable)
                }
            }
            m.SetNewRecords(res)
        }
        return nil
    }
    res["recordsToDelete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRecordsToDelete(res)
        }
        return nil
    }
    res["recordsToDeleteByQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordsToDeleteByQuery(val.(SearchRequestable))
        }
        return nil
    }
    return res
}
// GetNewRecords gets the newRecords property value. 0-n records to be created.
// returns a []RecordDataable when successful
func (m *InsertRemoveRequest) GetNewRecords()([]RecordDataable) {
    return m.newRecords
}
// GetRecordsToDelete gets the recordsToDelete property value. 0-n record ids which are to be deleted. Specifying a record as both 'delete' and 'new' has re-insert semantics. The specified ids have to be in one of the following formats: 1. they can resemble the 'id' as returned by the records endpoint (aka url-safe string). 2. they can be a folder id (only applicable if you are connected against a folder collection). 3. If connected with a folder collection, you can optionally specify folderId::VALUE here where VALUE is a folder id.
// returns a []string when successful
func (m *InsertRemoveRequest) GetRecordsToDelete()([]string) {
    return m.recordsToDelete
}
// GetRecordsToDeleteByQuery gets the recordsToDeleteByQuery property value. A query.
// returns a SearchRequestable when successful
func (m *InsertRemoveRequest) GetRecordsToDeleteByQuery()(SearchRequestable) {
    return m.recordsToDeleteByQuery
}
// Serialize serializes information the current object
func (m *InsertRemoveRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNewRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNewRecords()))
        for i, v := range m.GetNewRecords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("newRecords", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecordsToDelete() != nil {
        err := writer.WriteCollectionOfStringValues("recordsToDelete", m.GetRecordsToDelete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recordsToDeleteByQuery", m.GetRecordsToDeleteByQuery())
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
func (m *InsertRemoveRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewRecords sets the newRecords property value. 0-n records to be created.
func (m *InsertRemoveRequest) SetNewRecords(value []RecordDataable)() {
    m.newRecords = value
}
// SetRecordsToDelete sets the recordsToDelete property value. 0-n record ids which are to be deleted. Specifying a record as both 'delete' and 'new' has re-insert semantics. The specified ids have to be in one of the following formats: 1. they can resemble the 'id' as returned by the records endpoint (aka url-safe string). 2. they can be a folder id (only applicable if you are connected against a folder collection). 3. If connected with a folder collection, you can optionally specify folderId::VALUE here where VALUE is a folder id.
func (m *InsertRemoveRequest) SetRecordsToDelete(value []string)() {
    m.recordsToDelete = value
}
// SetRecordsToDeleteByQuery sets the recordsToDeleteByQuery property value. A query.
func (m *InsertRemoveRequest) SetRecordsToDeleteByQuery(value SearchRequestable)() {
    m.recordsToDeleteByQuery = value
}
type InsertRemoveRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNewRecords()([]RecordDataable)
    GetRecordsToDelete()([]string)
    GetRecordsToDeleteByQuery()(SearchRequestable)
    SetNewRecords(value []RecordDataable)()
    SetRecordsToDelete(value []string)()
    SetRecordsToDeleteByQuery(value SearchRequestable)()
}
