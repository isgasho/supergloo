package install_test

import (
	"fmt"

	"github.com/solo-io/supergloo/cli/pkg/helpers/clients"

	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/pkg/install/istio"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	skclients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/supergloo/cli/test/utils"
	"github.com/solo-io/supergloo/test/inputs"
)

var _ = Describe("Install", func() {

	BeforeEach(func() {
		clients.UseMemoryClients()
	})

	getInstall := func(name string) *v1.Install {
		in, err := clients.MustInstallClient().Read("supergloo-system", name, skclients.ReadOpts{})
		ExpectWithOffset(1, err).NotTo(HaveOccurred())
		return in
	}

	Context("non-interactive", func() {
		It("should create the expected install ", func() {
			installAndVerifyIstio := func(
				name,
				namespace,
				version string,
				mtls,
				autoInject,
				ingress,
				egress,
				prometheus,
				jaeger,
				grafana bool) {

				err := utils.Supergloo("install istio " +
					fmt.Sprintf("--name=%v ", name) +
					fmt.Sprintf("--installation-namespace istio ") +
					fmt.Sprintf("--version=%v ", version) +
					fmt.Sprintf("--mtls=%v ", mtls) +
					fmt.Sprintf("--auto-inject=%v ", autoInject) +
					fmt.Sprintf("--ingress=%v ", ingress) +
					fmt.Sprintf("--egress=%v ", egress) +
					fmt.Sprintf("--grafana=%v ", grafana) +
					fmt.Sprintf("--prometheus=%v ", prometheus) +
					fmt.Sprintf("--jaeger=%v", jaeger))
				if version == "badver" {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("is not a supported istio version"))
					return
				}

				Expect(err).NotTo(HaveOccurred())
				install := getInstall(name)
				istio := MustIstioInstallType(install)
				Expect(istio.Istio.Version).To(Equal(version))
				Expect(istio.Istio.EnableMtls).To(Equal(mtls))
				Expect(istio.Istio.EnableAutoInject).To(Equal(autoInject))
				Expect(istio.Istio.EnableIngress).To(Equal(ingress))
				Expect(istio.Istio.EnableEgress).To(Equal(egress))
				Expect(istio.Istio.InstallPrometheus).To(Equal(prometheus))
				Expect(istio.Istio.InstallJaeger).To(Equal(jaeger))
				Expect(istio.Istio.InstallGrafana).To(Equal(grafana))
			}

			installAndVerifyIstio("a1a", "ns", "1.0.3", true, true, true, true, true, true, true)
			installAndVerifyIstio("b1a", "ns", "1.0.5", false, false, false, false, false, false, false)
			installAndVerifyIstio("c1a", "ns", "badver", false, false, false, false, false, false, false)
		})
		It("should enable an existing + disabled install", func() {
			name := "input"
			namespace := "ns"
			inst := inputs.IstioInstall(name, namespace, "any", "1.0.5", true)
			ic := clients.MustInstallClient()
			_, err := ic.Write(inst, skclients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			err = utils.Supergloo("install istio " +
				fmt.Sprintf("--name=%v ", name) +
				fmt.Sprintf("--namespace=%v ", namespace))
			Expect(err).NotTo(HaveOccurred())

			updatedInstall, err := ic.Read(namespace, name, skclients.ReadOpts{})
			Expect(err).NotTo(HaveOccurred())
			Expect(updatedInstall.Disabled).To(BeFalse())

		})
		It("should error enable on existing enabled install", func() {
			name := "input"
			namespace := "ns"
			inst := inputs.IstioInstall(name, namespace, "any", "1.0.5", false)
			ic := clients.MustInstallClient()
			_, err := ic.Write(inst, skclients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			err = utils.Supergloo("install istio " +
				fmt.Sprintf("--name=%v ", name) +
				fmt.Sprintf("--namespace=%v ", namespace))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("already installed and enabled"))
		})
		It("should update existing enabled install if --update set to true", func() {
			name := "input"
			namespace := "ns"
			inst := inputs.IstioInstall(name, namespace, "istio-system", "1.0.5", false)
			Expect(inst.InstallType).To(BeAssignableToTypeOf(&v1.Install_Mesh{}))
			ic := clients.MustInstallClient()
			_, err := ic.Write(inst, skclients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			err = utils.Supergloo("install istio " +
				fmt.Sprintf("--name=%v ", name) +
				fmt.Sprintf("--namespace=%v ", namespace) +
				"--mtls=true " +
				"--update=true ")
			Expect(err).NotTo(HaveOccurred())

			updatedInstall, err := ic.Read(namespace, name, skclients.ReadOpts{})
			Expect(err).NotTo(HaveOccurred())
			Expect(*updatedInstall).To(Equal(v1.Install{
				Metadata: core.Metadata{
					Name:            name,
					Namespace:       namespace,
					ResourceVersion: updatedInstall.Metadata.ResourceVersion,
				},
				InstallationNamespace: "istio-system",
				Disabled:              false,
				InstallType: &v1.Install_Mesh{
					Mesh: &v1.MeshInstall{
						MeshInstallType: &v1.MeshInstall_Istio{
							Istio: &v1.IstioInstall{
								Version:          istio.IstioVersion106,
								EnableAutoInject: true,
								EnableMtls:       true,
							},
						},
					},
				},
			}))
		})
	})
})

func MustIstioInstallType(install *v1.Install) *v1.MeshInstall_Istio {
	Expect(install.InstallType).To(BeAssignableToTypeOf(&v1.Install_Mesh{}))
	mesh := install.InstallType.(*v1.Install_Mesh)
	Expect(mesh.Mesh.MeshInstallType).To(BeAssignableToTypeOf(&v1.MeshInstall_Istio{}))
	istioMesh := mesh.Mesh.MeshInstallType.(*v1.MeshInstall_Istio)
	return istioMesh
}
