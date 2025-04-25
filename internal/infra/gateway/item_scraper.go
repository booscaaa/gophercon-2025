package gateway

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/playwright-community/playwright-go"
)

type itemScraperGateway struct {
}

// ScrapeItems implements domain.ItemScraperGateway.
func (gateway *itemScraperGateway) ScrapeItems(ctx context.Context) ([]domain.Item, error) {
	items := []domain.Item{}

	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(
		playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(true),
		},
	)
	if err != nil {
		return nil, err
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		return nil, err
	}

	_, err = page.Goto("https://potatosburger.menudino.com")
	if err != nil {
		return nil, err
	}

	categories, err := page.Locator("#cardapio > section.cardapio-body > div > div.categories > div").All()
	if err != nil {
		return nil, err
	}
	for _, category := range categories {
		category.ScrollIntoViewIfNeeded()

		time.Sleep(time.Millisecond * 500)

		cards, err := category.Locator("div:nth-child(2) > div > div").All()
		if err != nil {
			return nil, err
		}

		for _, card := range cards {
			card.ScrollIntoViewIfNeeded()

			productName, err := card.Locator("a > div > div.media-body > div.name > span").TextContent()
			if err != nil {
				return nil, err
			}

			productPrice, err := card.Locator("a > div > div.media-body > div.priceDescription > div").TextContent()
			if err != nil {
				return nil, err
			}

			productPrice = strings.ReplaceAll(productPrice, "R$ ", "")
			productPrice = strings.ReplaceAll(productPrice, ",", ".")

			product := domain.Item{
				Name: productName,
			}

			if price, err := strconv.ParseFloat(productPrice, 64); err == nil {
				product.Price = price
			}

			items = append(items, product)
		}
	}

	return items, nil
}

func NewItemScraper() domain.ItemScraperGateway {
	return &itemScraperGateway{}
}
