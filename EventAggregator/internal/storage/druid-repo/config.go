package druid

type Config struct {
	IngestURL  string `toml:"ingest_url"`
	Datasource string `toml:"datasource"`
}
