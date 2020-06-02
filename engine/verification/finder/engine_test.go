package finder_test

import (
	"testing"

	"github.com/rs/zerolog"
	testifymock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/dapperlabs/flow-go/engine"
	"github.com/dapperlabs/flow-go/engine/verification/finder"
	"github.com/dapperlabs/flow-go/engine/verification/utils"
	"github.com/dapperlabs/flow-go/model/flow"
	mempool "github.com/dapperlabs/flow-go/module/mempool/mock"
	module "github.com/dapperlabs/flow-go/module/mock"
	network "github.com/dapperlabs/flow-go/network/mock"
	storage "github.com/dapperlabs/flow-go/storage/mock"
	"github.com/dapperlabs/flow-go/utils/unittest"
)

// FinderEngineTestSuite contains the unit tests of Finder engine.
type FinderEngineTestSuite struct {
	suite.Suite
	net *module.Network
	me  *module.Local

	// mock conduit for receiving receipts
	receiptsConduit *network.Conduit

	// mock mempool for receipts
	receipts *mempool.Receipts

	// mock mempool for processed result IDs
	processedResults *mempool.Identifiers
	// mock mempool for header storage of blocks
	headerStorage *storage.Headers
	// resources fixtures
	collection    *flow.Collection
	block         *flow.Block
	receipt       *flow.ExecutionReceipt
	chunk         *flow.Chunk
	chunkDataPack *flow.ChunkDataPack

	// identities
	verIdentity  *flow.Identity // verification node
	execIdentity *flow.Identity // execution node

	// other engine
	// mock Match engine, should be called when Finder engine completely
	// processes a receipt
	matchEng *network.Engine
}

// TestFinderEngine executes all FinderEngineTestSuite tests.
func TestFinderEngine(t *testing.T) {
	suite.Run(t, new(FinderEngineTestSuite))
}

// SetupTest initiates the test setups prior to each test.
func (suite *FinderEngineTestSuite) SetupTest() {
	suite.receiptsConduit = &network.Conduit{}
	suite.net = &module.Network{}
	suite.me = &module.Local{}
	suite.headerStorage = &storage.Headers{}
	suite.receipts = &mempool.Receipts{}
	suite.processedResults = &mempool.Identifiers{}
	suite.matchEng = &network.Engine{}

	// generates an execution result with a single collection, chunk, and transaction.
	completeER := utils.LightExecutionResultFixture(1)
	suite.collection = completeER.Collections[0]
	suite.block = completeER.Block
	suite.receipt = completeER.Receipt
	suite.chunk = completeER.Receipt.ExecutionResult.Chunks[0]
	suite.chunkDataPack = completeER.ChunkDataPacks[0]

	suite.verIdentity = unittest.IdentityFixture(unittest.WithRole(flow.RoleVerification))
	suite.execIdentity = unittest.IdentityFixture(unittest.WithRole(flow.RoleExecution))

	// mocking the network registration of the engine
	suite.net.On("Register", uint8(engine.ExecutionReceiptProvider), testifymock.Anything).
		Return(suite.receiptsConduit, nil).
		Once()

	// mocks identity of the verification node
	suite.me.On("NodeID").Return(suite.verIdentity.NodeID)
}

// TestNewFinderEngine tests the establishment of the network registration upon
// creation of an instance of FinderEngine using the New method.
// It also returns an instance of new engine to be used in the later tests.
func (suite *FinderEngineTestSuite) TestNewFinderEngine() *finder.Engine {
	e, err := finder.New(zerolog.Logger{},
		suite.net,
		suite.me,
		suite.matchEng,
		suite.receipts,
		suite.headerStorage,
		suite.processedResults)
	require.Nil(suite.T(), err, "could not create finder engine")

	suite.net.AssertExpectations(suite.T())

	return e
}

// TestHandleReceipt_HappyPath evaluates that handling a receipt that is not duplicate,
// and its result has not been processed yet ends by:
// - sending its result to match engine.
// - marking its result as processed.
// - removing it from mempool.
func (suite *FinderEngineTestSuite) TestHandleReceipt_HappyPath() {
	e := suite.TestNewFinderEngine()

	// mocks result has not yet processed
	suite.processedResults.On("Has", suite.receipt.ExecutionResult.ID()).Return(false)

	// mocks adding receipt to the receipts mempool
	suite.receipts.On("Add", suite.receipt).Return(true).Once()

	// mocks block associated with receipt
	suite.headerStorage.On("ByBlockID", suite.block.ID()).Return(&flow.Header{}, nil).Once()

	// mocks successful submission to match engine
	suite.matchEng.On("ProcessLocal", &suite.receipt.ExecutionResult).Return(nil).Once()

	// mocks marking receipt as processed
	suite.processedResults.On("Add", suite.receipt.ExecutionResult.ID()).Return(true)

	// mocks receipt clean up after result is processed
	suite.receipts.On("All").Return([]*flow.ExecutionReceipt{suite.receipt})
	suite.receipts.On("Rem", suite.receipt.ID()).Return(true)

	// sends receipt to finder engine
	err := e.Process(suite.execIdentity.NodeID, suite.receipt)
	require.NoError(suite.T(), err)

	suite.receipts.AssertExpectations(suite.T())
	suite.headerStorage.AssertExpectations(suite.T())
	suite.matchEng.AssertExpectations(suite.T())
}

// TestHandleReceipt_Duplicate evaluates that handling a duplicate receipt is dropped
// without attempting to process it.
func (suite *FinderEngineTestSuite) TestHandleReceipt_Duplicate() {
	e := suite.TestNewFinderEngine()

	// mocks result has not yet processed
	suite.processedResults.On("Has", suite.receipt.ExecutionResult.ID()).Return(false).Once()

	// mocks adding receipt to the receipts mempool
	suite.receipts.On("Add", suite.receipt).Return(false).Once()

	// sends receipt to finder engine
	err := e.Process(suite.execIdentity.NodeID, suite.receipt)
	require.NoError(suite.T(), err)

	suite.matchEng.AssertNotCalled(suite.T(), "ProcessLocal", suite.receipt.ExecutionResult)

	suite.receipts.AssertExpectations(suite.T())
	suite.processedResults.AssertExpectations(suite.T())
}
