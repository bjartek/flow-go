package errors

import (
	"fmt"
	"strings"

	"github.com/onflow/cadence/runtime"

	"github.com/onflow/flow-go/model/flow"
)

// CadenceRuntimeError captures a collection of errors provided by cadence runtime
// it cover cadence errors such as
// NotDeclaredError, NotInvokableError, ArgumentCountError, TransactionNotDeclaredError,
// ConditionError, RedeclarationError, DereferenceError,
// OverflowError, UnderflowError, DivisionByZeroError,
// DestroyedCompositeError,  ForceAssignmentToNonNilResourceError, ForceNilError,
// TypeMismatchError, InvalidPathDomainError, OverwriteError, CyclicLinkError,
// ArrayIndexOutOfBoundsError, ...
// The Cadence error might have occurred because of an inner fvm Error.
type CadenceRuntimeError struct {
	errorWrapper
}

// NewCadenceRuntimeError constructs a new CadenceRuntimeError and wraps a cadence runtime error
func NewCadenceRuntimeError(err runtime.Error) CadenceRuntimeError {
	return CadenceRuntimeError{
		errorWrapper: errorWrapper{
			err: err,
		},
	}
}

func (e CadenceRuntimeError) Error() string {
	return fmt.Sprintf("%s cadence runtime error %s", e.Code().String(), e.err.Error())
}

// Code returns the error code for this error
func (e CadenceRuntimeError) Code() ErrorCode {
	return ErrCodeCadenceRunTimeError
}

func IsCadenceRuntimeError(err error) bool {
	var t CadenceRuntimeError
	return As(err, &t)
}

// NewTransactionFeeDeductionFailedError constructs a new CodedError which
// indicates that a there was an error deducting transaction fees from the
// transaction Payer
func NewTransactionFeeDeductionFailedError(
	payer flow.Address,
	txFees uint64,
	err error,
) *CodedError {
	return WrapCodedError(
		ErrCodeTransactionFeeDeductionFailedError,
		err,
		"failed to deduct %d transaction fees from %s",
		txFees,
		payer)
}

// NewComputationLimitExceededError constructs a new CodedError which indicates
// that computation has exceeded its limit.
func NewComputationLimitExceededError(limit uint64) *CodedError {
	return NewCodedError(
		ErrCodeComputationLimitExceededError,
		"computation exceeds limit (%d)",
		limit)
}

// IsComputationLimitExceededError returns true if error has this code.
func IsComputationLimitExceededError(err error) bool {
	return HasErrorCode(err, ErrCodeComputationLimitExceededError)
}

// NewMemoryLimitExceededError constructs a new CodedError which indicates
// that execution has exceeded its memory limits.
func NewMemoryLimitExceededError(limit uint64) *CodedError {
	return NewCodedError(
		ErrCodeMemoryLimitExceededError,
		"memory usage exceeds limit (%d)",
		limit)
}

// IsMemoryLimitExceededError returns true if error has this code.
func IsMemoryLimitExceededError(err error) bool {
	return HasErrorCode(err, ErrCodeMemoryLimitExceededError)
}

// NewStorageCapacityExceededError constructs a new CodedError which indicates
// that an account used more storage than it has storage capacity.
func NewStorageCapacityExceededError(
	address flow.Address,
	storageUsed uint64,
	storageCapacity uint64,
) *CodedError {
	return NewCodedError(
		ErrCodeStorageCapacityExceeded,
		"The account with address (%s) uses %d bytes of storage which is "+
			"over its capacity (%d bytes). Capacity can be increased by "+
			"adding FLOW tokens to the account.",
		address,
		storageUsed,
		storageCapacity)
}

func IsStorageCapacityExceededError(err error) bool {
	return HasErrorCode(err, ErrCodeStorageCapacityExceeded)
}

// NewEventLimitExceededError constructs a CodedError which indicates that the
// transaction has produced events with size more than limit.
func NewEventLimitExceededError(
	totalByteSize uint64,
	limit uint64) *CodedError {
	return NewCodedError(
		ErrCodeEventLimitExceededError,
		"total event byte size (%d) exceeds limit (%d)",
		totalByteSize,
		limit)
}

// NewStateKeySizeLimitError constructs a CodedError which indicates that the
// provided key has exceeded the size limit allowed by the storage.
func NewStateKeySizeLimitError(
	owner string,
	key string,
	size uint64,
	limit uint64,
) *CodedError {
	return NewCodedError(
		ErrCodeStateKeySizeLimitError,
		"key %s has size %d which is higher than storage key size limit %d.",
		strings.Join([]string{owner, key}, "/"),
		size,
		limit)
}

// NewStateValueSizeLimitError constructs a CodedError which indicates that the
// provided value has exceeded the size limit allowed by the storage.
func NewStateValueSizeLimitError(
	value flow.RegisterValue,
	size uint64,
	limit uint64,
) *CodedError {
	valueStr := ""
	if len(value) > 23 {
		valueStr = string(value[0:10]) + "..." + string(value[len(value)-10:])
	} else {
		valueStr = string(value)
	}

	return NewCodedError(
		ErrCodeStateValueSizeLimitError,
		"value %s has size %d which is higher than storage value size "+
			"limit %d.",
		valueStr,
		size,
		limit)
}

// NewLedgerInteractionLimitExceededError constructs a CodeError. It is
// returned when a tx hits the maximum ledger interaction limit.
func NewLedgerInteractionLimitExceededError(
	used uint64,
	limit uint64,
) *CodedError {
	return NewCodedError(
		ErrCodeLedgerInteractionLimitExceededError,
		"max interaction with storage has exceeded the limit "+
			"(used: %d bytes, limit %d bytes)",
		used,
		limit)
}

func IsLedgerInteractionLimitExceededError(err error) bool {
	return HasErrorCode(err, ErrCodeLedgerInteractionLimitExceededError)
}

// NewOperationNotSupportedError construct a new CodedError. It is generated
// when an operation (e.g. getting block info) is not supported in the current
// environment.
func NewOperationNotSupportedError(operation string) *CodedError {
	return NewCodedError(
		ErrCodeOperationNotSupportedError,
		"operation (%s) is not supported in this environment",
		operation)
}

func IsOperationNotSupportedError(err error) bool {
	return HasErrorCode(err, ErrCodeOperationNotSupportedError)
}

// NewScriptExecutionCancelledError construct a new CodedError which indicates
// that Cadence Script execution has been cancelled (e.g. request connection
// has been droped)
//
// note: this error is used by scripts only and won't be emitted for
// transactions since transaction execution has to be deterministic.
func NewScriptExecutionCancelledError(err error) *CodedError {
	return WrapCodedError(
		ErrCodeScriptExecutionCancelledError,
		err,
		"script execution is cancelled")
}

// NewScriptExecutionTimedOutError construct a new CodedError which indicates
// that Cadence Script execution has been taking more time than what is allowed.
//
// note: this error is used by scripts only and won't be emitted for
// transactions since transaction execution has to be deterministic.
func NewScriptExecutionTimedOutError() *CodedError {
	return NewCodedError(
		ErrCodeScriptExecutionTimedOutError,
		"script execution is timed out and did not finish executing within "+
			"the maximum execution time allowed")
}

// NewCouldNotGetExecutionParameterFromStateError constructs a new CodedError
// which indicates that computation has exceeded its limit.
func NewCouldNotGetExecutionParameterFromStateError(
	address string,
	path string,
) *CodedError {
	return NewCodedError(
		ErrCodeCouldNotDecodeExecutionParameterFromState,
		"could not get execution parameter from the state "+
			"(address: %s path: %s)",
		address,
		path)
}
