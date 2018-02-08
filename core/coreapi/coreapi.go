package coreapi

import (
	core "github.com/ipfs/go-ipfs/core"
	coreiface "github.com/ipfs/go-ipfs/core/coreapi/interface"
)

type CoreAPI struct {
	node *core.IpfsNode
}

// NewCoreAPI creates new instance of IPFS CoreAPI backed by go-ipfs Node.
func NewCoreAPI(n *core.IpfsNode) coreiface.CoreAPI {
	api := &CoreAPI{n}
	return api
}

// Unixfs returns the UnixfsAPI interface implementation backed by the go-ipfs node
func (api *CoreAPI) Unixfs() coreiface.UnixfsAPI {
	return (*UnixfsAPI)(api)
}

// Block returns the BlockAPI interface implementation backed by the go-ipfs node
func (api *CoreAPI) Block() coreiface.BlockAPI {
	return &BlockAPI{api, nil}
}

// Dag returns the DagAPI interface implementation backed by the go-ipfs node
func (api *CoreAPI) Dag() coreiface.DagAPI {
	return &DagAPI{api, nil}
}

// Name returns the NameAPI interface implementation backed by the go-ipfs node
func (api *CoreAPI) Name() coreiface.NameAPI {
	return &NameAPI{api, nil}
}

// Key returns the KeyAPI interface implementation backed by the go-ipfs node
func (api *CoreAPI) Key() coreiface.KeyAPI {
	return &KeyAPI{api, nil}
}

//Object returns the ObjectAPI interface implementation backed by the go-ipfs node
func (api *CoreAPI) Object() coreiface.ObjectAPI {
	return &ObjectAPI{api, nil}
}

// Pin returns the PinAPI interface implementation backed by the go-ipfs node
func (api *CoreAPI) Pin() coreiface.PinAPI {
	return &PinAPI{api, nil}
}
