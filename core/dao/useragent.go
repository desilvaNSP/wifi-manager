package dao

type UserAgent struct {
	Browser Browser `json:"browser"`
	Cpu CPU         `json:"cpu"`
	Device Device   `json:"device"`
	Engine Engine   `json:"engine"`
	Os OS           `json:"os"`
	Ua string       `json:"ua"`
}

type Browser struct {
	Major string   `json:"major"`
	Name string    `json:"name"`
	Version string `json:"version"`
}

type CPU struct {
	Architecture string `json:"architecture"`
}

type Device struct {
	Model string `json:"model"`
	Type string `json:"type"`
	Vendor string `json:"vendor"`
}

type Engine struct {
	Name string     `json:"name"`
	Version string  `json:"version"`
}

type OS struct {
	Name string     `json:"name"`
	Version string  `json:"version"`
}