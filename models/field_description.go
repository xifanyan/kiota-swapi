package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FieldDescription the list of fields, including folder fields.
type FieldDescription struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A field display name. The value is never null, it falls back to the field id if there is no display name. Display names can be configured in the server configuration.
    displayName *string
    // The field ID. The ID is the one to be used where the API requires a field ID (or folder field ID, depending on the context).
    id *string
    // True if the field is a folder field (taxonomy, smart filter). A document can be associated with a folder in such a folder field (or with multiple folders, depending on folder field). A folder field can be used for filtering. Values (e.g. the document count) can be accumulated for folders.
    isFolderField *bool
    // True if the field is a multi-value folder collection (taxonomy). A folder collection can be used as filter. Each document can be associated with any number of folders of such a folder collection.
    isMultivalueFolderCollection *bool
    // True if a folder search on the field can be restricted by prefix.
    isPrefixSearchable *bool
    // True if the field can be used for sorting.
    isSortable *bool
    // The data type of the field: identifier (text), content (text), text, integer, float, date, single-value taxonomy, multi-value taxonomy.   The type 'binary' is special in that it accepts and returns binary streams. A main use-case is to access Binary storages configured for a singleMindServer. To this end, use projectId=singleMindServer.NAME and collectionId=rm_storage:Native files for 'Native files' or collectionId=rm_storage:Image files to access image files. The syntax is 'rm_storage:' followed by the value configured as 'Storage file type' in the single mind server configuration.   The type 'hierarchy' allows to implement tree structures. Its value is a pointer to the nodes parent id. It is intended to be used in folder collections only.   The type 'parameterizedLong' is a multi-value type: a single document can have multiple values in such a field, and each value has a 'parameter' (a string) and a long value. Such fields are supposed to contain 'few' parameters, but many different long values. It is possible to insert/change values of the form '<parameter>.<long>' where '<parameter>' is some string chosen by the API user (example: rank.1234). Assuming such a field is called ranks, the associated values are returned by the records endpoint with fields=ranks.A, ranks.B where 'A' and 'B' are parameters. It is possible to search for documents with such values by means of field bases searches of the form 'ranks.A > 42' or 'ranks.B = 1234'. It is also possible to sort by these values: specify the sort field as 'ranks.A:desc' or 'ranks.B:asc'. Fields of type 'texts' contain a list of strings.
    typeEscaped *FieldDescription_type
}
// NewFieldDescription instantiates a new FieldDescription and sets the default values.
func NewFieldDescription()(*FieldDescription) {
    m := &FieldDescription{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFieldDescriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFieldDescriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFieldDescription(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FieldDescription) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A field display name. The value is never null, it falls back to the field id if there is no display name. Display names can be configured in the server configuration.
// returns a *string when successful
func (m *FieldDescription) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FieldDescription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["isFolderField"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFolderField(val)
        }
        return nil
    }
    res["isMultivalueFolderCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMultivalueFolderCollection(val)
        }
        return nil
    }
    res["isPrefixSearchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrefixSearchable(val)
        }
        return nil
    }
    res["isSortable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSortable(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFieldDescription_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*FieldDescription_type))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The field ID. The ID is the one to be used where the API requires a field ID (or folder field ID, depending on the context).
// returns a *string when successful
func (m *FieldDescription) GetId()(*string) {
    return m.id
}
// GetIsFolderField gets the isFolderField property value. True if the field is a folder field (taxonomy, smart filter). A document can be associated with a folder in such a folder field (or with multiple folders, depending on folder field). A folder field can be used for filtering. Values (e.g. the document count) can be accumulated for folders.
// returns a *bool when successful
func (m *FieldDescription) GetIsFolderField()(*bool) {
    return m.isFolderField
}
// GetIsMultivalueFolderCollection gets the isMultivalueFolderCollection property value. True if the field is a multi-value folder collection (taxonomy). A folder collection can be used as filter. Each document can be associated with any number of folders of such a folder collection.
// returns a *bool when successful
func (m *FieldDescription) GetIsMultivalueFolderCollection()(*bool) {
    return m.isMultivalueFolderCollection
}
// GetIsPrefixSearchable gets the isPrefixSearchable property value. True if a folder search on the field can be restricted by prefix.
// returns a *bool when successful
func (m *FieldDescription) GetIsPrefixSearchable()(*bool) {
    return m.isPrefixSearchable
}
// GetIsSortable gets the isSortable property value. True if the field can be used for sorting.
// returns a *bool when successful
func (m *FieldDescription) GetIsSortable()(*bool) {
    return m.isSortable
}
// GetTypeEscaped gets the type property value. The data type of the field: identifier (text), content (text), text, integer, float, date, single-value taxonomy, multi-value taxonomy.   The type 'binary' is special in that it accepts and returns binary streams. A main use-case is to access Binary storages configured for a singleMindServer. To this end, use projectId=singleMindServer.NAME and collectionId=rm_storage:Native files for 'Native files' or collectionId=rm_storage:Image files to access image files. The syntax is 'rm_storage:' followed by the value configured as 'Storage file type' in the single mind server configuration.   The type 'hierarchy' allows to implement tree structures. Its value is a pointer to the nodes parent id. It is intended to be used in folder collections only.   The type 'parameterizedLong' is a multi-value type: a single document can have multiple values in such a field, and each value has a 'parameter' (a string) and a long value. Such fields are supposed to contain 'few' parameters, but many different long values. It is possible to insert/change values of the form '<parameter>.<long>' where '<parameter>' is some string chosen by the API user (example: rank.1234). Assuming such a field is called ranks, the associated values are returned by the records endpoint with fields=ranks.A, ranks.B where 'A' and 'B' are parameters. It is possible to search for documents with such values by means of field bases searches of the form 'ranks.A > 42' or 'ranks.B = 1234'. It is also possible to sort by these values: specify the sort field as 'ranks.A:desc' or 'ranks.B:asc'. Fields of type 'texts' contain a list of strings.
// returns a *FieldDescription_type when successful
func (m *FieldDescription) GetTypeEscaped()(*FieldDescription_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *FieldDescription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteBoolValue("isFolderField", m.GetIsFolderField())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMultivalueFolderCollection", m.GetIsMultivalueFolderCollection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPrefixSearchable", m.GetIsPrefixSearchable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSortable", m.GetIsSortable())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *FieldDescription) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A field display name. The value is never null, it falls back to the field id if there is no display name. Display names can be configured in the server configuration.
func (m *FieldDescription) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The field ID. The ID is the one to be used where the API requires a field ID (or folder field ID, depending on the context).
func (m *FieldDescription) SetId(value *string)() {
    m.id = value
}
// SetIsFolderField sets the isFolderField property value. True if the field is a folder field (taxonomy, smart filter). A document can be associated with a folder in such a folder field (or with multiple folders, depending on folder field). A folder field can be used for filtering. Values (e.g. the document count) can be accumulated for folders.
func (m *FieldDescription) SetIsFolderField(value *bool)() {
    m.isFolderField = value
}
// SetIsMultivalueFolderCollection sets the isMultivalueFolderCollection property value. True if the field is a multi-value folder collection (taxonomy). A folder collection can be used as filter. Each document can be associated with any number of folders of such a folder collection.
func (m *FieldDescription) SetIsMultivalueFolderCollection(value *bool)() {
    m.isMultivalueFolderCollection = value
}
// SetIsPrefixSearchable sets the isPrefixSearchable property value. True if a folder search on the field can be restricted by prefix.
func (m *FieldDescription) SetIsPrefixSearchable(value *bool)() {
    m.isPrefixSearchable = value
}
// SetIsSortable sets the isSortable property value. True if the field can be used for sorting.
func (m *FieldDescription) SetIsSortable(value *bool)() {
    m.isSortable = value
}
// SetTypeEscaped sets the type property value. The data type of the field: identifier (text), content (text), text, integer, float, date, single-value taxonomy, multi-value taxonomy.   The type 'binary' is special in that it accepts and returns binary streams. A main use-case is to access Binary storages configured for a singleMindServer. To this end, use projectId=singleMindServer.NAME and collectionId=rm_storage:Native files for 'Native files' or collectionId=rm_storage:Image files to access image files. The syntax is 'rm_storage:' followed by the value configured as 'Storage file type' in the single mind server configuration.   The type 'hierarchy' allows to implement tree structures. Its value is a pointer to the nodes parent id. It is intended to be used in folder collections only.   The type 'parameterizedLong' is a multi-value type: a single document can have multiple values in such a field, and each value has a 'parameter' (a string) and a long value. Such fields are supposed to contain 'few' parameters, but many different long values. It is possible to insert/change values of the form '<parameter>.<long>' where '<parameter>' is some string chosen by the API user (example: rank.1234). Assuming such a field is called ranks, the associated values are returned by the records endpoint with fields=ranks.A, ranks.B where 'A' and 'B' are parameters. It is possible to search for documents with such values by means of field bases searches of the form 'ranks.A > 42' or 'ranks.B = 1234'. It is also possible to sort by these values: specify the sort field as 'ranks.A:desc' or 'ranks.B:asc'. Fields of type 'texts' contain a list of strings.
func (m *FieldDescription) SetTypeEscaped(value *FieldDescription_type)() {
    m.typeEscaped = value
}
type FieldDescriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    GetIsFolderField()(*bool)
    GetIsMultivalueFolderCollection()(*bool)
    GetIsPrefixSearchable()(*bool)
    GetIsSortable()(*bool)
    GetTypeEscaped()(*FieldDescription_type)
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetIsFolderField(value *bool)()
    SetIsMultivalueFolderCollection(value *bool)()
    SetIsPrefixSearchable(value *bool)()
    SetIsSortable(value *bool)()
    SetTypeEscaped(value *FieldDescription_type)()
}
