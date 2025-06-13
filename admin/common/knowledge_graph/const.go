package knowledge_graph

type Entities struct {
	Entities []Entity `json:"entities"`
}

type Entity struct {
	EntityName string `json:"entity"`
	Type       string `json:"type"`
}

type Props struct {
	Props []Prop `json:"props"`
}

type Prop struct {
	PropName string `json:"prop"`
	Value    string `json:"value"`
}

type Triple struct {
	Source string `json:"source"`
	Rel    string `json:"rel"`
	Target string `json:"target"`
}
