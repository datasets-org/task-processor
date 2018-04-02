package structs

type Dataset struct {
	Id         string   `yaml:"id"`
	Data       []string `yaml:"data"`
	Name       string   `yaml:"name"`
	Internal   bool     `yaml:"internal"`
	From       string   `yaml:"from_ds"`
	Url        string   `yaml:"url"`
	Maintainer string   `yaml:"maintainer"`
	Tags       []string `yaml:"tags"`
	// todo uasage type
	Usages          []map[string]string `yaml:"usages"`
	Changelog       [][][]string        `yaml:"changelog"`
	Markdowns       []string            `yaml:"markdowns"`
	Characteristics map[string]string   `yaml:"characteristics"`
	Links           []string            `yaml:"links"`
	Path            []string            `yaml:"path"`
	Type            string              `yaml:"type"`
	Servers         []string            `yaml:"servers"`
}
