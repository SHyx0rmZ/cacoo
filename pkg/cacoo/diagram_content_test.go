package cacoo

import (
	"encoding/xml"
	"testing"
)

func TestClient_DiagramContentPlainText(t *testing.T) {
	var c DiagramContent
	err := xml.Unmarshal([]byte(`<?xml version="1.0" encoding="utf-8"?>
<diagram>
  <sheet name="test">
    <group attr-stencil-id="10000">
      <polygon/>
      <text>Plain text</text>
    </group>
  </sheet>
</diagram>`), &c)
	if err != nil {
		t.Fatal(err)
	}
	if len(c.Sheets) != 1 {
		t.Fatal("failed to parse sheets")
	}
	if len(c.Sheets[0].Groups) != 1 {
		t.Fatal("failed to parse group")
	}
	if len(c.Sheets[0].Groups[0].Texts) != 1 {
		t.Fatal("failed to parse texts")
	}
	if c.Sheets[0].Groups[0].Texts[0].Text != "Plain text" {
		t.Fatal("failed to parse plain text")
	}
}

func TestClient_DiagramContentStyledText(t *testing.T) {
	var c DiagramContent
	err := xml.Unmarshal([]byte(`<?xml version="1.0" encoding="utf-8"?>
<diagram>
  <sheet name="test">
    <group attr-stencil-id="10000">
      <polygon/>
      <text>
        <textStyle font="Arial">Styled text</textStyle>
      </text>
    </group>
  </sheet>
</diagram>`), &c)
	if err != nil {
		t.Fatal(err)
	}
	if len(c.Sheets) != 1 {
		t.Fatal("failed to parse sheets")
	}
	if len(c.Sheets[0].Groups) != 1 {
		t.Fatal("failed to parse group")
	}
	if len(c.Sheets[0].Groups[0].Texts) != 1 {
		t.Fatal("failed to parse texts")
	}
	if c.Sheets[0].Groups[0].Texts[0].TextStyle.Font != "Arial" {
		t.Fatal("failed to parse text style")
	}
	if c.Sheets[0].Groups[0].Texts[0].TextStyle.Text != "Styled text" {
		t.Fatal("failed to parse styled text")
	}
}
