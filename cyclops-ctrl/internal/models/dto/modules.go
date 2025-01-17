package dto

type Module struct {
	Name      string      `json:"name"`
	Namespace string      `json:"namespace"`
	Template  Template    `json:"template"`
	Version   string      `json:"version"`
	Values    interface{} `json:"values"`
	Status    string      `json:"status"`
}

type Template struct {
	Name    string         `json:"name"`
	Version string         `json:"version"`
	GitRef  TemplateGitRef `json:"git"`
}

type TemplateGitRef struct {
	Repo string `json:"repo"`
	Path string `json:"path"`
}

type TemplatesResponse struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

type DeleteResource struct {
	Group     string `json:"group"`
	Version   string `json:"version"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func (d *DeleteResource) GetGroupVersionKind() string {
	return d.Group + "/" + d.Version + ", Kind=" + d.Kind
}

func (d *DeleteResource) GetGroup() string {
	return d.Group
}

func (d *DeleteResource) GetVersion() string {
	return d.Version
}

func (d *DeleteResource) GetKind() string {
	return d.Kind
}

func (d *DeleteResource) GetName() string {
	return d.Name
}

func (d *DeleteResource) GetNamespace() string {
	return d.Namespace
}

func (d *DeleteResource) GetDeleted() bool {
	return false
}

func (d *DeleteResource) SetDeleted(_ bool) {}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
