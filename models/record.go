package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Record struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Content or part of the content of the record (document). This can be 'null'.
    body *string
    // A list of field-value pairs, if any requested.
    fields []Fieldable
    // A list of sets of folder-value pairs (also including the folder display name), if any requested.
    folderSets []FolderSetable
    // The ID of the record (document).
    id *string
    // Shows the page number of the current result. Only populated if body for single documents has been retrieved (not in the result listing of a search).
    page *int32
    // Shows the number of pages in the document. Only populated if body for single documents has been retrieved (not in the result listing of a search).
    pageCount *int32
    // The list position of the record in the list of found records. The first record has rank 1.
    rank *int64
    // Calculated relevance (between 0 and 1) corresponding to how well content and properties match the search expression.
    relevance *float32
    // A human-readable unique field. This field should only be used for logging or investigation, there is no guarantee about the value.
    uniqueField *string
}
// NewRecord instantiates a new Record and sets the default values.
func NewRecord()(*Record) {
    m := &Record{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecord(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Record) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. Content or part of the content of the record (document). This can be 'null'.
// returns a *string when successful
func (m *Record) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Record) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFieldFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fieldable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Fieldable)
                }
            }
            m.SetFields(res)
        }
        return nil
    }
    res["folderSets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFolderSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FolderSetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FolderSetable)
                }
            }
            m.SetFolderSets(res)
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
    res["page"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPage(val)
        }
        return nil
    }
    res["pageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageCount(val)
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
    res["uniqueField"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueField(val)
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. A list of field-value pairs, if any requested.
// returns a []Fieldable when successful
func (m *Record) GetFields()([]Fieldable) {
    return m.fields
}
// GetFolderSets gets the folderSets property value. A list of sets of folder-value pairs (also including the folder display name), if any requested.
// returns a []FolderSetable when successful
func (m *Record) GetFolderSets()([]FolderSetable) {
    return m.folderSets
}
// GetId gets the id property value. The ID of the record (document).
// returns a *string when successful
func (m *Record) GetId()(*string) {
    return m.id
}
// GetPage gets the page property value. Shows the page number of the current result. Only populated if body for single documents has been retrieved (not in the result listing of a search).
// returns a *int32 when successful
func (m *Record) GetPage()(*int32) {
    return m.page
}
// GetPageCount gets the pageCount property value. Shows the number of pages in the document. Only populated if body for single documents has been retrieved (not in the result listing of a search).
// returns a *int32 when successful
func (m *Record) GetPageCount()(*int32) {
    return m.pageCount
}
// GetRank gets the rank property value. The list position of the record in the list of found records. The first record has rank 1.
// returns a *int64 when successful
func (m *Record) GetRank()(*int64) {
    return m.rank
}
// GetRelevance gets the relevance property value. Calculated relevance (between 0 and 1) corresponding to how well content and properties match the search expression.
// returns a *float32 when successful
func (m *Record) GetRelevance()(*float32) {
    return m.relevance
}
// GetUniqueField gets the uniqueField property value. A human-readable unique field. This field should only be used for logging or investigation, there is no guarantee about the value.
// returns a *string when successful
func (m *Record) GetUniqueField()(*string) {
    return m.uniqueField
}
// Serialize serializes information the current object
func (m *Record) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    if m.GetFields() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFields()))
        for i, v := range m.GetFields() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("fields", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFolderSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFolderSets()))
        for i, v := range m.GetFolderSets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("folderSets", cast)
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
        err := writer.WriteInt32Value("page", m.GetPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pageCount", m.GetPageCount())
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
        err := writer.WriteStringValue("uniqueField", m.GetUniqueField())
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
func (m *Record) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. Content or part of the content of the record (document). This can be 'null'.
func (m *Record) SetBody(value *string)() {
    m.body = value
}
// SetFields sets the fields property value. A list of field-value pairs, if any requested.
func (m *Record) SetFields(value []Fieldable)() {
    m.fields = value
}
// SetFolderSets sets the folderSets property value. A list of sets of folder-value pairs (also including the folder display name), if any requested.
func (m *Record) SetFolderSets(value []FolderSetable)() {
    m.folderSets = value
}
// SetId sets the id property value. The ID of the record (document).
func (m *Record) SetId(value *string)() {
    m.id = value
}
// SetPage sets the page property value. Shows the page number of the current result. Only populated if body for single documents has been retrieved (not in the result listing of a search).
func (m *Record) SetPage(value *int32)() {
    m.page = value
}
// SetPageCount sets the pageCount property value. Shows the number of pages in the document. Only populated if body for single documents has been retrieved (not in the result listing of a search).
func (m *Record) SetPageCount(value *int32)() {
    m.pageCount = value
}
// SetRank sets the rank property value. The list position of the record in the list of found records. The first record has rank 1.
func (m *Record) SetRank(value *int64)() {
    m.rank = value
}
// SetRelevance sets the relevance property value. Calculated relevance (between 0 and 1) corresponding to how well content and properties match the search expression.
func (m *Record) SetRelevance(value *float32)() {
    m.relevance = value
}
// SetUniqueField sets the uniqueField property value. A human-readable unique field. This field should only be used for logging or investigation, there is no guarantee about the value.
func (m *Record) SetUniqueField(value *string)() {
    m.uniqueField = value
}
type Recordable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetFields()([]Fieldable)
    GetFolderSets()([]FolderSetable)
    GetId()(*string)
    GetPage()(*int32)
    GetPageCount()(*int32)
    GetRank()(*int64)
    GetRelevance()(*float32)
    GetUniqueField()(*string)
    SetBody(value *string)()
    SetFields(value []Fieldable)()
    SetFolderSets(value []FolderSetable)()
    SetId(value *string)()
    SetPage(value *int32)()
    SetPageCount(value *int32)()
    SetRank(value *int64)()
    SetRelevance(value *float32)()
    SetUniqueField(value *string)()
}
