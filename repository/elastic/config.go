package elastic

type Config struct {
	Addresses []string // A list of Elasticsearch nodes to use.
	Username  string   // Username for HTTP Basic Authentication.
	Password  string   // Password for HTTP Basic Authentication.

	Index string

	// PEM-encoded certificate authorities.
	// When set, an empty certificate pool will be created, and the certificates will be appended to it.
	// The option is only valid when the transport is not specified, or when it's http.Transport.
	CACert []byte

	RetryOnStatus        []int // List of status codes for retry. Default: 502, 503, 504.
	DisableRetry         bool  // Default: false.
	EnableRetryOnTimeout bool  // Default: false.
	MaxRetries           int   // Default: 3.

	EnableMetrics     bool // Enable the metrics collection.
	EnableDebugLogger bool // Enable the debug logging.
}

func DefaultConfig() Config {
	return Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "",
		Password:  "",
		Index:     "f1",
	}
}
