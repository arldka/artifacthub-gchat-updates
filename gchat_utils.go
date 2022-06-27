package artifacthub_gchat_updates

import (
	"log"
	"fmt"
	"net/http"
)


type gc_header struct {
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	ImageUrl   string `json:"imageUrl"`
	ImageStyle string `json:"imageStyle"`
}

type gc_key_value struct {
	Toplabel string `json:"topLabel, omitempty"`
	Content  string `json:"content"`
	Contentmultiline bool `json:"contentMultiline"`
}

type gc_openlink struct {
	Url string `json:"url"`
} 

type gc_onclick struct {
	Openlink gc_openlink `json:"openLink"`
}

type gc_text_button struct {
	Text string `json:"text"`
	Onclick gc_onclick `json:"onClick"`
}

type gc_button struct {
	Text_button gc_text_button `json:"textButton"`
}

type gc_widget struct {
	Keyvalue *gc_key_value `json:"keyValue, omitempty"`
	Buttons []gc_button `json:"buttons, omitempty"`
}

type gc_section struct {
	Header string `json:"header, omitempty"`
	Widgets []gc_widget `json:"widgets"`
}

type gc_card struct {
	Header gc_header `json:"header"`
	Sections []gc_section `json:"sections"`
}

type gc_cards struct {
	Cards []gc_card `json:"cards"`
}

func findLogo(name string) string {
	res, err := http.Get("https://raw.githubusercontent.com/cncf/artwork/master/projects/" + name + "/icon/color/" + name + "-icon-color.png")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusNotFound {
		return "https://raw.githubusercontent.com/cncf/artwork/master/projects/artifacthub/icon/color/artifacthub-icon-color.png"
	}
	return "https://raw.githubusercontent.com/cncf/artwork/master/projects/" + name + "/icon/color/" + name + "-icon-color.png"
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
					Toplabel: "",
					Content: change,
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
	var version_widget gc_widget = gc_widget{Keyvalue: &gc_key_value{Toplabel: "Version", Content: p.Pkg.Version, Contentmultiline: false}}
	card.Sections = []gc_section{gc_section{Widgets: []gc_widget{version_widget}}, changesSection(p), buttonSection(p)}
	return gc_cards{Cards: []gc_card{card}}
}
