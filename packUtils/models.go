package packUtils

type Package struct {
	Head HeadPackage `json:"head"`
	Body BodyPackage `json:"body"`
}

type HeadPackage struct {
	Rand    string `json:"rand"`
	Title   string `json:"title"`
	Sender  string `json:"sender"`
	Session string `json:"session"`
	Meta    string `json:"meta"`
}

type BodyPackage struct {
	Date string `json:"date"`
	Data string `json:"data"`
	Hash string `json:"hash"`
	Sign string `json:"sign"`
}
