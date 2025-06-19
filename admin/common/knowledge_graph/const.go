package knowledge_graph

type Entities struct {
	Entities []Entity `json:"entities"`
}

type Entity struct {
	ID         int    `json:"id"`
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

type Relationships struct {
	Relationships []Relationship `json:"relationships"`
}

type Relationship struct {
	Source string `json:"source"`
	Rel    string `json:"rel"`
	Target string `json:"target"`
}
