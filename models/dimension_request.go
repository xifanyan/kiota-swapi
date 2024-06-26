package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DimensionRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The field defining the dimension.
    field *string
    // Only for folder fields: provide an order criterion for the results. The syntax is the same as 'order' for the records endpoint, i.e. <fieldName> or <filedName>:asc or <fieldName>:desc. In this context, <fieldName> is any folder property. In addition, <fieldName> can be one of the special names 'count', 'relevance', or 'name'. Here, 'name' is the category display name.
    folderOrder *string
    // Define which is the first element to return (first is 0). The system will return the next 'pageSize' following items. Default is to return all matching members. You cannot specify both 'page' and 'offset'.
    offset *int32
    // Define which page to retrieve (counting starts at 1). Default is to return all matching members. You cannot specify both 'page' and 'offset'.
    page *int32
    // Defines the page size. Only used if 'page' or 'offset' is set.
    pageSize *int32
    // Only for folder fields: optional query to constrain the folders returned. Examples are rm_folder_id=uniquevalue or rm_display_name=displayname
    restrictFoldersByQuery *string
    // A list of values. For folder fields, the values are folder ids. For fields of numeric or date type, each value can have one of two forms: the inclusive variant '[start,end]' or the exclusive variant '[start,end[' . For DATE fields, both start and end are either milli seconds since 1970 in UTC time followed by the suffix 'L' or a date format in one of the formats 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX', 'yyyy-MM-dd'T'HH:mm:ss.SSS', 'yyyy-MM-dd'T'HH:mm:ss', 'yyyy-MM-dd'T'HH:mm', 'yyyy-MM-dd'T'HH', 'yyyy-MM-dd', 'yyyy-MM', 'yyyy' . Example for a numeric field: [4,5] which counts documents having either value 4 and the value 5. Another example for a numeric field: [4,5[ counts documents having value 4 but not 5. Example for a date field: [2000, 2000-10-01] counts documents with value between 2000-01-01 and 200-10-01 23:59:59.999 (inclusive).    
    restrictValuesByList []string
    // A comma-separated list of field names. It is only used for folder fields and allows to fetch properties of the returned folders (like rm_display_name).
    returnedFields *string
    // Defines if members with count=0 should be returned. Default is false.
    returnEmptyMembers *bool
}
// NewDimensionRequest instantiates a new DimensionRequest and sets the default values.
func NewDimensionRequest()(*DimensionRequest) {
    m := &DimensionRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDimensionRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDimensionRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDimensionRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DimensionRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetField gets the field property value. The field defining the dimension.
// returns a *string when successful
func (m *DimensionRequest) GetField()(*string) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DimensionRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["field"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetField(val)
        }
        return nil
    }
    res["folderOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderOrder(val)
        }
        return nil
    }
    res["offset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffset(val)
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
    res["pageSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageSize(val)
        }
        return nil
    }
    res["restrictFoldersByQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictFoldersByQuery(val)
        }
        return nil
    }
    res["restrictValuesByList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRestrictValuesByList(res)
        }
        return nil
    }
    res["returnedFields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnedFields(val)
        }
        return nil
    }
    res["returnEmptyMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnEmptyMembers(val)
        }
        return nil
    }
    return res
}
// GetFolderOrder gets the folderOrder property value. Only for folder fields: provide an order criterion for the results. The syntax is the same as 'order' for the records endpoint, i.e. <fieldName> or <filedName>:asc or <fieldName>:desc. In this context, <fieldName> is any folder property. In addition, <fieldName> can be one of the special names 'count', 'relevance', or 'name'. Here, 'name' is the category display name.
// returns a *string when successful
func (m *DimensionRequest) GetFolderOrder()(*string) {
    return m.folderOrder
}
// GetOffset gets the offset property value. Define which is the first element to return (first is 0). The system will return the next 'pageSize' following items. Default is to return all matching members. You cannot specify both 'page' and 'offset'.
// returns a *int32 when successful
func (m *DimensionRequest) GetOffset()(*int32) {
    return m.offset
}
// GetPage gets the page property value. Define which page to retrieve (counting starts at 1). Default is to return all matching members. You cannot specify both 'page' and 'offset'.
// returns a *int32 when successful
func (m *DimensionRequest) GetPage()(*int32) {
    return m.page
}
// GetPageSize gets the pageSize property value. Defines the page size. Only used if 'page' or 'offset' is set.
// returns a *int32 when successful
func (m *DimensionRequest) GetPageSize()(*int32) {
    return m.pageSize
}
// GetRestrictFoldersByQuery gets the restrictFoldersByQuery property value. Only for folder fields: optional query to constrain the folders returned. Examples are rm_folder_id=uniquevalue or rm_display_name=displayname
// returns a *string when successful
func (m *DimensionRequest) GetRestrictFoldersByQuery()(*string) {
    return m.restrictFoldersByQuery
}
// GetRestrictValuesByList gets the restrictValuesByList property value. A list of values. For folder fields, the values are folder ids. For fields of numeric or date type, each value can have one of two forms: the inclusive variant '[start,end]' or the exclusive variant '[start,end[' . For DATE fields, both start and end are either milli seconds since 1970 in UTC time followed by the suffix 'L' or a date format in one of the formats 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX', 'yyyy-MM-dd'T'HH:mm:ss.SSS', 'yyyy-MM-dd'T'HH:mm:ss', 'yyyy-MM-dd'T'HH:mm', 'yyyy-MM-dd'T'HH', 'yyyy-MM-dd', 'yyyy-MM', 'yyyy' . Example for a numeric field: [4,5] which counts documents having either value 4 and the value 5. Another example for a numeric field: [4,5[ counts documents having value 4 but not 5. Example for a date field: [2000, 2000-10-01] counts documents with value between 2000-01-01 and 200-10-01 23:59:59.999 (inclusive).    
// returns a []string when successful
func (m *DimensionRequest) GetRestrictValuesByList()([]string) {
    return m.restrictValuesByList
}
// GetReturnedFields gets the returnedFields property value. A comma-separated list of field names. It is only used for folder fields and allows to fetch properties of the returned folders (like rm_display_name).
// returns a *string when successful
func (m *DimensionRequest) GetReturnedFields()(*string) {
    return m.returnedFields
}
// GetReturnEmptyMembers gets the returnEmptyMembers property value. Defines if members with count=0 should be returned. Default is false.
// returns a *bool when successful
func (m *DimensionRequest) GetReturnEmptyMembers()(*bool) {
    return m.returnEmptyMembers
}
// Serialize serializes information the current object
func (m *DimensionRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("field", m.GetField())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("folderOrder", m.GetFolderOrder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("offset", m.GetOffset())
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
        err := writer.WriteInt32Value("pageSize", m.GetPageSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("restrictFoldersByQuery", m.GetRestrictFoldersByQuery())
        if err != nil {
            return err
        }
    }
    if m.GetRestrictValuesByList() != nil {
        err := writer.WriteCollectionOfStringValues("restrictValuesByList", m.GetRestrictValuesByList())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("returnedFields", m.GetReturnedFields())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("returnEmptyMembers", m.GetReturnEmptyMembers())
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
func (m *DimensionRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetField sets the field property value. The field defining the dimension.
func (m *DimensionRequest) SetField(value *string)() {
    m.field = value
}
// SetFolderOrder sets the folderOrder property value. Only for folder fields: provide an order criterion for the results. The syntax is the same as 'order' for the records endpoint, i.e. <fieldName> or <filedName>:asc or <fieldName>:desc. In this context, <fieldName> is any folder property. In addition, <fieldName> can be one of the special names 'count', 'relevance', or 'name'. Here, 'name' is the category display name.
func (m *DimensionRequest) SetFolderOrder(value *string)() {
    m.folderOrder = value
}
// SetOffset sets the offset property value. Define which is the first element to return (first is 0). The system will return the next 'pageSize' following items. Default is to return all matching members. You cannot specify both 'page' and 'offset'.
func (m *DimensionRequest) SetOffset(value *int32)() {
    m.offset = value
}
// SetPage sets the page property value. Define which page to retrieve (counting starts at 1). Default is to return all matching members. You cannot specify both 'page' and 'offset'.
func (m *DimensionRequest) SetPage(value *int32)() {
    m.page = value
}
// SetPageSize sets the pageSize property value. Defines the page size. Only used if 'page' or 'offset' is set.
func (m *DimensionRequest) SetPageSize(value *int32)() {
    m.pageSize = value
}
// SetRestrictFoldersByQuery sets the restrictFoldersByQuery property value. Only for folder fields: optional query to constrain the folders returned. Examples are rm_folder_id=uniquevalue or rm_display_name=displayname
func (m *DimensionRequest) SetRestrictFoldersByQuery(value *string)() {
    m.restrictFoldersByQuery = value
}
// SetRestrictValuesByList sets the restrictValuesByList property value. A list of values. For folder fields, the values are folder ids. For fields of numeric or date type, each value can have one of two forms: the inclusive variant '[start,end]' or the exclusive variant '[start,end[' . For DATE fields, both start and end are either milli seconds since 1970 in UTC time followed by the suffix 'L' or a date format in one of the formats 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX', 'yyyy-MM-dd'T'HH:mm:ss.SSS', 'yyyy-MM-dd'T'HH:mm:ss', 'yyyy-MM-dd'T'HH:mm', 'yyyy-MM-dd'T'HH', 'yyyy-MM-dd', 'yyyy-MM', 'yyyy' . Example for a numeric field: [4,5] which counts documents having either value 4 and the value 5. Another example for a numeric field: [4,5[ counts documents having value 4 but not 5. Example for a date field: [2000, 2000-10-01] counts documents with value between 2000-01-01 and 200-10-01 23:59:59.999 (inclusive).    
func (m *DimensionRequest) SetRestrictValuesByList(value []string)() {
    m.restrictValuesByList = value
}
// SetReturnedFields sets the returnedFields property value. A comma-separated list of field names. It is only used for folder fields and allows to fetch properties of the returned folders (like rm_display_name).
func (m *DimensionRequest) SetReturnedFields(value *string)() {
    m.returnedFields = value
}
// SetReturnEmptyMembers sets the returnEmptyMembers property value. Defines if members with count=0 should be returned. Default is false.
func (m *DimensionRequest) SetReturnEmptyMembers(value *bool)() {
    m.returnEmptyMembers = value
}
type DimensionRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetField()(*string)
    GetFolderOrder()(*string)
    GetOffset()(*int32)
    GetPage()(*int32)
    GetPageSize()(*int32)
    GetRestrictFoldersByQuery()(*string)
    GetRestrictValuesByList()([]string)
    GetReturnedFields()(*string)
    GetReturnEmptyMembers()(*bool)
    SetField(value *string)()
    SetFolderOrder(value *string)()
    SetOffset(value *int32)()
    SetPage(value *int32)()
    SetPageSize(value *int32)()
    SetRestrictFoldersByQuery(value *string)()
    SetRestrictValuesByList(value []string)()
    SetReturnedFields(value *string)()
    SetReturnEmptyMembers(value *bool)()
}
