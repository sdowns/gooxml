// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/sdowns/gooxml/schema/soo/dml/diagram"
)

func TestCT_ParameterConstructor(t *testing.T) {
	v := diagram.NewCT_Parameter()
	if v == nil {
		t.Errorf("diagram.NewCT_Parameter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Parameter should validate: %s", err)
	}
}

func TestCT_ParameterMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Parameter()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Parameter()
	xml.Unmarshal(buf, v2)
}
