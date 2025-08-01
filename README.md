## Use Case
This is a simple program to parse `.jsonl` files containing [OTLP](https://opentelemetry.io/docs/specs/otlp/)-formatted telemetry data outputted by the [OpenTelemetry Collector's File Exporter](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/protocol/file-exporter.md) into flat SQLite files.

This may be useful in those situations where a program emitting (potentially sensitive) telemetry data resides on an airgapped (or otherwise unable to communicate this telemetry data externally) machine. The telemetry can be written to a file rather than streamed over gRPC/HTTP as is usually done.
In such situations, the telemetry within the file may need to be read and analysed at a later date. There is currently no good way to do this. Using this tool, the file can be serialized into a standardised database format to make it more usable.
Once in SQLite format, you could (for example) use something like the [Grafana SQLite Plugin](https://grafana.com/grafana/plugins/frser-sqlite-datasource/) to view telemetry data on a dashboard.

## Things to Note
- This project is still in its early stages. Currently it is an MVP which supports the features required for my use case (parsing of traces and logs). I am continuing to work on and improve it, to take into account other cases.
- Metrics are still a work in progress. Their structure as expected by Grafana is less clear than that of traces and logs.
