package models
import (
    "errors"
)
// The ID of the resource. Typical resources available will be 'records' which provides access to the actual searchable objects, 'filters' (folder fields, taxonomies, smart filters) which can be used for filtering, etc.
type CollectionResource_id int

const (
    FIELDS_COLLECTIONRESOURCE_ID CollectionResource_id = iota
    FILTERS_COLLECTIONRESOURCE_ID
    RECORDS_COLLECTIONRESOURCE_ID
)

func (i CollectionResource_id) String() string {
    return []string{"fields", "filters", "records"}[i]
}
func ParseCollectionResource_id(v string) (any, error) {
    result := FIELDS_COLLECTIONRESOURCE_ID
    switch v {
        case "fields":
            result = FIELDS_COLLECTIONRESOURCE_ID
        case "filters":
            result = FILTERS_COLLECTIONRESOURCE_ID
        case "records":
            result = RECORDS_COLLECTIONRESOURCE_ID
        default:
            return 0, errors.New("Unknown CollectionResource_id value: " + v)
    }
    return &result, nil
}
func SerializeCollectionResource_id(values []CollectionResource_id) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CollectionResource_id) isMultiValue() bool {
    return false
}
