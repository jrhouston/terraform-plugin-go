package tfprotov5

import "context"

// StorageServer is an interface containing the methods a storage
// implementation needs to fill.
type StorageServer interface {
	// ValidateStorageConfig is called when Terraform to check that a
	// storage configuration is valid.
	ValidateStorageConfig(context.Context, *ValidateStorageConfigRequest) (*ValidateStorageConfigResponse, error)

	// ConfigureStorage is called to pass the user-specified storage
	// configuration to the provider.
	ConfigureStorage(context.Context, *ConfigureStorageRequest) (*ConfigureStorageResponse, error)

	// ReadState is called to lock a particular state.
	ReadState(context.Context, *ReadStateRequest) (*ReadStateResponse, error)

	// WriteState is called to unlock a particular state.
	WriteState(context.Context, *WriteStateRequest) (*WriteStateResponse, error)

	// LockState is called to lock a particular state.
	LockState(context.Context, *LockStateRequest) (*LockStateResponse, error)

	// UnlockState is called to unlock a particular state.
	UnlockState(context.Context, *UnlockStateRequest) (*UnlockStateResponse, error)

	// GetStates is called to get a list of the states managed by the configured storage implementation.
	GetStates(context.Context, *GetStatesRequest) (*GetStatesResponse, error)

	// DeleteState is called to delete a state that is being managed by the provider.
	DeleteState(context.Context, *DeleteStateRequest) (*DeleteStateResponse, error)
}

// ValidateStorageConfigRequest is the request Terraform sends when it
// wants to validate a storage configuration.
type ValidateStorageConfigRequest struct {
	// StorageName is the type of storage implementation
	StorageName string

	// Config is the configuration the user supplied for that storage implementation.
	// Seethe documentation on `DynamicValue` for more information about
	// safely accessing the configuration.
	//
	// The configuration is represented as a tftypes.Object, with each
	// attribute and nested block getting its own key and value.
	//
	// This configuration may contain unknown values if a user uses
	// interpolation or other functionality that would prevent Terraform
	// from knowing the value at request time. Any attributes not directly
	// set in the configuration will be null.
	Config *DynamicValue
}

// ValidateStorageConfigResponse is the response from the provider about
// the validity of a resource's configuration.
type ValidateStorageConfigResponse struct {
	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}

// ConfigureStorageRequest represents a Terraform RPC request to supply the
// provider with information about what the user entered in the storage
// configuration
type ConfigureStorageRequest struct {
	// StorageName is the name of the storage implementation to use
	StorageName string

	// Config is the configuration the user supplied for the storage. This
	// information should usually be persisted to the underlying type
	// that's implementing the ProviderServer interface, for use in later
	// RPC requests. See the documentation on `DynamicValue` for more
	// information about safely accessing the configuration.
	//
	// The configuration is represented as a tftypes.Object, with each
	// attribute and nested block getting its own key and value.
	Config *DynamicValue
}

// ConfigureStorageResponse represents a Terraform RPC response to the
// configuration block that Terraform supplied for the storage.
type ConfigureStorageResponse struct {
	// Diagnostics report errors or warnings related to the storage
	// configuration. Returning an empty slice indicates success, with no
	// errors or warnings generated.
	Diagnostics []*Diagnostic
}

// ReadStateRequest is the request Terraform sends when it
// wants to read a state from storage
type ReadStateRequest struct {
	// StateName is the name of the state to be unlocked
	StateName string
}

// ReadStateResponse is the response from the provider to the
// the request to read a state from storage
type ReadStateResponse struct {
	// State is the raw state read from the storage
	State []byte

	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}

// WriteStateRequest is the request Terraform sends when it
// wants to write a state to storage
type WriteStateRequest struct {
	// StateName is the name of the state to be unlocked
	StateName string

	// State is the raw state to write to the storage
	State []byte
}

// WriteStateResponse is the response from the provider to the
// the request to write a state to storage
type WriteStateResponse struct {
	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}

// LockStateRequest is the request Terraform sends to lock
// a particular state
type LockStateRequest struct {
	// StateName is the name of the state to be locked
	StateName string
}

// LockStateResponse is the response from the provider to the
// the request to lock the state
type LockStateResponse struct {
	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}

// UnlockStateRequest is the request Terraform sends when it
// wants to validate a storage configuration.
type UnlockStateRequest struct {
	// StateName is the name of the state to be unlocked
	StateName string

	// Force tells the provider to forcibly unlock the state
	Force bool
}

// UnlockStateResponse is the response from the provider to the
// the request to unlock the state
type UnlockStateResponse struct {
	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}

// GetStatesRequest is the request Terraform sends when it
// wants to get a list of states managed by the configured storage
type GetStatesRequest struct {
}

// GetStatesResponse is the response from the provider with the
// the list of states it is managing
type GetStatesResponse struct {
	StateNames []string

	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}

// DeleteStateRequest is the request Terraform sends when it
// wants to get a list of states managed by the configured storage
type DeleteStateRequest struct {
	// StateName is the name of the state to be deleted
	StateName string
}

// DeleteStateResponse is the response from the provider with the
// the list of states it is managing
type DeleteStateResponse struct {
	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}
