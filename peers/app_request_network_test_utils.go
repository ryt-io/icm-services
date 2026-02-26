package peers

import (
	"sync"

	"github.com/ryt-io/ryt-v2/ids"
	"github.com/ryt-io/ryt-v2/network"
	snowVdrs "github.com/ryt-io/ryt-v2/snow/validators"
	"github.com/ryt-io/ryt-v2/utils/linked"
	"github.com/ryt-io/ryt-v2/utils/logging"
	"github.com/ryt-io/ryt-v2/utils/set"
	"github.com/ryt-io/icm-services/cache"
	"github.com/ryt-io/icm-services/peers/clients"
)

// NewAppRequestNetworkForTesting creates an AppRequestNetwork instance for testing
// with injected mock dependencies
func NewAppRequestNetworkForTesting(
	mockNetwork network.Network,
	handler *RelayerExternalHandler,
	logger logging.Logger,
	metrics *AppRequestNetworkMetrics,
	mockValidatorClient clients.CanonicalValidatorState,
	manager snowVdrs.Manager,
) *AppRequestNetwork {
	// Initialize cache for testing
	epochedValidatorSetCache := cache.NewFIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet](10)

	validatorManager := &ValidatorManager{
		logger:                   logger,
		validatorClient:          mockValidatorClient,
		metrics:                  metrics,
		maxPChainLookback:        -1, // No lookback limit for tests
		epochedValidatorSetCache: epochedValidatorSetCache,
		manager:                  manager,
		validatorSetLock:         new(sync.Mutex),
	}

	return &AppRequestNetwork{
		network:            mockNetwork,
		handler:            handler,
		logger:             logger,
		metrics:            metrics,
		trackedSubnets:     set.NewSet[ids.ID](0),
		lruSubnets:         linked.NewHashmap[ids.ID, interface{}](),
		trackedSubnetsLock: new(sync.RWMutex),
		validatorManager:   validatorManager,
	}
}
