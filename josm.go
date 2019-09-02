package josm

import (
	"encoding/xml"
	"io"
	"time"
)

type Root struct {
	Bounds    Bounds
	Nodes     []Node
	Ways      []Way
	Relations []Relation
}

type Bounds struct {
	XMLName xml.Name `xml:"bounds"`
	Minlat  float64  `xml:"minlat,attr"`
	Minlon  float64  `xml:"minlon,attr"`
	Maxlat  float64  `xml:"maxlat,attr"`
	Maxlon  float64  `xml:"maxlon,attr"`
}

type Element struct {
	ID        int64 `xml:"id,attr"`
	Location  Location
	Version   int       `xml:"version,attr"`
	Timestamp time.Time `xml:"timestamp,attr"`
	UID       int64     `xml:"uid,attr"`
	User      string    `xml:"user,attr"`
	ChangeSet int64     `xml:"changeset,attr"`
}

type Location struct {
	Type        string
	Coordinates []float64
}

type Member struct {
	Type string `xml:"type,attr"`
	Ref  int64  `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}

type Node struct {
	Element
	XMLName xml.Name `xml:"node"`
	Lat     float64  `xml:"lat,attr"`
	Lon     float64  `xml:"lon,attr"`
	Tag     []Tag    `xml:"tag"`
}

type Way struct {
	Element
	XMLName xml.Name `xml:"way"`
	Tags    []Tag    `xml:"tag"`
	Nds     []struct {
		ID int64 `xml:"ref,attr"`
	} `xml:"nd"`
}

type Relation struct {
	Element
	Visible bool     `xml:"visible,attr"`
	Version string   `xml:"version,attr"`
	Members []Member `xml:"member"`
	Tags    []Tag    `xml:"tag"`
}

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Key     string   `xml:"k,attr"`
	Value   string   `xml:"v,attr"`
}

func Decode(reader io.Reader) (*Root, error) {
	var (
		r   = new(Root)
		err error
	)

	decoder := xml.NewDecoder(reader)

	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch typedToken := token.(type) {
		case xml.StartElement:
			switch typedToken.Name.Local {
			case "bounds":
				var b Bounds
				err = decoder.DecodeElement(&b, &typedToken)
				if err != nil {
					return nil, err
				}
				r.Bounds = b
			case "node":
				var node Node
				err = decoder.DecodeElement(&node, &typedToken)
				if err != nil {
					return nil, err
				}
				r.Nodes = append(r.Nodes, node)
			case "way":
				var way Way
				err = decoder.DecodeElement(&way, &typedToken)
				if err != nil {
					return nil, err
				}
				r.Ways = append(r.Ways, way)
			case "relation":
				var relation Relation
				err = decoder.DecodeElement(&relation, &typedToken)
				if err != nil {
					return nil, err
				}
				r.Relations = append(r.Relations, relation)
			}
		}
	}

	return r, nil
}
