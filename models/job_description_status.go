package models
import (
    "errors"
)
type JobDescription_status int

const (
    RUNNING_JOBDESCRIPTION_STATUS JobDescription_status = iota
    STOPPED_JOBDESCRIPTION_STATUS
    STOPPING_JOBDESCRIPTION_STATUS
    PAUSED_JOBDESCRIPTION_STATUS
    PAUSING_JOBDESCRIPTION_STATUS
    FATAL_JOBDESCRIPTION_STATUS
    FINISHING_JOBDESCRIPTION_STATUS
    FINISHED_JOBDESCRIPTION_STATUS
    SCHEDULED_JOBDESCRIPTION_STATUS
    WAITING_JOBDESCRIPTION_STATUS
    CREATED_JOBDESCRIPTION_STATUS
)

func (i JobDescription_status) String() string {
    return []string{"RUNNING", "STOPPED", "STOPPING", "PAUSED", "PAUSING", "FATAL", "FINISHING", "FINISHED", "SCHEDULED", "WAITING", "CREATED"}[i]
}
func ParseJobDescription_status(v string) (any, error) {
    result := RUNNING_JOBDESCRIPTION_STATUS
    switch v {
        case "RUNNING":
            result = RUNNING_JOBDESCRIPTION_STATUS
        case "STOPPED":
            result = STOPPED_JOBDESCRIPTION_STATUS
        case "STOPPING":
            result = STOPPING_JOBDESCRIPTION_STATUS
        case "PAUSED":
            result = PAUSED_JOBDESCRIPTION_STATUS
        case "PAUSING":
            result = PAUSING_JOBDESCRIPTION_STATUS
        case "FATAL":
            result = FATAL_JOBDESCRIPTION_STATUS
        case "FINISHING":
            result = FINISHING_JOBDESCRIPTION_STATUS
        case "FINISHED":
            result = FINISHED_JOBDESCRIPTION_STATUS
        case "SCHEDULED":
            result = SCHEDULED_JOBDESCRIPTION_STATUS
        case "WAITING":
            result = WAITING_JOBDESCRIPTION_STATUS
        case "CREATED":
            result = CREATED_JOBDESCRIPTION_STATUS
        default:
            return 0, errors.New("Unknown JobDescription_status value: " + v)
    }
    return &result, nil
}
func SerializeJobDescription_status(values []JobDescription_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i JobDescription_status) isMultiValue() bool {
    return false
}
