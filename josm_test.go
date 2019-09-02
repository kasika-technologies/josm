package josm

import (
	"bytes"
	"testing"
)

func TestUnmarshalData(t *testing.T) {
	root, err := Decode(bytes.NewReader(testData))
	if err != nil {
		t.Fatal(err)
	}

	if len(root.Nodes) == 0 {
		t.Fatal("Root must have at least one node.")
	}
}

var testData = []byte(`<?xml version='1.0' encoding='UTF-8'?>
<osm version='0.6' generator='JOSM'>
  <bounds minlat='51.5076478723889' minlon='-0.127989783553507' maxlat='51.5077445145483' maxlon='-0.127774884645096' origin='OpenStreetMap server' />
  <node id='26821100' timestamp='2009-02-16T21:34:57+00:00' user='dankarran' visible='true' lat='51.5077286' lon='-0.1279688'>
    <tag k='created_by' v='Potlatch 0.10f' />
    <tag k='name' v='Nelson&apos;s Column' />
    <tag k='tourism' v='attraction' />
    <tag k='monument' v='statue' />
    <tag k='historic' v='monument' />
  </node>
  <node id='-1' visible='true' lat='51.507661490456606' lon='-0.1278000843634869' />
  <node id='346364767' action='delete' timestamp='2009-02-16T21:34:44+00:00' user='dankarran' visible='true' lat='51.5076698' lon='-0.1278143' />
</osm>`)
