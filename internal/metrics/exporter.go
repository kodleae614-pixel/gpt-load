package metrics

import (
    "encoding/json"
    "net/http"
    "strings"
)

// Metrics represents the structure for holding metric data
 type Metrics struct {
    RequestCount int `json:"request_count"`
    SuccessCount int `json:"success_count"`
    ErrorCount   int `json:"error_count"`
}

// ExportJSON exports metrics in JSON format
func (m *Metrics) ExportJSON(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)
}

// ExportOpenMetrics exports metrics in OpenMetrics format
func (m *Metrics) ExportOpenMetrics(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/openmetrics-text")
    sb := strings.Builder{}
    sb.WriteString("# HELP request_count Total number of requests\n")
    sb.WriteString("# TYPE request_count counter\n")
    sb.WriteString("request_count " + strconv.Itoa(m.RequestCount) + "\n")
    sb.WriteString("# HELP success_count Total number of successful requests\n")
    sb.WriteString("# TYPE success_count counter\n")
    sb.WriteString("success_count " + strconv.Itoa(m.SuccessCount) + "\n")
    sb.WriteString("# HELP error_count Total number of errors\n")
    sb.WriteString("# TYPE error_count counter\n")
    sb.WriteString("error_count " + strconv.Itoa(m.ErrorCount) + "\n")
    
    w.Write([]byte(sb.String()))
}