package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onflow/flow/protobuf/go/flow/access"

	"github.com/dapperlabs/flow-go/engine/common/convert"
	"github.com/dapperlabs/flow-go/model/flow"
)

func (h *Handler) GetLatestBlockHeader(ctx context.Context, req *access.GetLatestBlockHeaderRequest) (*access.BlockHeaderResponse, error) {

	var header *flow.Header
	var err error
	if req.IsSealed {
		// get the latest seal header from storage
		header, err = h.getLatestSealedHeader()
	} else {
		// get the finalized header from state
		header, err = h.state.Final().Head()
	}

	if err != nil {
		err = convertStorageError(err)
		return nil, err
	}

	return createBlockHeaderResponse(header)
}

func (h *Handler) GetBlockHeaderByID(_ context.Context, req *access.GetBlockHeaderByIDRequest) (*access.BlockHeaderResponse, error) {

	id := flow.HashToID(req.Id)
	header, err := h.headers.ByBlockID(id)
	if err != nil {
		err = convertStorageError(err)
		return nil, err
	}

	return createBlockHeaderResponse(header)
}

func (h *Handler) GetBlockHeaderByHeight(_ context.Context, req *access.GetBlockHeaderByHeightRequest) (*access.BlockHeaderResponse, error) {

	header, err := h.headers.ByNumber(req.Height)
	if err != nil {
		err = convertStorageError(err)
		return nil, err
	}

	return createBlockHeaderResponse(header)
}

func createBlockHeaderResponse(header *flow.Header) (*access.BlockHeaderResponse, error) {
	msg, err := convert.BlockHeaderToMessage(header)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not convert block header: %s", err.Error())
	}

	resp := &access.BlockHeaderResponse{
		Block: &msg,
	}
	return resp, nil
}
