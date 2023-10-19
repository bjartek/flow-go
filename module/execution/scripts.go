package execution

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/engine/execution/computation"
	"github.com/onflow/flow-go/engine/execution/computation/query"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/storage/derived"
	"github.com/onflow/flow-go/fvm/storage/snapshot"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/storage"
)

// ErrDataNotAvailable indicates that the data for a given block was not available
//
// This generally indicates a request was made for execution data at a block height that was not
// not locally indexed
var ErrDataNotAvailable = errors.New("data for block is not available")

// RegistersAtHeight returns register value for provided register ID at the block height.
// Even if the register wasn't indexed at the provided height, returns the highest height the register was indexed at.
// Expected errors:
// - storage.ErrNotFound if the register by the ID was never indexed
// - storage.ErrHeightNotIndexed if the given height was not indexed yet or lower than the first indexed height.
type RegistersAtHeight func(ID flow.RegisterID, height uint64) (flow.RegisterValue, error)

type ScriptExecutor interface {
	// ExecuteAtBlockHeight executes provided script against the block height.
	// A result value is returned encoded as byte array. An error will be returned if script
	// doesn't successfully execute.
	// Expected errors:
	// - storage.ErrNotFound if block or register value at height was not found.
	// - ErrDataNotAvailable if the data for the block height is not available
	ExecuteAtBlockHeight(
		ctx context.Context,
		script []byte,
		arguments [][]byte,
		height uint64,
	) ([]byte, error)

	// GetAccountAtBlockHeight returns a Flow account by the provided address and block height.
	// Expected errors:
	// - storage.ErrNotFound if block or register value at height was not found.
	// - ErrDataNotAvailable if the data for the block height is not available
	GetAccountAtBlockHeight(ctx context.Context, address flow.Address, height uint64) (*flow.Account, error)
}

var _ ScriptExecutor = (*Scripts)(nil)

type Scripts struct {
	executor          *query.QueryExecutor
	headers           storage.Headers
	registersAtHeight RegistersAtHeight
}

func NewScripts(
	log zerolog.Logger,
	metrics *metrics.ExecutionCollector,
	chainID flow.ChainID,
	entropy query.EntropyProviderPerBlock,
	header storage.Headers,
	registersAtHeight RegistersAtHeight,
) (*Scripts, error) {
	vm := fvm.NewVirtualMachine()

	options := computation.DefaultFVMOptions(chainID, false, false)
	vmCtx := fvm.NewContext(options...)

	derivedChainData, err := derived.NewDerivedChainData(derived.DefaultDerivedDataCacheSize)
	if err != nil {
		return nil, err
	}

	queryExecutor := query.NewQueryExecutor(
		query.NewDefaultConfig(),
		log,
		metrics,
		vm,
		vmCtx,
		derivedChainData,
		entropy,
	)

	return &Scripts{
		executor:          queryExecutor,
		headers:           header,
		registersAtHeight: registersAtHeight,
	}, nil
}

// ExecuteAtBlockHeight executes provided script against the block height.
// A result value is returned encoded as byte array. An error will be returned if script
// doesn't successfully execute.
// Expected errors:
// - storage.ErrNotFound if block or register value at height was not found.
// - ErrDataNotAvailable if the data for the block height is not available
func (s *Scripts) ExecuteAtBlockHeight(
	ctx context.Context,
	script []byte,
	arguments [][]byte,
	height uint64,
) ([]byte, error) {

	snap, header, err := s.snapshotWithBlock(height)
	if err != nil {
		return nil, err
	}

	return s.executor.ExecuteScript(ctx, script, arguments, header, snap)
}

// GetAccountAtBlockHeight returns a Flow account by the provided address and block height.
// Expected errors:
// - storage.ErrNotFound if block or register value at height was not found.
// - ErrDataNotAvailable if the data for the block height is not available
func (s *Scripts) GetAccountAtBlockHeight(ctx context.Context, address flow.Address, height uint64) (*flow.Account, error) {
	snap, header, err := s.snapshotWithBlock(height)
	if err != nil {
		return nil, err
	}

	return s.executor.GetAccount(ctx, address, header, snap)
}

// snapshotWithBlock is a common function for executing scripts and get account functionality.
// It creates a storage snapshot that is needed by the FVM to execute scripts.
func (s *Scripts) snapshotWithBlock(height uint64) (snapshot.StorageSnapshot, *flow.Header, error) {
	header, err := s.headers.ByHeight(height)
	if err != nil {
		return nil, nil, err
	}

	storageSnapshot := snapshot.NewReadFuncStorageSnapshot(func(ID flow.RegisterID) (flow.RegisterValue, error) {
		registers, err := s.registersAtHeight(ID, height)
		if errors.Is(err, storage.ErrHeightNotIndexed) {
			return nil, errors.Join(ErrDataNotAvailable, err)
		}
		return registers, err
	})

	return storageSnapshot, header, nil
}

// IndexRegisterAdapter an adapter for using indexer register values function that takes a slice of IDs in the
// script executor that only uses a single register ID at a time. It also does additional sanity checks if multiple values
// are returned, which shouldn't occur in normal operation.
func IndexRegisterAdapter(registerFun func(IDs flow.RegisterIDs, height uint64) ([]flow.RegisterValue, error)) func(flow.RegisterID, uint64) (flow.RegisterValue, error) {
	return func(ID flow.RegisterID, height uint64) (flow.RegisterValue, error) {
		values, err := registerFun([]flow.RegisterID{ID}, height)
		if err != nil {
			return nil, err
		}

		// even though this shouldn't occur in correct implementation we check that function returned either a single register or error
		if len(values) != 1 {
			return nil, fmt.Errorf("invalid number of returned values for a single register: %d", len(values))
		}
		return values[0], nil
	}
}
