package chat

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type gc_header struct {
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	ImageUrl   string `json:"imageUrl"`
	ImageStyle string `json:"imageStyle"`
}

type gc_key_value struct {
	Toplabel         string `json:"topLabel,omitempty"`
	Content          string `json:"content"`
	Contentmultiline bool   `json:"contentMultiline"`
}

type gc_openlink struct {
	Url string `json:"url"`
}

type gc_onclick struct {
	Openlink gc_openlink `json:"openLink"`
}

type gc_text_button struct {
	Text    string     `json:"text"`
	Onclick gc_onclick `json:"onClick"`
}

type gc_button struct {
	Text_button gc_text_button `json:"textButton"`
}

type gc_widget struct {
	Keyvalue *gc_key_value `json:"keyValue,omitempty"`
	Buttons  []gc_button   `json:"buttons,omitempty"`
}

type gc_section struct {
	Header  string      `json:"header,omitempty"`
	Widgets []gc_widget `json:"widgets"`
}

type gc_card struct {
	Header   gc_header    `json:"header"`
	Sections []gc_section `json:"sections"`
}

type gc_cards struct {
	Cards []gc_card `json:"cards"`
}

var cncf_projects []string = []string{
	"aerakimesh",
	"akri",
	"antrea",
	"argo",
	"artifacthub",
	"athenz",
	"backstage",
	"bfe",
	"brigade",
	"buildpacks",
	"cdk8s",
	"cert-manager",
	"chaosblade",
	"chaosmesh",
	"chubaofs",
	"cilium",
	"cloudcustodian",
	"cloudevents",
	"clusterpedia",
	"cncf-distribution",
	"cni",
	"confidential-containers",
	"containerd",
	"contour",
	"coredns",
	"cortex",
	"crio",
	"crossplane",
	"cubefs",
	"curiefense",
	"curve",
	"dapr",
	"devfile",
	"devstream",
	"dex",
	"dragonfly",
	"emissary-ingress",
	"envoy",
	"etcd",
	"fabedge",
	"falco",
	"file.txt",
	"fluentd",
	"fluid",
	"flux",
	"fonio",
	"grpc",
	"harbor",
	"helm",
	"inclavare",
	"in-toto",
	"jaeger",
	"k3s",
	"k8gb",
	"k8up",
	"karmada",
	"keda",
	"keptn",
	"keylime",
	"knative",
	"krator",
	"krustlet",
	"kubearmor",
	"kubedl",
	"kubeedge",
	"kube-ovn",
	"kuberhealthy",
	"kubernetes",
	"kube-rs",
	"kubevela",
	"kubevirt",
	"kudo",
	"kuma",
	"kyverno",
	"linkerd",
	"litmus",
	"longhorn",
	"meshery",
	"metal3",
	"metallb",
	"nats",
	"networkservicemesh",
	"nocalhost",
	"notary",
	"opa",
	"open-cluster-management",
	"opencost",
	"openebs",
	"openfeature",
	"openfunction",
	"opengitops",
	"openkruise",
	"openmetrics",
	"openservicemesh",
	"opentelemetry",
	"opentracing",
	"openyurt",
	"operatorframework",
	"oras",
	"parsec",
	"piraeus",
	"pixie",
	"porter",
	"pravega",
	"prometheus",
	"rkt",
	"rook",
	"schemahero",
	"serverlessworkflow",
	"servicemeshinterface",
	"servicemeshperformance",
	"skooner",
	"spiffe",
	"spire",
	"strimzi",
	"submariner",
	"superedge",
	"telepresence",
	"teller",
	"thanos",
	"tikv",
	"tinkerbell",
	"tremor",
	"trickster",
	"tuf",
	"vineyard",
	"virtualkubelet",
	"vitess",
	"volcano",
	"wasmcloud",
	"wasm-edge-runtime",
}

func is_cncf(str string) string {
	for _, project := range cncf_projects {
		if strings.Contains(str, project) {
			return project
		}
	}
	return ""
}

func findLogo(name string) string {
	cncf_project := is_cncf(name)
	if cncf_project != "" {
		res, err := http.Get("https://raw.githubusercontent.com/cncf/artwork/master/projects/" + cncf_project + "/icon/color/" + cncf_project + "-icon-color.png")
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode == http.StatusNotFound {
			return "https://raw.githubusercontent.com/cncf/artwork/master/projects/artifacthub/icon/color/artifacthub-icon-color.png"
		} else {
			return "https://raw.githubusercontent.com/cncf/artwork/master/projects/" + cncf_project + "/icon/color/" + cncf_project + "-icon-color.png"
		}
	}
	return "https://raw.githubusercontent.com/cncf/artwork/master/projects/artifacthub/icon/color/artifacthub-icon-color.png"
}

func messageHeader(p *ah_payload) gc_header {
	var header gc_header
	header.Title = p.Pkg.Name
	header.Subtitle = p.Pkg.Version
	header.ImageUrl = findLogo(p.Pkg.Name)
	header.ImageStyle = "IMAGE"
	return header
}

func changesSection(p *ah_payload) gc_section {
	var section gc_section
	if len(p.Pkg.Changes) == 0 {
		section.Header = "No changes"
	} else {
		section.Header = "Changes"
		section.Widgets = []gc_widget{}
		for _, change := range p.Pkg.Changes {
			section.Widgets = append(section.Widgets, gc_widget{
				Keyvalue: &gc_key_value{
					Toplabel:         "",
					Content:          change,
					Contentmultiline: true,
				},
			})
		}
	}
	return section
}

func buttonSection(p *ah_payload) gc_section {
	var section gc_section
	var button gc_button
	button.Text_button.Text = "View on Artifact Hub"
	button.Text_button.Onclick.Openlink.Url = p.Pkg.Url
	var buttons []gc_button = []gc_button{button}
	var widget gc_widget = gc_widget{Buttons: buttons}
	section.Widgets = []gc_widget{widget}
	return section
}

func gcMessageGenerator(p *ah_payload) gc_cards {
	var card gc_card
	card.Header = messageHeader(p)
	fmt.Printf("Header: %+v\n", card.Header)
	if changesSection(p).Widgets == nil {
		card.Sections = []gc_section{buttonSection(p)}
	} else {
		card.Sections = []gc_section{changesSection(p), buttonSection(p)}
	}
	return gc_cards{Cards: []gc_card{card}}
}
