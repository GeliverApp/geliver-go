package geliver

// Version is the current SDK version.
// Note: Go modules are versioned via git tags; keep this in sync with releases.
const Version = "1.0.0"

// DefaultUserAgent is sent with every request unless overridden on Client.UserAgent.
const DefaultUserAgent = "geliver-go/" + Version
