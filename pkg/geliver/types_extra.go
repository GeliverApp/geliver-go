package geliver

// Hand-authored typed models for endpoints not present in OpenAPI output

type ProviderAccount struct {
    ID                string                 `json:"id,omitempty"`
    CreatedAt         string                 `json:"createdAt,omitempty"`
    UpdatedAt         string                 `json:"updatedAt,omitempty"`
    ProviderCode      string                 `json:"providerCode,omitempty"`
    Username          string                 `json:"username,omitempty"`
    Name              string                 `json:"name,omitempty"`
    IsActive          bool                   `json:"isActive,omitempty"`
    Parameters        map[string]any         `json:"parameters,omitempty"`
    Version           int                    `json:"version,omitempty,string"`
    IsC2C             bool                   `json:"isC2C,omitempty"`
    IntegrationType   string                 `json:"integrationType,omitempty"`
    LabelFileType     string                 `json:"labelFileType,omitempty"`
    Sharable          bool                   `json:"sharable,omitempty"`
    IsPublic          bool                   `json:"isPublic,omitempty"`
    IsDynamicPrice    bool                   `json:"isDynamicPrice,omitempty"`
    PriceUpdatedAt    string                 `json:"priceUpdatedAt,omitempty"`
}

type ProviderAccountRequest struct {
    Username        string         `json:"username"`
    Password        string         `json:"password,omitempty"`
    Name            string         `json:"name"`
    ProviderCode    string         `json:"providerCode"`
    Version         int            `json:"version"`
    IsActive        bool           `json:"isActive"`
    Parameters      map[string]any `json:"parameters,omitempty"`
    IsPublic        bool           `json:"isPublic"`
    Sharable        bool           `json:"sharable"`
    IsDynamicPrice  bool           `json:"isDynamicPrice"`
}

type ParcelTemplate struct {
    ID           string `json:"id,omitempty"`
    CreatedAt    string `json:"createdAt,omitempty"`
    UpdatedAt    string `json:"updatedAt,omitempty"`
    Length       string `json:"length,omitempty"`
    Width        string `json:"width,omitempty"`
    Height       string `json:"height,omitempty"`
    Desi         string `json:"desi,omitempty"`
    OldDesi      string `json:"oldDesi,omitempty"`
    DistanceUnit string `json:"distanceUnit,omitempty"`
    Weight       string `json:"weight,omitempty"`
    OldWeight    string `json:"oldWeight,omitempty"`
    MassUnit     string `json:"massUnit,omitempty"`
    IsActive     bool   `json:"isActive,omitempty"`
    Name         string `json:"name,omitempty"`
    LanguageCode string `json:"LanguageCode,omitempty"`
}

type Transaction struct {
    ID                 string    `json:"id,omitempty"`
    CreatedAt          string    `json:"createdAt,omitempty"`
    UpdatedAt          string    `json:"updatedAt,omitempty"`
    Amount             string    `json:"amount,omitempty"`
    Currency           string    `json:"currency,omitempty"`
    AmountLocal        string    `json:"amountLocal,omitempty"`
    CurrencyLocal      string    `json:"currencyLocal,omitempty"`
    AmountVat          string    `json:"amountVat,omitempty"`
    AmountLocalVat     string    `json:"amountLocalVat,omitempty"`
    AmountTax          string    `json:"amountTax,omitempty"`
    AmountLocalTax     string    `json:"amountLocalTax,omitempty"`
    TotalAmount        string    `json:"totalAmount,omitempty"`
    TotalAmountLocal   string    `json:"totalAmountLocal,omitempty"`
    AmountOld          string    `json:"amountOld,omitempty"`
    AmountLocalOld     string    `json:"amountLocalOld,omitempty"`
    DiscountRate       string    `json:"discountRate,omitempty"`
    BonusBalance       string    `json:"bonusBalance,omitempty"`
    OfferID            string    `json:"offerID,omitempty"`
    Shipment           *Shipment `json:"shipment,omitempty"`
    Description        string    `json:"description,omitempty"`
    IsRefund           bool      `json:"isRefund,omitempty"`
    IsCustomAccountCharge bool   `json:"isCustomAccountCharge,omitempty"`
    IsPayed            bool      `json:"isPayed,omitempty"`
    PayedVia           string    `json:"payedVia,omitempty"`
    TransactionType    string    `json:"transactionType,omitempty"`
    InvoiceID          string    `json:"invoiceID,omitempty"`
    CancelDescription  string    `json:"cancelDescription,omitempty"`
    IsCanceled         bool      `json:"isCanceled,omitempty"`
    OldBalance         string    `json:"oldBalance,omitempty"`
    NewBalance         string    `json:"newBalance,omitempty"`
    InvoiceOldDebt     string    `json:"invoiceOldDebt,omitempty"`
    InvoiceNewDebt     string    `json:"invoiceNewDebt,omitempty"`
}

type OrganizationBalance struct {
    Result            bool   `json:"result,omitempty"`
    AdditionalMessage string `json:"additionalMessage,omitempty"`
    Data              string `json:"data,omitempty"`
    Debt              string `json:"debt,omitempty"`
}

