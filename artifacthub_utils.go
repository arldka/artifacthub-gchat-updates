package artifacthub_gchat_updates

type ah_payload struct {
	Specversion     string  `json:"specversion"`
	Id              string  `json:"id"`
	Source          string  `json:"source"`
	Type            string  `json:"type"`
	Datacontenttype string  `json:"datacontenttype"`
	Data            ah_data `json:"data"`
}

type ah_data struct {
	Package ah_pkg `json:"package"`
}

type ah_pkg struct {
	Name       string        `json:"name"`
	Version    string        `json:"version"`
	Url        string        `json:"url"`
	Changes    []string      `json:"changes"`
	Repository ah_repository `json:"repository"`
}

type ah_repository struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}
