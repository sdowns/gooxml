// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"log"

	"github.com/sdowns/gooxml/common"
	"github.com/sdowns/gooxml/document"
	"github.com/sdowns/gooxml/measurement"

	"github.com/sdowns/gooxml/schema/soo/wml"
)

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()

	img, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	iref, err := doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}

	para := doc.AddParagraph()
	anchored, err := para.AddRun().AddDrawingAnchored(iref)
	if err != nil {
		log.Fatalf("unable to add anchored image: %s", err)
	}
	anchored.SetName("Gopher")
	anchored.SetSize(2*measurement.Inch, 2*measurement.Inch)
	anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
	anchored.SetHAlignment(wml.WdST_AlignHCenter)
	anchored.SetYOffset(3 * measurement.Inch)
	anchored.SetTextWrapSquare(wml.WdST_WrapTextBothSides)

	run := para.AddRun()
	for i := 0; i < 16; i++ {
		run.AddText(lorem)

		// drop an inline image in
		if i == 13 {
			inl, err := run.AddDrawingInline(iref)
			if err != nil {
				log.Fatalf("unable to add inline image: %s", err)
			}
			inl.SetSize(1*measurement.Inch, 1*measurement.Inch)
		}
	}
	doc.SaveToFile("image.docx")
}
