package verification

import (
	"github.com/dapperlabs/flow-go/model/flow"
)

// PendingCollection represents a collection that its origin ID is pending to be verified
// It is utilized whenever the reference blockID for the pending collection is not available
type PendingCollection struct {
	Collection *flow.Collection
	OriginID   flow.Identifier
	Counter    uint64 // keeps track of number of retries
}

// ID returns the unique identifier for the pending receipt which is the
// id of its collection
func (p *PendingCollection) ID() flow.Identifier {
	return p.Collection.ID()
}

// Checksum returns the checksum of the pending collection.
func (p *PendingCollection) Checksum() flow.Identifier {
	return flow.MakeID(p)
}

// NewPendingCollection creates a new PendingCollection structure out of the collection and
// its originID. It also sets the counter value of receipt to one.
func NewPendingCollection(collection *flow.Collection, originID flow.Identifier) *PendingCollection {
	return &PendingCollection{
		Collection: collection,
		OriginID:   originID,
		Counter:    1,
	}
}
