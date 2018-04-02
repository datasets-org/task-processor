package structs

type Dataset struct {
	Id         string   `yaml:"id" json:"id"`
	Data       []string `yaml:"data" json:"data"`
	Name       string   `yaml:"name" json:"name"`
	Internal   bool     `yaml:"internal" json:"internal"`
	From       string   `yaml:"from_ds" json:"from_ds"`
	Url        string   `yaml:"url" json:"url"`
	Maintainer string   `yaml:"maintainer" json:"maintainer"`
	Tags       []string `yaml:"tags" json:"tags"`
	// todo usage type struct
	Usages          []map[string]string `yaml:"usages" json:"tags"`
	// todo changelog type struct
	Changelog       [][][]string        `yaml:"changelog" json:"changelog"`
	Markdowns       []string            `yaml:"markdowns" json:"markdowns"`
	// todo characteristics type struct
	Characteristics map[string]string   `yaml:"characteristics" json:"characteristics"`
	Links           []string            `yaml:"links" json:"links"`
	Path            []string            `yaml:"path" json:"path"`
	Type            string              `yaml:"type" json:"type"`
	Servers         []string            `yaml:"servers" json:"servers"`
}
