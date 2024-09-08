package model

type GetProductsResponse struct {
	Products []*Product `json:"products"`
}

type GetProductVariantsResponse struct {
	Variants []*Variant `json:"variants"`
}

type Product struct {
	ID                int64       `json:"id"`
	Title             string      `json:"title"`
	BodyHTML          string      `json:"body_html"`
	Vendor            string      `json:"vendor"`
	ProductType       string      `json:"product_type"`
	CreatedAt         string      `json:"created_at"`
	Handle            string      `json:"handle"`
	UpdatedAt         string      `json:"updated_at"`
	PublishedAt       string      `json:"published_at"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	PublishedScope    string      `json:"published_scope"`
	Tags              string      `json:"tags"`
	Status            string      `json:"status"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	Variants          []*Variant  `json:"variants"`
	Options           []*Option   `json:"options"`
	Images            []*Image    `json:"images"`
	Image             *Image      `json:"image"`
}

type Image struct {
	ID                int64       `json:"id"`
	Alt               interface{} `json:"alt"`
	Position          int64       `json:"position"`
	ProductID         int64       `json:"product_id"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	Width             int64       `json:"width"`
	Height            int64       `json:"height"`
	Src               *string     `json:"src"`
	VariantIDS        []int64     `json:"variant_ids"`
}

type Option struct {
	ID        int64    `json:"id"`
	ProductID int64    `json:"product_id"`
	Name      string   `json:"name"`
	Position  int64    `json:"position"`
	Values    []string `json:"values"`
}

type Variant struct {
	ID                   int64               `json:"id"`
	ProductID            int64               `json:"product_id"`
	Title                string              `json:"title"`
	Price                string              `json:"price"`
	Sku                  string              `json:"sku"`
	Position             int64               `json:"position"`
	InventoryPolicy      string              `json:"inventory_policy"`
	CompareAtPrice       interface{}         `json:"compare_at_price"`
	FulfillmentService   string              `json:"fulfillment_service"`
	InventoryManagement  string              `json:"inventory_management"`
	Option1              string              `json:"option1"`
	Option2              string              `json:"option2"`
	Option3              string              `json:"option3"`
	CreatedAt            string              `json:"created_at"`
	UpdatedAt            string              `json:"updated_at"`
	Taxable              bool                `json:"taxable"`
	Barcode              string              `json:"barcode"`
	Grams                int64               `json:"grams"`
	Weight               float64             `json:"weight"`
	WeightUnit           string              `json:"weight_unit"`
	InventoryItemID      int64               `json:"inventory_item_id"`
	InventoryQuantity    int64               `json:"inventory_quantity"`
	OldInventoryQuantity int64               `json:"old_inventory_quantity"`
	PresentmentPrices    []*PresentmentPrice `json:"presentment_prices"`
	RequiresShipping     bool                `json:"requires_shipping"`
	AdminGraphqlAPIID    string              `json:"admin_graphql_api_id"`
	ImageID              *int64              `json:"image_id"`
}

type PresentmentPrice struct {
	Price          Price       `json:"price"`
	CompareAtPrice interface{} `json:"compare_at_price"`
}

type Price struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}
