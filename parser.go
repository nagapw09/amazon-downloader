package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type AmazonParser struct {
	body     string
	product  *Product
	document *goquery.Document
}

func (amp *AmazonParser) Parse() error {
	var err error
	amp.document, err = goquery.NewDocumentFromReader(strings.NewReader(amp.body))
	if err != nil {
		return err
	}

	amp.product.Title = amp.getTitle()
	amp.product.Image = amp.getImage()
	amp.product.IsSale = amp.isAvailableForSale()
	if !amp.product.IsSale {
		amp.product.Price = "None"
	}else{
		amp.product.Price = amp.getPrice()
	}

	return nil
}

func (amp *AmazonParser) getTitle() string {
	text := amp.document.Find("span#productTitle").Text()
	return strings.TrimSpace(text)
}

func (amp *AmazonParser) getPrice() string {
	price := amp.document.Find("span.a-color-price").First().Text()
	return strings.TrimSpace(strings.TrimSpace(price))
}

func (amp *AmazonParser) getImage() string {
	imgContainerSelector := amp.document.Find("#imageBlock").Find("#main-image-container").Find("img").First()
	value, ok := imgContainerSelector.Attr("data-old-hires")
	if !ok{
		value, _ = imgContainerSelector.Attr("src")
	}
	return strings.TrimSpace(value)
}

func (amp *AmazonParser) isAvailableForSale() bool {
	return !strings.Contains(amp.body, "Currently unavailable")
}
