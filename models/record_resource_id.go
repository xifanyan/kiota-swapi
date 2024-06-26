package models
import (
    "errors"
)
// The ID of the resource.
type RecordResource_id int

const (
    CONTENT_RECORDRESOURCE_ID RecordResource_id = iota
)

func (i RecordResource_id) String() string {
    return []string{"content"}[i]
}
func ParseRecordResource_id(v string) (any, error) {
    result := CONTENT_RECORDRESOURCE_ID
    switch v {
        case "content":
            result = CONTENT_RECORDRESOURCE_ID
        default:
            return 0, errors.New("Unknown RecordResource_id value: " + v)
    }
    return &result, nil
}
func SerializeRecordResource_id(values []RecordResource_id) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RecordResource_id) isMultiValue() bool {
    return false
}
