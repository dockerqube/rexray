package types

// PathConfig contains the path configuration for the application.
type PathConfig struct {

	// Home is the path to the root data directory.
	Home string

	// Etc is the path to the etc directory.
	Etc string

	// Lib is the path to the lib directory.
	Lib string

	// Log is the path to the log directory.
	Log string

	// Run is the path to the run directory.
	Run string

	// TLS is the path to the tls directory.
	TLS string

	// LSX is the path to the executor.
	LSX string

	// DefaultTLSCertFile is the path to the default TLS cert file.
	DefaultTLSCertFile string

	// DefaultTLSKeyFile is the path to the default TLS key file.
	DefaultTLSKeyFile string

	// DefaultTLSTrustedRootsFile is the path to the default TLS trusted roots
	// file.
	DefaultTLSTrustedRootsFile string

	// DefaultTLSKnownHosts is the default path to the TLS known hosts file.
	DefaultTLSKnownHosts string
}
