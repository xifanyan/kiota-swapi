package models
import (
    "errors"
)
// The ID of the resource.
type ProjectResource_id int

const (
    COLLECTIONS_PROJECTRESOURCE_ID ProjectResource_id = iota
)

func (i ProjectResource_id) String() string {
    return []string{"collections"}[i]
}
func ParseProjectResource_id(v string) (any, error) {
    result := COLLECTIONS_PROJECTRESOURCE_ID
    switch v {
        case "collections":
            result = COLLECTIONS_PROJECTRESOURCE_ID
        default:
            return 0, errors.New("Unknown ProjectResource_id value: " + v)
    }
    return &result, nil
}
func SerializeProjectResource_id(values []ProjectResource_id) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProjectResource_id) isMultiValue() bool {
    return false
}
