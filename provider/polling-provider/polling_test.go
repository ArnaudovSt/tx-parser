package pollingprovider

import (
	"context"
	"math/big"
	"testing"
	"time"

	clientMocks "github.com/ArnaudovSt/tx-parser/client/mocks"
	storageMocks "github.com/ArnaudovSt/tx-parser/storage/mocks"
	"github.com/ArnaudovSt/tx-parser/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewPollingProvider(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := clientMocks.NewMockIClient(ctrl)
	mockStorage := storageMocks.NewMockIStorage(ctrl)
	avgBlockTime := time.Second * 10

	provider := NewPollingProvider(mockClient, mockStorage, avgBlockTime)
	assert.NotNil(t, provider)
	assert.Equal(t, mockClient, provider.client)
	assert.Equal(t, mockStorage, provider.storage)
	assert.Equal(t, avgBlockTime, provider.avgBlockTime)
}

func TestPollingProvider_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := clientMocks.NewMockIClient(ctrl)
	mockStorage := storageMocks.NewMockIStorage(ctrl)
	avgBlockTime := time.Millisecond * 100

	provider := NewPollingProvider(mockClient, mockStorage, avgBlockTime)
	ctx, cancel := context.WithCancel(context.Background())

	mockStorage.EXPECT().GetLatestBlock().Return(nil, nil).AnyTimes()
	mockClient.EXPECT().GetLatestBlock(gomock.Any()).Return(&types.Block{Hash: "hash1", Number: big.NewInt(1)}, nil).AnyTimes()
	mockStorage.EXPECT().AtomicWrite(gomock.Any()).Return(nil).AnyTimes()

	go func() {
		time.Sleep(time.Millisecond * 300)
		cancel()
	}()

	err := provider.Start(ctx)
	assert.NoError(t, err)
}

func TestPollingProvider_poll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := clientMocks.NewMockIClient(ctrl)
	mockStorage := storageMocks.NewMockIStorage(ctrl)
	provider := NewPollingProvider(mockClient, mockStorage, time.Second)

	t.Run("no latest block in storage", func(t *testing.T) {
		mockStorage.EXPECT().GetLatestBlock().Return(nil, nil)
		mockClient.EXPECT().GetLatestBlock(gomock.Any()).Return(&types.Block{Hash: "hash1", Number: big.NewInt(1)}, nil)
		mockStorage.EXPECT().AtomicWrite(gomock.Any()).Return(nil)

		err := provider.poll()
		assert.NoError(t, err)
	})

	t.Run("no new block", func(t *testing.T) {
		mockStorage.EXPECT().GetLatestBlock().Return(&types.Block{Hash: "hash1", Number: big.NewInt(1)}, nil)
		mockClient.EXPECT().GetLatestBlock(gomock.Any()).Return(&types.Block{Hash: "hash1", Number: big.NewInt(1)}, nil)

		err := provider.poll()
		assert.NoError(t, err)
	})

	t.Run("new consecutive block", func(t *testing.T) {
		mockStorage.EXPECT().GetLatestBlock().Return(&types.Block{Hash: "hash1", Number: big.NewInt(1)}, nil)
		mockClient.EXPECT().GetLatestBlock(gomock.Any()).Return(&types.Block{Hash: "hash2", ParentHash: "hash1", Number: big.NewInt(2)}, nil)
		mockStorage.EXPECT().AtomicWrite(gomock.Any()).Return(nil)

		err := provider.poll()
		assert.NoError(t, err)
	})

	t.Run("reorg", func(t *testing.T) {
		mockStorage.EXPECT().GetLatestBlock().Return(&types.Block{Number: big.NewInt(1), Hash: "hash1", ParentHash: ""}, nil)
		mockClient.EXPECT().GetLatestBlock(gomock.Any()).Return(&types.Block{Number: big.NewInt(3), Hash: "hash3", ParentHash: "hash2"}, nil)
		gomock.InOrder(
			mockClient.EXPECT().GetBlockByHash(gomock.Any(), "hash2").Return(&types.Block{Number: big.NewInt(2), Hash: "hash2", ParentHash: "hash1"}, nil),
			mockClient.EXPECT().GetBlockByHash(gomock.Any(), "hash1").Return(&types.Block{Number: big.NewInt(1), Hash: "hash1", ParentHash: ""}, nil),
		)
		mockStorage.EXPECT().AtomicWrite(gomock.Any()).Return(nil)

		err := provider.poll()
		assert.NoError(t, err)
	})
}
