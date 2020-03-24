package fdfs_client

type Configuration struct {
	Trackers    []string `json:"trackers" xml:"trackers" yaml:"trackers"`
	MaxOpenCons int      `json:"max_open_cons" xml:"max_open_cons" yaml:"trackers"`
}
