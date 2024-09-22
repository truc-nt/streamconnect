package model

type CreateOrderRequest struct {
	Order *OrderRequest `json:"order"`
}

type OrderRequest struct {
	LineItems       []*LineItemRequest `json:"line_items"`
	Customer        *CustomerRequest   `json:"customer"`
	BillingAddress  *BillingAddress    `json:"billing_address"`
	ShippingAddress *ShippingAddress   `json:"shipping_address"`
	Email           string             `json:"email"`
	/*Transactions    []*Transactions    `json:"transactions"`*/
	FinancialStatus string          `json:"financial_status"`
	DiscountCodes   []*DiscountCode `json:"discount_codes"`
}

type LineItemRequest struct {
	VariantID int   `json:"variant_id"`
	Quantity  int64 `json:"quantity"`
}

type CustomerRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type BillingAddress struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address1"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Zip       string `json:"zip"`
}
type ShippingAddress struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address1"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Zip       string `json:"zip"`
}
type Transactions struct {
	Kind   string  `json:"kind"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

type DiscountCode struct {
	Code   string `json:"code"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type CreateOrderResponse struct {
	Order *OrderResponse `json:"order"`
}

type OrderResponse struct {
	ID                             int64         `json:"id"`
	AdminGraphqlAPIID              string        `json:"admin_graphql_api_id"`
	AppID                          int64         `json:"app_id"`
	BrowserIP                      interface{}   `json:"browser_ip"`
	BuyerAcceptsMarketing          bool          `json:"buyer_accepts_marketing"`
	CancelReason                   interface{}   `json:"cancel_reason"`
	CancelledAt                    interface{}   `json:"cancelled_at"`
	CartToken                      interface{}   `json:"cart_token"`
	CheckoutID                     interface{}   `json:"checkout_id"`
	CheckoutToken                  interface{}   `json:"checkout_token"`
	ClientDetails                  interface{}   `json:"client_details"`
	ClosedAt                       interface{}   `json:"closed_at"`
	ConfirmationNumber             string        `json:"confirmation_number"`
	Confirmed                      bool          `json:"confirmed"`
	ContactEmail                   string        `json:"contact_email"`
	CreatedAt                      string        `json:"created_at"`
	Currency                       Currency      `json:"currency"`
	CurrentSubtotalPrice           string        `json:"current_subtotal_price"`
	CurrentSubtotalPriceSet        Set           `json:"current_subtotal_price_set"`
	CurrentTotalAdditionalFeesSet  interface{}   `json:"current_total_additional_fees_set"`
	CurrentTotalDiscounts          string        `json:"current_total_discounts"`
	CurrentTotalDiscountsSet       Set           `json:"current_total_discounts_set"`
	CurrentTotalDutiesSet          interface{}   `json:"current_total_duties_set"`
	CurrentTotalPrice              string        `json:"current_total_price"`
	CurrentTotalPriceSet           Set           `json:"current_total_price_set"`
	CurrentTotalTax                string        `json:"current_total_tax"`
	CurrentTotalTaxSet             Set           `json:"current_total_tax_set"`
	CustomerLocale                 interface{}   `json:"customer_locale"`
	DeviceID                       interface{}   `json:"device_id"`
	DiscountCodes                  []interface{} `json:"discount_codes"`
	DutiesIncluded                 bool          `json:"duties_included"`
	Email                          string        `json:"email"`
	EstimatedTaxes                 bool          `json:"estimated_taxes"`
	FinancialStatus                string        `json:"financial_status"`
	FulfillmentStatus              interface{}   `json:"fulfillment_status"`
	LandingSite                    interface{}   `json:"landing_site"`
	LandingSiteRef                 interface{}   `json:"landing_site_ref"`
	LocationID                     interface{}   `json:"location_id"`
	MerchantOfRecordAppID          interface{}   `json:"merchant_of_record_app_id"`
	Name                           string        `json:"name"`
	Note                           interface{}   `json:"note"`
	NoteAttributes                 []interface{} `json:"note_attributes"`
	Number                         int64         `json:"number"`
	OrderNumber                    int64         `json:"order_number"`
	OrderStatusURL                 string        `json:"order_status_url"`
	OriginalTotalAdditionalFeesSet interface{}   `json:"original_total_additional_fees_set"`
	OriginalTotalDutiesSet         interface{}   `json:"original_total_duties_set"`
	PaymentGatewayNames            []string      `json:"payment_gateway_names"`
	Phone                          interface{}   `json:"phone"`
	PoNumber                       interface{}   `json:"po_number"`
	PresentmentCurrency            Currency      `json:"presentment_currency"`
	ProcessedAt                    string        `json:"processed_at"`
	Reference                      interface{}   `json:"reference"`
	ReferringSite                  interface{}   `json:"referring_site"`
	SourceIdentifier               interface{}   `json:"source_identifier"`
	SourceName                     string        `json:"source_name"`
	SourceURL                      interface{}   `json:"source_url"`
	SubtotalPrice                  string        `json:"subtotal_price"`
	SubtotalPriceSet               Set           `json:"subtotal_price_set"`
	Tags                           string        `json:"tags"`
	TaxExempt                      bool          `json:"tax_exempt"`
	TaxLines                       []interface{} `json:"tax_lines"`
	TaxesIncluded                  bool          `json:"taxes_included"`
	Test                           bool          `json:"test"`
	Token                          string        `json:"token"`
	TotalDiscounts                 string        `json:"total_discounts"`
	TotalDiscountsSet              Set           `json:"total_discounts_set"`
	TotalLineItemsPrice            string        `json:"total_line_items_price"`
	TotalLineItemsPriceSet         Set           `json:"total_line_items_price_set"`
	TotalOutstanding               string        `json:"total_outstanding"`
	TotalPrice                     string        `json:"total_price"`
	TotalPriceSet                  Set           `json:"total_price_set"`
	TotalShippingPriceSet          Set           `json:"total_shipping_price_set"`
	TotalTax                       string        `json:"total_tax"`
	TotalTaxSet                    Set           `json:"total_tax_set"`
	TotalTipReceived               string        `json:"total_tip_received"`
	TotalWeight                    int64         `json:"total_weight"`
	UpdatedAt                      string        `json:"updated_at"`
	UserID                         interface{}   `json:"user_id"`
	BillingAddress                 Address       `json:"billing_address"`
	Customer                       Customer      `json:"customer"`
	DiscountApplications           []interface{} `json:"discount_applications"`
	Fulfillments                   []interface{} `json:"fulfillments"`
	LineItems                      []LineItem    `json:"line_items"`
	PaymentTerms                   interface{}   `json:"payment_terms"`
	Refunds                        []interface{} `json:"refunds"`
	ShippingAddress                Address       `json:"shipping_address"`
	ShippingLines                  []interface{} `json:"shipping_lines"`
}

type Address struct {
	FirstName    string      `json:"first_name"`
	Address1     string      `json:"address1"`
	Phone        string      `json:"phone"`
	City         string      `json:"city"`
	Zip          string      `json:"zip"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	LastName     string      `json:"last_name"`
	Address2     interface{} `json:"address2"`
	Company      interface{} `json:"company"`
	Latitude     interface{} `json:"latitude"`
	Longitude    interface{} `json:"longitude"`
	Name         string      `json:"name"`
	CountryCode  string      `json:"country_code"`
	ProvinceCode string      `json:"province_code"`
	ID           *int64      `json:"id,omitempty"`
	CustomerID   *int64      `json:"customer_id,omitempty"`
	CountryName  *string     `json:"country_name,omitempty"`
	Default      *bool       `json:"default,omitempty"`
}

type Set struct {
	ShopMoney        Money `json:"shop_money"`
	PresentmentMoney Money `json:"presentment_money"`
}

type Money struct {
	Amount       string   `json:"amount"`
	CurrencyCode Currency `json:"currency_code"`
}

type Customer struct {
	ID                    int64                 `json:"id"`
	Email                 string                `json:"email"`
	CreatedAt             string                `json:"created_at"`
	UpdatedAt             string                `json:"updated_at"`
	FirstName             string                `json:"first_name"`
	LastName              string                `json:"last_name"`
	State                 string                `json:"state"`
	Note                  interface{}           `json:"note"`
	VerifiedEmail         bool                  `json:"verified_email"`
	MultipassIdentifier   interface{}           `json:"multipass_identifier"`
	TaxExempt             bool                  `json:"tax_exempt"`
	Phone                 interface{}           `json:"phone"`
	EmailMarketingConsent EmailMarketingConsent `json:"email_marketing_consent"`
	SMSMarketingConsent   interface{}           `json:"sms_marketing_consent"`
	Tags                  string                `json:"tags"`
	Currency              Currency              `json:"currency"`
	TaxExemptions         []interface{}         `json:"tax_exemptions"`
	AdminGraphqlAPIID     string                `json:"admin_graphql_api_id"`
	DefaultAddress        Address               `json:"default_address"`
}

type EmailMarketingConsent struct {
	State            string      `json:"state"`
	OptInLevel       string      `json:"opt_in_level"`
	ConsentUpdatedAt interface{} `json:"consent_updated_at"`
}

type LineItem struct {
	ID                         int64         `json:"id"`
	AdminGraphqlAPIID          string        `json:"admin_graphql_api_id"`
	AttributedStaffs           []interface{} `json:"attributed_staffs"`
	CurrentQuantity            int64         `json:"current_quantity"`
	FulfillableQuantity        int64         `json:"fulfillable_quantity"`
	FulfillmentService         string        `json:"fulfillment_service"`
	FulfillmentStatus          interface{}   `json:"fulfillment_status"`
	GiftCard                   bool          `json:"gift_card"`
	Grams                      int64         `json:"grams"`
	Name                       string        `json:"name"`
	Price                      string        `json:"price"`
	PriceSet                   Set           `json:"price_set"`
	ProductExists              bool          `json:"product_exists"`
	ProductID                  int64         `json:"product_id"`
	Properties                 []interface{} `json:"properties"`
	Quantity                   int64         `json:"quantity"`
	RequiresShipping           bool          `json:"requires_shipping"`
	Sku                        string        `json:"sku"`
	Taxable                    bool          `json:"taxable"`
	Title                      string        `json:"title"`
	TotalDiscount              string        `json:"total_discount"`
	TotalDiscountSet           Set           `json:"total_discount_set"`
	VariantID                  int64         `json:"variant_id"`
	VariantInventoryManagement string        `json:"variant_inventory_management"`
	VariantTitle               string        `json:"variant_title"`
	Vendor                     string        `json:"vendor"`
	TaxLines                   []interface{} `json:"tax_lines"`
	Duties                     []interface{} `json:"duties"`
	DiscountAllocations        []interface{} `json:"discount_allocations"`
}

type Currency string

const (
	Usd Currency = "USD"
)
