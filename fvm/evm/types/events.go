package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/runtime/common"
	gethCommon "github.com/onflow/go-ethereum/common"
	"github.com/onflow/go-ethereum/rlp"

	"github.com/onflow/flow-go/model/flow"
)

const (
	EventTypeBlockExecuted       flow.EventType = "BlockExecuted"
	EventTypeTransactionExecuted flow.EventType = "TransactionExecuted"
	locationDivider                             = "."
)

type EventPayload interface {
	// CadenceEvent creates a Cadence event type
	CadenceEvent() (cadence.Event, error)
}

type Event struct {
	Etype   flow.EventType
	Payload EventPayload
}

var _ common.Location = EVMLocation{}

type EVMLocation struct{}

func (l EVMLocation) TypeID(memoryGauge common.MemoryGauge, qualifiedIdentifier string) common.TypeID {
	id := fmt.Sprintf("%s%s%s", flow.EVMLocationPrefix, locationDivider, qualifiedIdentifier)
	common.UseMemory(memoryGauge, common.NewRawStringMemoryUsage(len(id)))

	return common.TypeID(id)
}

func (l EVMLocation) QualifiedIdentifier(typeID common.TypeID) string {
	pieces := strings.SplitN(string(typeID), locationDivider, 2)

	if len(pieces) < 2 {
		return ""
	}

	return pieces[1]
}

func (l EVMLocation) String() string {
	return flow.EVMLocationPrefix
}

func (l EVMLocation) Description() string {
	return flow.EVMLocationPrefix
}

func (l EVMLocation) ID() string {
	return flow.EVMLocationPrefix
}

func (l EVMLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string
	}{
		Type: "EVMLocation",
	})
}

func init() {
	common.RegisterTypeIDDecoder(
		flow.EVMLocationPrefix,
		func(_ common.MemoryGauge, typeID string) (common.Location, string, error) {
			if typeID == "" {
				return nil, "", fmt.Errorf("invalid EVM type location ID: missing type prefix")
			}

			parts := strings.SplitN(typeID, ".", 2)
			prefix := parts[0]
			if prefix != flow.EVMLocationPrefix {
				return EVMLocation{}, "", fmt.Errorf("invalid EVM type location ID: invalid prefix")
			}

			var qualifiedIdentifier string
			pieceCount := len(parts)
			if pieceCount > 1 {
				qualifiedIdentifier = parts[1]
			}

			return EVMLocation{}, qualifiedIdentifier, nil
		},
	)
}

// todo we might have to break this event into two (tx included /tx executed) if size becomes an issue

type TransactionExecutedPayload struct {
	BlockHeight uint64
	TxEncoded   []byte
	BlockHash   gethCommon.Hash
	TxHash      gethCommon.Hash
	Result      *Result
}

func (p *TransactionExecutedPayload) CadenceEvent() (cadence.Event, error) {
	var encodedLogs []byte
	var err error
	if len(p.Result.Logs) > 0 {
		encodedLogs, err = rlp.EncodeToBytes(p.Result.Logs)
		if err != nil {
			return cadence.Event{}, err
		}
	}

	return cadence.Event{
		EventType: cadence.NewEventType(
			EVMLocation{},
			string(EventTypeTransactionExecuted),
			[]cadence.Field{
				cadence.NewField("blockHeight", cadence.UInt64Type{}),
				cadence.NewField("blockHash", cadence.StringType{}),
				cadence.NewField("transactionHash", cadence.StringType{}),
				cadence.NewField("transaction", cadence.StringType{}),
				cadence.NewField("failed", cadence.BoolType{}),
				cadence.NewField("vmError", cadence.StringType{}),
				cadence.NewField("transactionType", cadence.UInt8Type{}),
				cadence.NewField("gasConsumed", cadence.UInt64Type{}),
				cadence.NewField("deployedContractAddress", cadence.StringType{}),
				cadence.NewField("returnedValue", cadence.StringType{}),
				cadence.NewField("logs", cadence.StringType{}),
			},
			nil,
		),
		Fields: []cadence.Value{
			cadence.NewUInt64(p.BlockHeight),
			cadence.String(p.BlockHash.String()),
			cadence.String(p.TxHash.String()),
			cadence.String(hex.EncodeToString(p.TxEncoded)),
			cadence.Bool(p.Result.Failed()),
			cadence.String(p.Result.VMErrorString()),
			cadence.NewUInt8(p.Result.TxType),
			cadence.NewUInt64(p.Result.GasConsumed),
			cadence.String(p.Result.DeployedContractAddress.String()),
			cadence.String(hex.EncodeToString(p.Result.ReturnedValue)),
			cadence.String(hex.EncodeToString(encodedLogs)),
		},
	}, nil
}

func NewTransactionExecutedEvent(
	height uint64,
	txEncoded []byte,
	blockHash gethCommon.Hash,
	txHash gethCommon.Hash,
	result *Result,
) *Event {
	return &Event{
		Etype: EventTypeTransactionExecuted,
		Payload: &TransactionExecutedPayload{
			BlockHeight: height,
			BlockHash:   blockHash,
			TxEncoded:   txEncoded,
			TxHash:      txHash,
			Result:      result,
		},
	}
}

var blockExecutedEventCadenceType = &cadence.EventType{
	Location:            EVMLocation{},
	QualifiedIdentifier: string(EventTypeBlockExecuted),
	Fields: []cadence.Field{
		cadence.NewField("height", cadence.UInt64Type{}),
		cadence.NewField("hash", cadence.StringType{}),
		cadence.NewField("timestamp", cadence.UInt64Type{}),
		cadence.NewField("totalSupply", cadence.IntType{}),
		cadence.NewField("parentHash", cadence.StringType{}),
		cadence.NewField("receiptRoot", cadence.StringType{}),
		cadence.NewField(
			"transactionHashes",
			cadence.NewVariableSizedArrayType(cadence.StringType{}),
		),
	},
}

type BlockExecutedEventPayload struct {
	Block *Block
}

func (p *BlockExecutedEventPayload) CadenceEvent() (cadence.Event, error) {
	hashes := make([]cadence.Value, len(p.Block.TransactionHashes))
	for i, hash := range p.Block.TransactionHashes {
		hashes[i] = cadence.String(hash.String())
	}

	blockHash, err := p.Block.Hash()
	if err != nil {
		return cadence.Event{}, err
	}

	fields := []cadence.Value{
		cadence.NewUInt64(p.Block.Height),
		cadence.String(blockHash.String()),
		cadence.NewUInt64(p.Block.Timestamp),
		cadence.NewIntFromBig(p.Block.TotalSupply),
		cadence.String(p.Block.ParentBlockHash.String()),
		cadence.String(p.Block.ReceiptRoot.String()),
		cadence.NewArray(hashes).WithType(cadence.NewVariableSizedArrayType(cadence.StringType{})),
	}

	return cadence.
		NewEvent(fields).
		WithType(blockExecutedEventCadenceType), nil
}

func NewBlockExecutedEvent(block *Block) *Event {
	return &Event{
		Etype: EventTypeBlockExecuted,
		Payload: &BlockExecutedEventPayload{
			Block: block,
		},
	}
}
