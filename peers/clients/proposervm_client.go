// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package clients

import (
	"context"

	"github.com/ryt-io/ryt-v2/utils/rpc"
	"github.com/ryt-io/ryt-v2/vms/proposervm"
	"github.com/ryt-io/ryt-v2/vms/proposervm/block"
	"github.com/ryt-io/icm-services/config"
)

// ProposerVMAPI is a wrapper around a [proposervm.JSONRPCClient],
// and provides additional options for the API passed in the config.
type ProposerVMAPI struct {
	client  *proposervm.JSONRPCClient
	options []rpc.Option
}

func NewProposerVMAPI(uri string, chain string, cfg *config.APIConfig) *ProposerVMAPI {
	return &ProposerVMAPI{
		client:  proposervm.NewJSONRPCClient(uri, chain),
		options: cfg.Options(),
	}
}

func (p *ProposerVMAPI) GetCurrentEpoch(ctx context.Context) (block.Epoch, error) {
	return p.client.GetCurrentEpoch(ctx, p.options...)
}

func (p *ProposerVMAPI) GetProposedHeight(ctx context.Context) (uint64, error) {
	return p.client.GetProposedHeight(ctx, p.options...)
}
