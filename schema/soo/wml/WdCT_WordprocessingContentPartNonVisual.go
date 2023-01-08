// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"github.com/sdowns/gooxml"
	"github.com/sdowns/gooxml/schema/soo/dml"
)

type WdCT_WordprocessingContentPartNonVisual struct {
	CNvPr            *dml.CT_NonVisualDrawingProps
	CNvContentPartPr *dml.CT_NonVisualContentPartProperties
}

func NewWdCT_WordprocessingContentPartNonVisual() *WdCT_WordprocessingContentPartNonVisual {
	ret := &WdCT_WordprocessingContentPartNonVisual{}
	return ret
}

func (m *WdCT_WordprocessingContentPartNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CNvPr != nil {
		secNvPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvPr"}}
		e.EncodeElement(m.CNvPr, secNvPr)
	}
	if m.CNvContentPartPr != nil {
		secNvContentPartPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvContentPartPr"}}
		e.EncodeElement(m.CNvContentPartPr, secNvContentPartPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_WordprocessingContentPartNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_WordprocessingContentPartNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvPr"}:
				m.CNvPr = dml.NewCT_NonVisualDrawingProps()
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvContentPartPr"}:
				m.CNvContentPartPr = dml.NewCT_NonVisualContentPartProperties()
				if err := d.DecodeElement(m.CNvContentPartPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdCT_WordprocessingContentPartNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WordprocessingContentPartNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WordprocessingContentPartNonVisual and its children
func (m *WdCT_WordprocessingContentPartNonVisual) Validate() error {
	return m.ValidateWithPath("WdCT_WordprocessingContentPartNonVisual")
}

// ValidateWithPath validates the WdCT_WordprocessingContentPartNonVisual and its children, prefixing error messages with path
func (m *WdCT_WordprocessingContentPartNonVisual) ValidateWithPath(path string) error {
	if m.CNvPr != nil {
		if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
			return err
		}
	}
	if m.CNvContentPartPr != nil {
		if err := m.CNvContentPartPr.ValidateWithPath(path + "/CNvContentPartPr"); err != nil {
			return err
		}
	}
	return nil
}
