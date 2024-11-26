package pollingprovider

import (
	"context"
	"log"
	"time"

	"github.com/ArnaudovSt/tx-parser/client"
	"github.com/ArnaudovSt/tx-parser/provider"
	"github.com/ArnaudovSt/tx-parser/storage"
	"github.com/ArnaudovSt/tx-parser/types"
)

var _ provider.IProvider = (*pollingProvider)(nil)

type pollingProvider struct {
	client       client.IClient
	storage      storage.IStorage
	avgBlockTime time.Duration
}

func NewPollingProvider(client client.IClient, storage storage.IStorage, avgBlockTime time.Duration) *pollingProvider {
	return &pollingProvider{
		client:       client,
		storage:      storage,
		avgBlockTime: avgBlockTime,
	}
}

func (p *pollingProvider) Start(ctx context.Context) error {
	ticker := time.NewTicker(p.avgBlockTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := p.poll()
			if err != nil {
				log.Printf("[WARN] polling failed: %v\n", err)
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (p *pollingProvider) poll() error {
	currentLatestBlock, err := p.storage.GetLatestBlock()
	if err != nil {
		return err
	}

	newLatestBlock, err := p.client.GetLatestBlock(context.Background())
	if err != nil {
		return err
	}

	switch {
	case currentLatestBlock == nil:
		log.Printf("[DEBUG] Initialising storage with block %s %d\n", newLatestBlock.Hash, newLatestBlock.Number.Uint64())
		return p.storage.AtomicWrite(func(w storage.IAtomicWriter) error {
			return w.AppendBlock(newLatestBlock)
		})
	case currentLatestBlock.Hash == newLatestBlock.Hash:
		// no new block
		log.Printf("[DEBUG] No new block identified. Currently at block %s %d\n", currentLatestBlock.Hash, currentLatestBlock.Number.Uint64())
		return nil
	case currentLatestBlock.Hash == newLatestBlock.ParentHash:
		// new consecutive block
		log.Printf("[DEBUG] Appending new consecutive block %s %d\n", newLatestBlock.Hash, newLatestBlock.Number.Uint64())
		return p.storage.AtomicWrite(func(w storage.IAtomicWriter) error {
			return w.AppendBlock(newLatestBlock)
		})
	default:
		// reorg
		blocksToAdd := make([]*types.Block, 0)

		// reach the same block height as the current latest block
		blockPointer := newLatestBlock
		var err error
		for blockPointer.Number.Uint64() > currentLatestBlock.Number.Uint64() {
			blocksToAdd = append(blocksToAdd, blockPointer)

			blockPointer, err = p.client.GetBlockByHash(context.Background(), blockPointer.ParentHash)
			if err != nil {
				return err
			}
		}

		return p.storage.AtomicWrite(func(w storage.IAtomicWriter) error {
			return p.handleReorg(w, currentLatestBlock, blockPointer, blocksToAdd)
		})
	}
}

func (p *pollingProvider) handleReorg(w storage.IAtomicWriter, currentLatestBlock *types.Block, blockPointer *types.Block, blocksToAdd []*types.Block) error {
	// compare the blocks by hashes and iterate until the common ancestor is found
	var err error
	for blockPointer.Hash != currentLatestBlock.Hash {
		blocksToAdd = append(blocksToAdd, blockPointer)

		blockPointer, err = p.client.GetBlockByHash(context.Background(), blockPointer.ParentHash)
		if err != nil {
			return err
		}

		log.Printf("[DEBUG] Removing reorged block %s %d\n", currentLatestBlock.Hash, currentLatestBlock.Number.Uint64())
		currentLatestBlock, err = w.PopLatestBlock()
		if err != nil {
			return err
		}
	}

	// add the blocks in reverse order
	for i := len(blocksToAdd) - 1; i >= 0; i-- {
		log.Printf("[DEBUG] Appending block to state %s %d\n", blocksToAdd[i].Hash, blocksToAdd[i].Number.Uint64())
		err = w.AppendBlock(blocksToAdd[i])
		if err != nil {
			return err
		}
	}

	return nil
}
