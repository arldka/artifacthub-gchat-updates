package artifacthub_gchat_updates

import (
	"log"
	"net/http"
)

type gc_header struct {
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	ImageUrl   string `json:"imageUrl"`
	ImageStyle string `json:"imageStyle"`
}

func findLogo(name string) string {
	res, err := http.Get("https://landscape.cncf.io/logos/" + name + ".svg")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusNotFound {
		return "https://landscape.cncf.io/logos/artifact-hub.svg"
	}
	return "https://landscape.cncf.io/logos/" + name + ".svg"
}

func messageHeader(p ah_payload) gc_header {
	var header gc_header
	header.Title = p.Data.Package.Name
	header.Subtitle = p.Data.Package.Version
	header.ImageUrl = findLogo(p.Data.Package.Name)
	header.ImageStyle = "IMAGE"
	return header
}
