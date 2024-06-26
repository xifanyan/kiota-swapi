package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FolderRecord the list of folders (categories).
type FolderRecord struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of documents associated with the given folder (category).
    count *int64
    // The display name of the folder. Generally, for most folders, this will be the same as the ID.
    displayName *string
    // The folder ID.
    id *string
    // The list position of the record in the sorted list of folders. The first folder has rank 1.
    rank *int64
    // The accumulated relevancy (between 0 and 1) of collection records in the search result relating to this folder.
    relevance *float32
}
// NewFolderRecord instantiates a new FolderRecord and sets the default values.
func NewFolderRecord()(*FolderRecord) {
    m := &FolderRecord{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFolderRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFolderRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFolderRecord(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FolderRecord) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. The number of documents associated with the given folder (category).
// returns a *int64 when successful
func (m *FolderRecord) GetCount()(*int64) {
    return m.count
}
// GetDisplayName gets the displayName property value. The display name of the folder. Generally, for most folders, this will be the same as the ID.
// returns a *string when successful
func (m *FolderRecord) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FolderRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
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
    res["rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["relevance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelevance(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The folder ID.
// returns a *string when successful
func (m *FolderRecord) GetId()(*string) {
    return m.id
}
// GetRank gets the rank property value. The list position of the record in the sorted list of folders. The first folder has rank 1.
// returns a *int64 when successful
func (m *FolderRecord) GetRank()(*int64) {
    return m.rank
}
// GetRelevance gets the relevance property value. The accumulated relevancy (between 0 and 1) of collection records in the search result relating to this folder.
// returns a *float32 when successful
func (m *FolderRecord) GetRelevance()(*float32) {
    return m.relevance
}
// Serialize serializes information the current object
func (m *FolderRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("count", m.GetCount())
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
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("relevance", m.GetRelevance())
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
func (m *FolderRecord) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. The number of documents associated with the given folder (category).
func (m *FolderRecord) SetCount(value *int64)() {
    m.count = value
}
// SetDisplayName sets the displayName property value. The display name of the folder. Generally, for most folders, this will be the same as the ID.
func (m *FolderRecord) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The folder ID.
func (m *FolderRecord) SetId(value *string)() {
    m.id = value
}
// SetRank sets the rank property value. The list position of the record in the sorted list of folders. The first folder has rank 1.
func (m *FolderRecord) SetRank(value *int64)() {
    m.rank = value
}
// SetRelevance sets the relevance property value. The accumulated relevancy (between 0 and 1) of collection records in the search result relating to this folder.
func (m *FolderRecord) SetRelevance(value *float32)() {
    m.relevance = value
}
type FolderRecordable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetDisplayName()(*string)
    GetId()(*string)
    GetRank()(*int64)
    GetRelevance()(*float32)
    SetCount(value *int64)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetRank(value *int64)()
    SetRelevance(value *float32)()
}
