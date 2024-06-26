package models
import (
    "errors"
)
// The status of the backend response if there is one. This can be 'null'. The value 'meaningless search terms' indicates that the query consists only of stop words and the-like and results in HTTP code 200. The value 'all search terms unknown' is similiar, but the server does not know any of the provided terms. The special value 'incomplete search results due to pending synchronization' can only occur in early startup phases of the target server (if synchronizes searchable folder properties asynchronously); it results in HTTP code 200. The value 'child collection invalid' means that a sub-engine is currently unreachable and its results are not part of the result and also results in code 200. The value 'unexpandable wildcard' indicates that a wildcard prefix is too short (it results in HTTP code 400). The value 'expanded query too long' indicates that a wildcard expression resulted in too many terms (which is also code 400). The values here are constants and are subject to api stability requirements. The value 'syntax error' indicates a syntax error in a query. The value 'unsupported expression' indicates that the query as such is valid, but cannot be used in the current context.
type StatusObject_backendStatus int

const (
    OK_STATUSOBJECT_BACKENDSTATUS StatusObject_backendStatus = iota
    MEANINGLESSSEARCHTERMS_STATUSOBJECT_BACKENDSTATUS
    UNKNOWNERROR_STATUSOBJECT_BACKENDSTATUS
    ALLSEARCHTERMSUNKNOWN_STATUSOBJECT_BACKENDSTATUS
    CHILDCOLLECTIONINVALID_STATUSOBJECT_BACKENDSTATUS
    INCOMPLETESEARCHRESULTSDUETOPENDINGSYNCHRONIZATION_STATUSOBJECT_BACKENDSTATUS
    UNEXPANDABLEWILDCARD_STATUSOBJECT_BACKENDSTATUS
    EXPANDEDQUERYTOOLONG_STATUSOBJECT_BACKENDSTATUS
    SYNTAXERROR_STATUSOBJECT_BACKENDSTATUS
    UNSUPPORTEDEXPRESSION_STATUSOBJECT_BACKENDSTATUS
)

func (i StatusObject_backendStatus) String() string {
    return []string{"ok", "meaningless search terms", "unknown error", "all search terms unknown", "child collection invalid", "incomplete search results due to pending synchronization", "unexpandable wildcard", "expanded query too long", "syntax error", "unsupported expression"}[i]
}
func ParseStatusObject_backendStatus(v string) (any, error) {
    result := OK_STATUSOBJECT_BACKENDSTATUS
    switch v {
        case "ok":
            result = OK_STATUSOBJECT_BACKENDSTATUS
        case "meaningless search terms":
            result = MEANINGLESSSEARCHTERMS_STATUSOBJECT_BACKENDSTATUS
        case "unknown error":
            result = UNKNOWNERROR_STATUSOBJECT_BACKENDSTATUS
        case "all search terms unknown":
            result = ALLSEARCHTERMSUNKNOWN_STATUSOBJECT_BACKENDSTATUS
        case "child collection invalid":
            result = CHILDCOLLECTIONINVALID_STATUSOBJECT_BACKENDSTATUS
        case "incomplete search results due to pending synchronization":
            result = INCOMPLETESEARCHRESULTSDUETOPENDINGSYNCHRONIZATION_STATUSOBJECT_BACKENDSTATUS
        case "unexpandable wildcard":
            result = UNEXPANDABLEWILDCARD_STATUSOBJECT_BACKENDSTATUS
        case "expanded query too long":
            result = EXPANDEDQUERYTOOLONG_STATUSOBJECT_BACKENDSTATUS
        case "syntax error":
            result = SYNTAXERROR_STATUSOBJECT_BACKENDSTATUS
        case "unsupported expression":
            result = UNSUPPORTEDEXPRESSION_STATUSOBJECT_BACKENDSTATUS
        default:
            return 0, errors.New("Unknown StatusObject_backendStatus value: " + v)
    }
    return &result, nil
}
func SerializeStatusObject_backendStatus(values []StatusObject_backendStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StatusObject_backendStatus) isMultiValue() bool {
    return false
}
