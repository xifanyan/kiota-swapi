package models
import (
    "errors"
)
// The type of change (case sensitive!). The semantics is as follows:  * ADD_FOLDERS tags all documents in scope into the folders specified in folderIds.  * SET_FOLDERS ensures that all documents in scope have exactly the folders specified in folderIds.  * REMOVE_FOLDERS untags all documents in scope from the folders specified in folderIds.  * REMOVE_FOLDERS_AND_BRANCH is similiar to removeFolders. However, for a hierarchical folder collection with enabled 'full hierarchy', parent associations will also be removed unless another child references the documents.  * MAKE_ROOT requires that the identified collection is a hierarchical folder collection (=taxonomy), and that field is set to 'rm_prop_parent'. In this case, it ensures that the documents in scope (which are folder ids) will become top-level folders (example: projectId='singleMindServer.project', collectionId='my_workspace', field='rm_prop_parent', type='MAKE_ROOT', recordId='unencoded category id').  * SET_PARENT has the same requirements as MAKE_ROOT, but it also expects exactly one argument in 'folderIds' (the new parent). It changes the hierarchy of the specified categories.  * SET_TEXT changes the content to the value specified in 'text' in the specified 'field' for all documents in scope. This is also possible for numeric fields or date fields. Note that date fields can be changed to a string resembling the milli seconds since 1970 suffixed with 'L' or a date pattern which must be configured in the singleMindServer configuration.  * APPEND_TEXT is similar to SET_TEXT, but it appends to the existing text (if any). This is impossible for numeric or date fields.  * SET_TEXTS can be used to modify folder properties by assigning a list of strings. It expects the value in 'texts'.
type ChangeRequest_type int

const (
    ADD_FOLDERS_CHANGEREQUEST_TYPE ChangeRequest_type = iota
    SET_FOLDERS_CHANGEREQUEST_TYPE
    REMOVE_FOLDERS_CHANGEREQUEST_TYPE
    REMOVE_FOLDERS_AND_BRANCH_CHANGEREQUEST_TYPE
    MAKE_ROOT_CHANGEREQUEST_TYPE
    SET_PARENT_CHANGEREQUEST_TYPE
    SET_TEXT_CHANGEREQUEST_TYPE
    SET_TEXTS_CHANGEREQUEST_TYPE
    APPEND_TEXT_CHANGEREQUEST_TYPE
)

func (i ChangeRequest_type) String() string {
    return []string{"ADD_FOLDERS", "SET_FOLDERS", "REMOVE_FOLDERS", "REMOVE_FOLDERS_AND_BRANCH", "MAKE_ROOT", "SET_PARENT", "SET_TEXT", "SET_TEXTS", "APPEND_TEXT"}[i]
}
func ParseChangeRequest_type(v string) (any, error) {
    result := ADD_FOLDERS_CHANGEREQUEST_TYPE
    switch v {
        case "ADD_FOLDERS":
            result = ADD_FOLDERS_CHANGEREQUEST_TYPE
        case "SET_FOLDERS":
            result = SET_FOLDERS_CHANGEREQUEST_TYPE
        case "REMOVE_FOLDERS":
            result = REMOVE_FOLDERS_CHANGEREQUEST_TYPE
        case "REMOVE_FOLDERS_AND_BRANCH":
            result = REMOVE_FOLDERS_AND_BRANCH_CHANGEREQUEST_TYPE
        case "MAKE_ROOT":
            result = MAKE_ROOT_CHANGEREQUEST_TYPE
        case "SET_PARENT":
            result = SET_PARENT_CHANGEREQUEST_TYPE
        case "SET_TEXT":
            result = SET_TEXT_CHANGEREQUEST_TYPE
        case "SET_TEXTS":
            result = SET_TEXTS_CHANGEREQUEST_TYPE
        case "APPEND_TEXT":
            result = APPEND_TEXT_CHANGEREQUEST_TYPE
        default:
            return 0, errors.New("Unknown ChangeRequest_type value: " + v)
    }
    return &result, nil
}
func SerializeChangeRequest_type(values []ChangeRequest_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ChangeRequest_type) isMultiValue() bool {
    return false
}
