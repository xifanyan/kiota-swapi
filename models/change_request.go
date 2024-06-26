package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ChangeRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The field to change.
    field *string
    // Zero, one, or more folder ids. This value is evaluated for the following choices of 'type': 'ADD_FOLDERS', 'SET_FOLDERS', 'REMOVE_FOLDERS', 'REMOVE_FOLDERS_AND_BRANCH', 'SET_PARENT'. A folder id can be one of two possible formats: either it is the (unencoded) folder id as returned by other software components or it is the 'id' returned by the records endpoint when searching in the folder collection.
    folderIds []string
    // A string text. This value is evaluated for the following choices of 'type': 'SET_TEXT', 'APPEND_TEXT'. The type 'SET_TEXT' without a 'text' clears the field. Note that changes in a numeric field also use this parameter. If 'pages' is a configured numeric type, you can use field=pages, text='42', type='SET_TEXT'
    text *string
    // A sequence of strings. This value is evaluated for the following choices of 'type': 'SET_TEXTS'.
    texts []string
    // The type of change (case sensitive!). The semantics is as follows:  * ADD_FOLDERS tags all documents in scope into the folders specified in folderIds.  * SET_FOLDERS ensures that all documents in scope have exactly the folders specified in folderIds.  * REMOVE_FOLDERS untags all documents in scope from the folders specified in folderIds.  * REMOVE_FOLDERS_AND_BRANCH is similiar to removeFolders. However, for a hierarchical folder collection with enabled 'full hierarchy', parent associations will also be removed unless another child references the documents.  * MAKE_ROOT requires that the identified collection is a hierarchical folder collection (=taxonomy), and that field is set to 'rm_prop_parent'. In this case, it ensures that the documents in scope (which are folder ids) will become top-level folders (example: projectId='singleMindServer.project', collectionId='my_workspace', field='rm_prop_parent', type='MAKE_ROOT', recordId='unencoded category id').  * SET_PARENT has the same requirements as MAKE_ROOT, but it also expects exactly one argument in 'folderIds' (the new parent). It changes the hierarchy of the specified categories.  * SET_TEXT changes the content to the value specified in 'text' in the specified 'field' for all documents in scope. This is also possible for numeric fields or date fields. Note that date fields can be changed to a string resembling the milli seconds since 1970 suffixed with 'L' or a date pattern which must be configured in the singleMindServer configuration.  * APPEND_TEXT is similar to SET_TEXT, but it appends to the existing text (if any). This is impossible for numeric or date fields.  * SET_TEXTS can be used to modify folder properties by assigning a list of strings. It expects the value in 'texts'.
    typeEscaped *ChangeRequest_type
}
// NewChangeRequest instantiates a new ChangeRequest and sets the default values.
func NewChangeRequest()(*ChangeRequest) {
    m := &ChangeRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChangeRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChangeRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChangeRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ChangeRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetField gets the field property value. The field to change.
// returns a *string when successful
func (m *ChangeRequest) GetField()(*string) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChangeRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["folderIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetFolderIds(res)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    res["texts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTexts(res)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChangeRequest_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*ChangeRequest_type))
        }
        return nil
    }
    return res
}
// GetFolderIds gets the folderIds property value. Zero, one, or more folder ids. This value is evaluated for the following choices of 'type': 'ADD_FOLDERS', 'SET_FOLDERS', 'REMOVE_FOLDERS', 'REMOVE_FOLDERS_AND_BRANCH', 'SET_PARENT'. A folder id can be one of two possible formats: either it is the (unencoded) folder id as returned by other software components or it is the 'id' returned by the records endpoint when searching in the folder collection.
// returns a []string when successful
func (m *ChangeRequest) GetFolderIds()([]string) {
    return m.folderIds
}
// GetText gets the text property value. A string text. This value is evaluated for the following choices of 'type': 'SET_TEXT', 'APPEND_TEXT'. The type 'SET_TEXT' without a 'text' clears the field. Note that changes in a numeric field also use this parameter. If 'pages' is a configured numeric type, you can use field=pages, text='42', type='SET_TEXT'
// returns a *string when successful
func (m *ChangeRequest) GetText()(*string) {
    return m.text
}
// GetTexts gets the texts property value. A sequence of strings. This value is evaluated for the following choices of 'type': 'SET_TEXTS'.
// returns a []string when successful
func (m *ChangeRequest) GetTexts()([]string) {
    return m.texts
}
// GetTypeEscaped gets the type property value. The type of change (case sensitive!). The semantics is as follows:  * ADD_FOLDERS tags all documents in scope into the folders specified in folderIds.  * SET_FOLDERS ensures that all documents in scope have exactly the folders specified in folderIds.  * REMOVE_FOLDERS untags all documents in scope from the folders specified in folderIds.  * REMOVE_FOLDERS_AND_BRANCH is similiar to removeFolders. However, for a hierarchical folder collection with enabled 'full hierarchy', parent associations will also be removed unless another child references the documents.  * MAKE_ROOT requires that the identified collection is a hierarchical folder collection (=taxonomy), and that field is set to 'rm_prop_parent'. In this case, it ensures that the documents in scope (which are folder ids) will become top-level folders (example: projectId='singleMindServer.project', collectionId='my_workspace', field='rm_prop_parent', type='MAKE_ROOT', recordId='unencoded category id').  * SET_PARENT has the same requirements as MAKE_ROOT, but it also expects exactly one argument in 'folderIds' (the new parent). It changes the hierarchy of the specified categories.  * SET_TEXT changes the content to the value specified in 'text' in the specified 'field' for all documents in scope. This is also possible for numeric fields or date fields. Note that date fields can be changed to a string resembling the milli seconds since 1970 suffixed with 'L' or a date pattern which must be configured in the singleMindServer configuration.  * APPEND_TEXT is similar to SET_TEXT, but it appends to the existing text (if any). This is impossible for numeric or date fields.  * SET_TEXTS can be used to modify folder properties by assigning a list of strings. It expects the value in 'texts'.
// returns a *ChangeRequest_type when successful
func (m *ChangeRequest) GetTypeEscaped()(*ChangeRequest_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ChangeRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("field", m.GetField())
        if err != nil {
            return err
        }
    }
    if m.GetFolderIds() != nil {
        err := writer.WriteCollectionOfStringValues("folderIds", m.GetFolderIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    if m.GetTexts() != nil {
        err := writer.WriteCollectionOfStringValues("texts", m.GetTexts())
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
func (m *ChangeRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetField sets the field property value. The field to change.
func (m *ChangeRequest) SetField(value *string)() {
    m.field = value
}
// SetFolderIds sets the folderIds property value. Zero, one, or more folder ids. This value is evaluated for the following choices of 'type': 'ADD_FOLDERS', 'SET_FOLDERS', 'REMOVE_FOLDERS', 'REMOVE_FOLDERS_AND_BRANCH', 'SET_PARENT'. A folder id can be one of two possible formats: either it is the (unencoded) folder id as returned by other software components or it is the 'id' returned by the records endpoint when searching in the folder collection.
func (m *ChangeRequest) SetFolderIds(value []string)() {
    m.folderIds = value
}
// SetText sets the text property value. A string text. This value is evaluated for the following choices of 'type': 'SET_TEXT', 'APPEND_TEXT'. The type 'SET_TEXT' without a 'text' clears the field. Note that changes in a numeric field also use this parameter. If 'pages' is a configured numeric type, you can use field=pages, text='42', type='SET_TEXT'
func (m *ChangeRequest) SetText(value *string)() {
    m.text = value
}
// SetTexts sets the texts property value. A sequence of strings. This value is evaluated for the following choices of 'type': 'SET_TEXTS'.
func (m *ChangeRequest) SetTexts(value []string)() {
    m.texts = value
}
// SetTypeEscaped sets the type property value. The type of change (case sensitive!). The semantics is as follows:  * ADD_FOLDERS tags all documents in scope into the folders specified in folderIds.  * SET_FOLDERS ensures that all documents in scope have exactly the folders specified in folderIds.  * REMOVE_FOLDERS untags all documents in scope from the folders specified in folderIds.  * REMOVE_FOLDERS_AND_BRANCH is similiar to removeFolders. However, for a hierarchical folder collection with enabled 'full hierarchy', parent associations will also be removed unless another child references the documents.  * MAKE_ROOT requires that the identified collection is a hierarchical folder collection (=taxonomy), and that field is set to 'rm_prop_parent'. In this case, it ensures that the documents in scope (which are folder ids) will become top-level folders (example: projectId='singleMindServer.project', collectionId='my_workspace', field='rm_prop_parent', type='MAKE_ROOT', recordId='unencoded category id').  * SET_PARENT has the same requirements as MAKE_ROOT, but it also expects exactly one argument in 'folderIds' (the new parent). It changes the hierarchy of the specified categories.  * SET_TEXT changes the content to the value specified in 'text' in the specified 'field' for all documents in scope. This is also possible for numeric fields or date fields. Note that date fields can be changed to a string resembling the milli seconds since 1970 suffixed with 'L' or a date pattern which must be configured in the singleMindServer configuration.  * APPEND_TEXT is similar to SET_TEXT, but it appends to the existing text (if any). This is impossible for numeric or date fields.  * SET_TEXTS can be used to modify folder properties by assigning a list of strings. It expects the value in 'texts'.
func (m *ChangeRequest) SetTypeEscaped(value *ChangeRequest_type)() {
    m.typeEscaped = value
}
type ChangeRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetField()(*string)
    GetFolderIds()([]string)
    GetText()(*string)
    GetTexts()([]string)
    GetTypeEscaped()(*ChangeRequest_type)
    SetField(value *string)()
    SetFolderIds(value []string)()
    SetText(value *string)()
    SetTexts(value []string)()
    SetTypeEscaped(value *ChangeRequest_type)()
}
