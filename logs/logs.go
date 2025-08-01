package logs

// Types based on OpenTelemetry's spec for the logs produced by the file exporter
type LogFile struct {
	ResourceLogs []struct {
		Resource struct {
			Attributes []Attribute `json:"attributes"`
		} `json:"resource"`
		ScopeLogs []struct {
			LogRecords []LogRecord `json:"logRecords"`
		} `json:"scopeLogs"`
	} `json:"resourceLogs"`
}

type LogRecord struct {
	TimeUnixNano           string      `json:"timeUnixNano"`
	SeverityNumber         int         `json:"severityNumber"`
	SeverityText           string      `json:"severityText"`
	Body                   ValueField  `json:"body"`
	Attributes             []Attribute `json:"attributes"`
	DroppedAttributesCount int         `json:"droppedAttributesCount,omitempty"` // This doesn't parse the values from OTLP correctly, instead showing a conversion error or a very large number
	TraceID                string      `json:"traceId,omitempty"`
	SpanID                 string      `json:"spanId,omitempty"`
}

type Attribute struct {
	Key   string     `json:"key"`
	Value ValueField `json:"value"`
}

type ValueField struct {
	// Assume just strings for now
	StringValue string `json:"stringValue,omitempty"`
}
