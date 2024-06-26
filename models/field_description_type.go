package models
import (
    "errors"
)
// The data type of the field: identifier (text), content (text), text, integer, float, date, single-value taxonomy, multi-value taxonomy.   The type 'binary' is special in that it accepts and returns binary streams. A main use-case is to access Binary storages configured for a singleMindServer. To this end, use projectId=singleMindServer.NAME and collectionId=rm_storage:Native files for 'Native files' or collectionId=rm_storage:Image files to access image files. The syntax is 'rm_storage:' followed by the value configured as 'Storage file type' in the single mind server configuration.   The type 'hierarchy' allows to implement tree structures. Its value is a pointer to the nodes parent id. It is intended to be used in folder collections only.   The type 'parameterizedLong' is a multi-value type: a single document can have multiple values in such a field, and each value has a 'parameter' (a string) and a long value. Such fields are supposed to contain 'few' parameters, but many different long values. It is possible to insert/change values of the form '<parameter>.<long>' where '<parameter>' is some string chosen by the API user (example: rank.1234). Assuming such a field is called ranks, the associated values are returned by the records endpoint with fields=ranks.A, ranks.B where 'A' and 'B' are parameters. It is possible to search for documents with such values by means of field bases searches of the form 'ranks.A > 42' or 'ranks.B = 1234'. It is also possible to sort by these values: specify the sort field as 'ranks.A:desc' or 'ranks.B:asc'. Fields of type 'texts' contain a list of strings.
type FieldDescription_type int

const (
    IDENTIFIER_FIELDDESCRIPTION_TYPE FieldDescription_type = iota
    CONTENT_FIELDDESCRIPTION_TYPE
    TEXT_FIELDDESCRIPTION_TYPE
    INTEGER_FIELDDESCRIPTION_TYPE
    FLOAT_FIELDDESCRIPTION_TYPE
    DATE_FIELDDESCRIPTION_TYPE
    SINGLEVALUE_FIELDDESCRIPTION_TYPE
    MULTIVALUE_FIELDDESCRIPTION_TYPE
    BINARY_FIELDDESCRIPTION_TYPE
    HIERARCHY_FIELDDESCRIPTION_TYPE
    PARAMETERIZEDLONG_FIELDDESCRIPTION_TYPE
    TEXTS_FIELDDESCRIPTION_TYPE
)

func (i FieldDescription_type) String() string {
    return []string{"identifier", "content", "text", "integer", "float", "date", "singleValue", "multiValue", "binary", "hierarchy", "parameterizedLong", "texts"}[i]
}
func ParseFieldDescription_type(v string) (any, error) {
    result := IDENTIFIER_FIELDDESCRIPTION_TYPE
    switch v {
        case "identifier":
            result = IDENTIFIER_FIELDDESCRIPTION_TYPE
        case "content":
            result = CONTENT_FIELDDESCRIPTION_TYPE
        case "text":
            result = TEXT_FIELDDESCRIPTION_TYPE
        case "integer":
            result = INTEGER_FIELDDESCRIPTION_TYPE
        case "float":
            result = FLOAT_FIELDDESCRIPTION_TYPE
        case "date":
            result = DATE_FIELDDESCRIPTION_TYPE
        case "singleValue":
            result = SINGLEVALUE_FIELDDESCRIPTION_TYPE
        case "multiValue":
            result = MULTIVALUE_FIELDDESCRIPTION_TYPE
        case "binary":
            result = BINARY_FIELDDESCRIPTION_TYPE
        case "hierarchy":
            result = HIERARCHY_FIELDDESCRIPTION_TYPE
        case "parameterizedLong":
            result = PARAMETERIZEDLONG_FIELDDESCRIPTION_TYPE
        case "texts":
            result = TEXTS_FIELDDESCRIPTION_TYPE
        default:
            return 0, errors.New("Unknown FieldDescription_type value: " + v)
    }
    return &result, nil
}
func SerializeFieldDescription_type(values []FieldDescription_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FieldDescription_type) isMultiValue() bool {
    return false
}
