// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"sync"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
)

var _ = Describe("LinkerdDiscoveryEventLoop", func() {
	var (
		namespace string
		emitter   LinkerdDiscoveryEmitter
		err       error
	)

	BeforeEach(func() {

		meshClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		meshClient, err := NewMeshClient(meshClientFactory)
		Expect(err).NotTo(HaveOccurred())

		installClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		installClient, err := NewInstallClient(installClientFactory)
		Expect(err).NotTo(HaveOccurred())

		podClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		podClient, err := github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPodClient(podClientFactory)
		Expect(err).NotTo(HaveOccurred())

		upstreamClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		upstreamClient, err := gloo_solo_io.NewUpstreamClient(upstreamClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewLinkerdDiscoveryEmitter(meshClient, installClient, podClient, upstreamClient)
	})
	It("runs sync function on a new snapshot", func() {
		_, err = emitter.Mesh().Write(NewMesh(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.Install().Write(NewInstall(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.Pod().Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.Upstream().Write(gloo_solo_io.NewUpstream(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockLinkerdDiscoverySyncer{}
		el := NewLinkerdDiscoveryEventLoop(emitter, sync)
		_, err := el.Run([]string{namespace}, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(sync.Synced, 5*time.Second).Should(BeTrue())
	})
})

type mockLinkerdDiscoverySyncer struct {
	synced bool
	mutex  sync.Mutex
}

func (s *mockLinkerdDiscoverySyncer) Synced() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.synced
}

func (s *mockLinkerdDiscoverySyncer) Sync(ctx context.Context, snap *LinkerdDiscoverySnapshot) error {
	s.mutex.Lock()
	s.synced = true
	s.mutex.Unlock()
	return nil
}
