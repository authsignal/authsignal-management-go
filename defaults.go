package authsignal

import "encoding/json"

// Defaults maps a setting's JSON field name to the value the setting takes when it is blank. A
// setting with no entry has no reported default, and Defaults is nil when the API reports none.
type Defaults map[string]json.RawMessage
