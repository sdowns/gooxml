// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/sdowns/gooxml"
)

type CT_SmartTags struct {
	// Cell Smart Tags
	CellSmartTags []*CT_CellSmartTags
}

func NewCT_SmartTags() *CT_SmartTags {
	ret := &CT_SmartTags{}
	return ret
}

func (m *CT_SmartTags) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secellSmartTags := xml.StartElement{Name: xml.Name{Local: "ma:cellSmartTags"}}
	for _, c := range m.CellSmartTags {
		e.EncodeElement(c, secellSmartTags)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SmartTags) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SmartTags:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cellSmartTags"}:
				tmp := NewCT_CellSmartTags()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CellSmartTags = append(m.CellSmartTags, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_SmartTags %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SmartTags
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SmartTags and its children
func (m *CT_SmartTags) Validate() error {
	return m.ValidateWithPath("CT_SmartTags")
}

// ValidateWithPath validates the CT_SmartTags and its children, prefixing error messages with path
func (m *CT_SmartTags) ValidateWithPath(path string) error {
	for i, v := range m.CellSmartTags {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CellSmartTags[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
