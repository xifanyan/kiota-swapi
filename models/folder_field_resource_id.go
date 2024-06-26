package models
import (
    "errors"
)
// The ID of the resource.
type FolderFieldResource_id int

const (
    VALUES_FOLDERFIELDRESOURCE_ID FolderFieldResource_id = iota
)

func (i FolderFieldResource_id) String() string {
    return []string{"values"}[i]
}
func ParseFolderFieldResource_id(v string) (any, error) {
    result := VALUES_FOLDERFIELDRESOURCE_ID
    switch v {
        case "values":
            result = VALUES_FOLDERFIELDRESOURCE_ID
        default:
            return 0, errors.New("Unknown FolderFieldResource_id value: " + v)
    }
    return &result, nil
}
func SerializeFolderFieldResource_id(values []FolderFieldResource_id) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FolderFieldResource_id) isMultiValue() bool {
    return false
}
