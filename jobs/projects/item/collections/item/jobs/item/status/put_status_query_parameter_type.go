package status
import (
    "errors"
)
type PutStatusQueryParameterType int

const (
    START_PUTSTATUSQUERYPARAMETERTYPE PutStatusQueryParameterType = iota
    PAUSE_PUTSTATUSQUERYPARAMETERTYPE
    STOP_PUTSTATUSQUERYPARAMETERTYPE
)

func (i PutStatusQueryParameterType) String() string {
    return []string{"START", "PAUSE", "STOP"}[i]
}
func ParsePutStatusQueryParameterType(v string) (any, error) {
    result := START_PUTSTATUSQUERYPARAMETERTYPE
    switch v {
        case "START":
            result = START_PUTSTATUSQUERYPARAMETERTYPE
        case "PAUSE":
            result = PAUSE_PUTSTATUSQUERYPARAMETERTYPE
        case "STOP":
            result = STOP_PUTSTATUSQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown PutStatusQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializePutStatusQueryParameterType(values []PutStatusQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PutStatusQueryParameterType) isMultiValue() bool {
    return false
}
