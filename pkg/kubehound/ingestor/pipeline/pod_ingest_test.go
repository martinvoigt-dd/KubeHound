package pipeline

import (
	"context"
	"testing"

	"github.com/DataDog/KubeHound/pkg/collector"
	mockcollect "github.com/DataDog/KubeHound/pkg/collector/mockcollector"
	"github.com/DataDog/KubeHound/pkg/globals/types"
	"github.com/DataDog/KubeHound/pkg/kubehound/models/store"
	cache "github.com/DataDog/KubeHound/pkg/kubehound/storage/cache/mocks"
	graphdb "github.com/DataDog/KubeHound/pkg/kubehound/storage/graphdb/mocks"
	storedb "github.com/DataDog/KubeHound/pkg/kubehound/storage/storedb/mocks"
	"github.com/DataDog/KubeHound/pkg/kubehound/store/collections"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPodIngest_Pipeline(t *testing.T) {
	pi := &PodIngest{}

	ctx := context.Background()
	fakePod, err := loadTestObject[types.PodType]("testdata/pod.json")
	assert.NoError(t, err)

	client := mockcollect.NewCollectorClient(t)
	client.EXPECT().StreamPods(ctx, pi).
		RunAndReturn(func(ctx context.Context, i collector.PodIngestor) error {
			// Fake the stream of a single pod from the collector client
			err := i.IngestPod(ctx, fakePod)
			if err != nil {
				return err
			}

			return i.Complete(ctx)
		})

	// Cache setup
	c := cache.NewCacheProvider(t)
	cw := cache.NewAsyncWriter(t)
	cw.EXPECT().Queue(ctx, mock.AnythingOfType("*cache.containerCacheKey"), mock.AnythingOfType("string")).Return(nil).Once()
	cw.EXPECT().Flush(ctx).Return(nil)
	cw.EXPECT().Close(ctx).Return(nil)
	c.EXPECT().BulkWriter(ctx).Return(cw, nil)
	c.EXPECT().Get(ctx, mock.AnythingOfType("*cache.nodeCacheKey")).Return(store.ObjectID().Hex(), nil)
	c.EXPECT().Get(ctx, mock.AnythingOfType("*cache.containerCacheKey")).Return(store.ObjectID().Hex(), nil)

	// Store setup - pods
	sdb := storedb.NewProvider(t)
	psw := storedb.NewAsyncWriter(t)
	pods := collections.Pod{}
	psw.EXPECT().Queue(ctx, mock.AnythingOfType("*store.Pod")).Return(nil).Once()
	psw.EXPECT().Flush(ctx).Return(nil)
	psw.EXPECT().Close(ctx).Return(nil)

	// Store setup - containers
	csw := storedb.NewAsyncWriter(t)
	containers := collections.Container{}
	csw.EXPECT().Queue(ctx, mock.AnythingOfType("*store.Container")).Return(nil).Once()
	csw.EXPECT().Flush(ctx).Return(nil)
	csw.EXPECT().Close(ctx).Return(nil)

	// Store setup - volumes
	vsw := storedb.NewAsyncWriter(t)
	volumes := collections.Volume{}
	vsw.EXPECT().Queue(ctx, mock.AnythingOfType("*store.Volume")).Return(nil).Once()
	vsw.EXPECT().Flush(ctx).Return(nil)
	vsw.EXPECT().Close(ctx).Return(nil)

	sdb.EXPECT().BulkWriter(ctx, pods).Return(psw, nil)
	sdb.EXPECT().BulkWriter(ctx, containers).Return(csw, nil)
	sdb.EXPECT().BulkWriter(ctx, volumes).Return(vsw, nil)

	// Graph setup - pods
	gdb := graphdb.NewProvider(t)
	pgw := graphdb.NewAsyncVertexWriter(t)
	pgw.EXPECT().Queue(ctx, mock.AnythingOfType("*graph.Pod")).Return(nil).Once()
	pgw.EXPECT().Flush(ctx).Return(nil)
	pgw.EXPECT().Close(ctx).Return(nil)

	// Graph setup - containers
	cgw := graphdb.NewAsyncVertexWriter(t)
	cgw.EXPECT().Queue(ctx, mock.AnythingOfType("*graph.Container")).Return(nil).Once()
	cgw.EXPECT().Flush(ctx).Return(nil)
	cgw.EXPECT().Close(ctx).Return(nil)

	// Graph setup - volumes
	vgw := graphdb.NewAsyncVertexWriter(t)
	vgw.EXPECT().Queue(ctx, mock.AnythingOfType("*graph.Volume")).Return(nil).Once()
	vgw.EXPECT().Flush(ctx).Return(nil)
	vgw.EXPECT().Close(ctx).Return(nil)

	gdb.EXPECT().VertexWriter(ctx, mock.AnythingOfType("vertex.Pod")).Return(pgw, nil)
	gdb.EXPECT().VertexWriter(ctx, mock.AnythingOfType("vertex.Container")).Return(cgw, nil)
	gdb.EXPECT().VertexWriter(ctx, mock.AnythingOfType("vertex.Volume")).Return(vgw, nil)

	deps := &Dependencies{
		Collector: client,
		Cache:     c,
		GraphDB:   gdb,
		StoreDB:   sdb,
	}

	// Initialize
	err = pi.Initialize(ctx, deps)
	assert.NoError(t, err)

	// Run
	err = pi.Run(ctx)
	assert.NoError(t, err)

	// Close
	err = pi.Close(ctx)
	assert.NoError(t, err)
}