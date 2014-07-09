package ngsi10

type Attribute struct {
	Name	string
	Type	string
	Value	string
}

func (a Attribute) ToString() string {
	return "attribute"
}

type ContextElement struct {
	Id			string
	Type		string
	IsPattern	string
 	Attributes	[]Attribute
}

func (ce ContextElement) ToString() string {
	return "contextElement"
}

type UpdateContextRequest struct {
	ContextElements	[]ContextElement
	UpdateAction	string
}

func (ucr UpdateContextRequest) ToString() (out string) {
	out = "{\n"
// 	out += ucr.ContextElements[0].ToString()
	
	if len(ucr.ContextElements) > 1 {
		out += "\t\"contextElements\": [\n"
		for _, ce := range ucr.ContextElements {
		//for i := 0; i < len(ucr.ContextElements); i++ {
			out += "\t" + ce.ToString()
		}
		out += "\t]\n"
	}
	
	out += "\t\"updateAction\": \"" + ucr.UpdateAction +"\"\n}"
	
	return out
}
