// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "github.com/sdowns/gooxml/schema/soo/wml"

// Footer is a footer for a document section.
type Footer struct {
	d *Document
	x *wml.Ftr
}

// X returns the inner wrapped XML type.
func (f Footer) X() *wml.Ftr {
	return f.x
}

// AddParagraph adds a paragraph to the footer.
func (f Footer) AddParagraph() Paragraph {
	bc := wml.NewEG_ContentBlockContent()
	f.x.EG_ContentBlockContent = append(f.x.EG_ContentBlockContent, bc)
	p := wml.NewCT_P()
	bc.P = append(bc.P, p)
	return Paragraph{f.d, p}
}

// Index returns the index of the footer within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (f Footer) Index() int {
	for i, hdr := range f.d.footers {
		if hdr == f.x {
			return i
		}
	}
	return -1
}

// Paragraphs returns the paragraphs defined in a footer.
func (f Footer) Paragraphs() []Paragraph {
	ret := []Paragraph{}
	for _, ec := range f.x.EG_ContentBlockContent {
		for _, p := range ec.P {
			ret = append(ret, Paragraph{f.d, p})
		}
	}
	return ret
}

// RemoveParagraph removes a paragraph from a footer.
func (f Footer) RemoveParagraph(p Paragraph) {
	for _, ec := range f.x.EG_ContentBlockContent {
		for i, pa := range ec.P {
			// do we need to remove this paragraph
			if pa == p.x {
				copy(ec.P[i:], ec.P[i+1:])
				ec.P = ec.P[0 : len(ec.P)-1]
				return
			}
		}
	}
}

// Clear clears all content within a footer
func (f Footer) Clear() {
	f.x.EG_ContentBlockContent = nil
}
