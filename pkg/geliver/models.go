// Auto-generated from openapi.yaml
package geliver

// Address model
type Address struct {
  Address1 string `json:"address1,omitempty"`
  Address2 string `json:"address2,omitempty"`
  City *City `json:"city,omitempty"`
  CityCode string `json:"cityCode,omitempty"`
  CityName string `json:"cityName,omitempty"`
  CountryCode string `json:"countryCode,omitempty"`
  CountryName string `json:"countryName,omitempty"`
  CreatedAt string `json:"createdAt,omitempty"`
  District *District `json:"district,omitempty"`
  DistrictID int `json:"districtID,omitempty"`
  DistrictName string `json:"districtName,omitempty"`
  Email string `json:"email,omitempty"`
  ID string `json:"id,omitempty"`
  IsActive bool `json:"isActive,omitempty"`
  IsDefaultReturnAddress bool `json:"isDefaultReturnAddress,omitempty"`
  IsDefaultSenderAddress bool `json:"isDefaultSenderAddress,omitempty"`
  IsInvoiceAddress bool `json:"isInvoiceAddress,omitempty"`
  IsRecipientAddress bool `json:"isRecipientAddress,omitempty"`
  Metadata *JSONContent `json:"metadata,omitempty"`
  Name string `json:"name,omitempty"`
  Owner string `json:"owner,omitempty"`
  Phone string `json:"phone,omitempty"`
  ShortName string `json:"shortName,omitempty"`
  Source string `json:"source,omitempty"`
  State string `json:"state,omitempty"`
  StreetID string `json:"streetID,omitempty"`
  StreetName string `json:"streetName,omitempty"`
  Test bool `json:"test,omitempty"`
  UpdatedAt string `json:"updatedAt,omitempty"`
  Zip string `json:"zip,omitempty"`
}

// City model
type City struct {
  AreaCode string `json:"areaCode,omitempty"`
  CityCode string `json:"cityCode,omitempty"`
  CountryCode string `json:"countryCode,omitempty"`
  // Model
  Name string `json:"name,omitempty"`
}

// DbStringArray model
type DbStringArray struct {
}

// District model
type District struct {
  CityCode string `json:"cityCode,omitempty"`
  CountryCode string `json:"countryCode,omitempty"`
  DistrictID int `json:"districtID,omitempty"`
  // Model
  Name string `json:"name,omitempty"`
  RegionCode string `json:"regionCode,omitempty"`
}

// Duration model removed; API returns integer timestamp for duration fields.

// Item model
type Item struct {
  CountryOfOrigin string `json:"countryOfOrigin,omitempty"`
  CreatedAt string `json:"createdAt,omitempty"`
  Currency string `json:"currency,omitempty"`
  CurrencyLocal string `json:"currencyLocal,omitempty"`
  ID string `json:"id,omitempty"`
  MassUnit string `json:"massUnit,omitempty"`
  MaxDeliveryTime string `json:"maxDeliveryTime,omitempty"`
  MaxShipTime string `json:"maxShipTime,omitempty"`
  Owner string `json:"owner,omitempty"`
  Quantity int `json:"quantity,omitempty"`
  Sku string `json:"sku,omitempty"`
  Test bool `json:"test,omitempty"`
  Title string `json:"title,omitempty"`
  TotalPrice string `json:"totalPrice,omitempty"`
  TotalPriceLocal string `json:"totalPriceLocal,omitempty"`
  UnitPrice string `json:"unitPrice,omitempty"`
  UnitPriceLocal string `json:"unitPriceLocal,omitempty"`
  UnitWeight string `json:"unitWeight,omitempty"`
  UpdatedAt string `json:"updatedAt,omitempty"`
  VariantTitle string `json:"variantTitle,omitempty"`
}

// JSONContent model
type JSONContent struct {
}

// Offer model
type Offer struct {
  Amount string `json:"amount,omitempty"`
  AmountLocal string `json:"amountLocal,omitempty"`
  AmountLocalOld string `json:"amountLocalOld,omitempty"`
  AmountLocalTax string `json:"amountLocalTax,omitempty"`
  AmountLocalVat string `json:"amountLocalVat,omitempty"`
  AmountOld string `json:"amountOld,omitempty"`
  AmountTax string `json:"amountTax,omitempty"`
  AmountVat string `json:"amountVat,omitempty"`
  AverageEstimatedTime *int64 `json:"averageEstimatedTime,omitempty"`
  AverageEstimatedTimeHumanReadible string `json:"averageEstimatedTimeHumanReadible,omitempty"`
  BonusBalance string `json:"bonusBalance,omitempty"`
  CreatedAt string `json:"createdAt,omitempty"`
  Currency string `json:"currency,omitempty"`
  CurrencyLocal string `json:"currencyLocal,omitempty"`
  DiscountRate string `json:"discountRate,omitempty"`
  DurationTerms string `json:"durationTerms,omitempty"`
  EstimatedArrivalTime string `json:"estimatedArrivalTime,omitempty"`
  ID string `json:"id,omitempty"`
  IntegrationType string `json:"integrationType,omitempty"`
  IsAccepted bool `json:"isAccepted,omitempty"`
  IsC2C bool `json:"isC2C,omitempty"`
  IsGlobal bool `json:"isGlobal,omitempty"`
  IsMainOffer bool `json:"isMainOffer,omitempty"`
  IsProviderAccountOffer bool `json:"isProviderAccountOffer,omitempty"`
  MaxEstimatedTime *int64 `json:"maxEstimatedTime,omitempty"`
  MinEstimatedTime *int64 `json:"minEstimatedTime,omitempty"`
  Owner string `json:"owner,omitempty"`
  PredictedDeliveryTime float64 `json:"predictedDeliveryTime,omitempty"`
  ProviderAccountID string `json:"providerAccountID,omitempty"`
  ProviderAccountName string `json:"providerAccountName,omitempty"`
  ProviderAccountOwnerType string `json:"providerAccountOwnerType,omitempty"`
  ProviderCode string `json:"providerCode,omitempty"`
  ProviderServiceCode string `json:"providerServiceCode,omitempty"`
  ProviderTotalAmount string `json:"providerTotalAmount,omitempty"`
  Rating float64 `json:"rating,omitempty"`
  ScheduleDate string `json:"scheduleDate,omitempty"`
  ShipmentTime string `json:"shipmentTime,omitempty"`
  Test bool `json:"test,omitempty"`
  TotalAmount string `json:"totalAmount,omitempty"`
  TotalAmountLocal string `json:"totalAmountLocal,omitempty"`
  UpdatedAt string `json:"updatedAt,omitempty"`
}

// OfferList model
type OfferList struct {
  AllowOfferFallback bool `json:"allowOfferFallback,omitempty"`
  Cheapest *Offer `json:"cheapest,omitempty"`
  CreatedAt string `json:"createdAt,omitempty"`
  Fastest *Offer `json:"fastest,omitempty"`
  Height string `json:"height,omitempty"`
  ItemIDs []string `json:"itemIDs,omitempty"`
  Length string `json:"length,omitempty"`
  List []Offer `json:"list,omitempty"`
  Owner string `json:"owner,omitempty"`
  ParcelIDs []string `json:"parcelIDs,omitempty"`
  ParcelTemplateID string `json:"parcelTemplateID,omitempty"`
  PercentageCompleted float64 `json:"percentageCompleted,omitempty"`
  ProviderAccountIDs []string `json:"providerAccountIDs,omitempty"`
  ProviderCodes []string `json:"providerCodes,omitempty"`
  ProviderServiceCodes []string `json:"providerServiceCodes,omitempty"`
  Test bool `json:"test,omitempty"`
  TotalOffersCompleted int `json:"totalOffersCompleted,omitempty"`
  TotalOffersRequested int `json:"totalOffersRequested,omitempty"`
  UpdatedAt string `json:"updatedAt,omitempty"`
  Weight string `json:"weight,omitempty"`
  Width string `json:"width,omitempty"`
}

// Order model
type Order struct {
  BuyerShipmentMethod string `json:"buyerShipmentMethod,omitempty"`
  BuyerShippingCost string `json:"buyerShippingCost,omitempty"`
  BuyerShippingCostCurrency string `json:"buyerShippingCostCurrency,omitempty"`
  CreatedAt string `json:"createdAt,omitempty"`
  ID string `json:"id,omitempty"`
  ItemIDs *DbStringArray `json:"itemIDs,omitempty"`
  MerchantCode string `json:"merchantCode,omitempty"`
  Notes string `json:"notes,omitempty"`
  OrderCode string `json:"orderCode,omitempty"`
  OrderNumber string `json:"orderNumber,omitempty"`
  OrderStatus string `json:"orderStatus,omitempty"`
  OrganizationID string `json:"organizationID,omitempty"`
  Owner string `json:"owner,omitempty"`
  Shipment *Shipment `json:"shipment,omitempty"`
  SourceCode string `json:"sourceCode,omitempty"`
  SourceIdentifier string `json:"sourceIdentifier,omitempty"`
  Test bool `json:"test,omitempty"`
  TotalAmount string `json:"totalAmount,omitempty"`
  TotalAmountCurrency string `json:"totalAmountCurrency,omitempty"`
  TotalTax string `json:"totalTax,omitempty"`
  UpdatedAt string `json:"updatedAt,omitempty"`
}

// Parcel model
type Parcel struct {
  Amount string `json:"amount,omitempty"`
  AmountLocal string `json:"amountLocal,omitempty"`
  AmountLocalOld string `json:"amountLocalOld,omitempty"`
  AmountLocalTax string `json:"amountLocalTax,omitempty"`
  AmountLocalVat string `json:"amountLocalVat,omitempty"`
  AmountOld string `json:"amountOld,omitempty"`
  AmountTax string `json:"amountTax,omitempty"`
  AmountVat string `json:"amountVat,omitempty"`
  Barcode string `json:"barcode,omitempty"`
  BonusBalance string `json:"bonusBalance,omitempty"`
  CommercialInvoiceURL string `json:"commercialInvoiceUrl,omitempty"`
  CreatedAt string `json:"createdAt,omitempty"`
  Currency string `json:"currency,omitempty"`
  CurrencyLocal string `json:"currencyLocal,omitempty"`
  CustomsDeclaration string `json:"customsDeclaration,omitempty"`
  // Desi of parcel
  Desi string `json:"desi,omitempty"`
  DiscountRate string `json:"discountRate,omitempty"`
  // Distance unit of parcel
  DistanceUnit string `json:"distanceUnit,omitempty"`
  Eta string `json:"eta,omitempty"`
  Extra *JSONContent `json:"extra,omitempty"`
  // Height of parcel
  Height string `json:"height,omitempty"`
  HidePackageContentOnTag bool `json:"hidePackageContentOnTag,omitempty"`
  ID string `json:"id,omitempty"`
  InvoiceGenerated bool `json:"invoiceGenerated,omitempty"`
  InvoiceID string `json:"invoiceID,omitempty"`
  IsMainParcel bool `json:"isMainParcel,omitempty"`
  ItemIDs []string `json:"itemIDs,omitempty"`
  LabelFileType string `json:"labelFileType,omitempty"`
  LabelURL string `json:"labelURL,omitempty"`
  // Length of parcel
  Length string `json:"length,omitempty"`
  // Weight unit of parcel
  MassUnit string `json:"massUnit,omitempty"`
  Metadata *JSONContent `json:"metadata,omitempty"`
  // Meta string to add additional info on your shipment/parcel
  MetadataText string `json:"metadataText,omitempty"`
  OldDesi string `json:"oldDesi,omitempty"`
  OldWeight string `json:"oldWeight,omitempty"`
  Owner string `json:"owner,omitempty"`
  ParcelReferenceCode string `json:"parcelReferenceCode,omitempty"`
  // Instead of setting parcel size manually, you can set this to a predefined Parcel Template
  ParcelTemplateID string `json:"parcelTemplateID,omitempty"`
  ProductPaymentOnDelivery bool `json:"productPaymentOnDelivery,omitempty"`
  ProviderTotalAmount string `json:"providerTotalAmount,omitempty"`
  QrCodeURL string `json:"qrCodeUrl,omitempty"`
  RefundInvoiceID string `json:"refundInvoiceID,omitempty"`
  ResponsiveLabelURL string `json:"responsiveLabelURL,omitempty"`
  ShipmentDate string `json:"shipmentDate,omitempty"`
  ShipmentID string `json:"shipmentID,omitempty"`
  StateCode string `json:"stateCode,omitempty"`
  Template string `json:"template,omitempty"`
  Test bool `json:"test,omitempty"`
  TotalAmount string `json:"totalAmount,omitempty"`
  TotalAmountLocal string `json:"totalAmountLocal,omitempty"`
  // Tracking number
  TrackingNumber string `json:"trackingNumber,omitempty"`
  TrackingStatus *Tracking `json:"trackingStatus,omitempty"`
  TrackingURL string `json:"trackingUrl,omitempty"`
  UpdatedAt string `json:"updatedAt,omitempty"`
  // If true, auto calculates total parcel size using the size of items
  UseDimensionsOfItems bool `json:"useDimensionsOfItems,omitempty"`
  // If true, auto calculates total parcel weight using the weight of items
  UseWeightOfItems bool `json:"useWeightOfItems,omitempty"`
  // Weight of parcel
  Weight string `json:"weight,omitempty"`
  // Width of parcel
  Width string `json:"width,omitempty"`
}

// Shipment model
type Shipment struct {
  AcceptedOffer *Offer `json:"acceptedOffer,omitempty"`
  AcceptedOfferID string `json:"acceptedOfferID,omitempty"`
  Amount string `json:"amount,omitempty"`
  AmountLocal string `json:"amountLocal,omitempty"`
  AmountLocalOld string `json:"amountLocalOld,omitempty"`
  AmountLocalTax string `json:"amountLocalTax,omitempty"`
  AmountLocalVat string `json:"amountLocalVat,omitempty"`
  AmountOld string `json:"amountOld,omitempty"`
  AmountTax string `json:"amountTax,omitempty"`
  AmountVat string `json:"amountVat,omitempty"`
  Barcode string `json:"barcode,omitempty"`
  BonusBalance string `json:"bonusBalance,omitempty"`
  BuyerNote string `json:"buyerNote,omitempty"`
  CancelDate string `json:"cancelDate,omitempty"`
  CategoryCode string `json:"categoryCode,omitempty"`
  CommercialInvoiceURL string `json:"commercialInvoiceUrl,omitempty"`
  CreateReturnLabel bool `json:"createReturnLabel,omitempty"`
  CreatedAt string `json:"createdAt,omitempty"`
  Currency string `json:"currency,omitempty"`
  CurrencyLocal string `json:"currencyLocal,omitempty"`
  CustomsDeclaration string `json:"customsDeclaration,omitempty"`
  // Desi of parcel
  Desi string `json:"desi,omitempty"`
  DiscountRate string `json:"discountRate,omitempty"`
  // Distance unit of parcel
  DistanceUnit string `json:"distanceUnit,omitempty"`
  EnableAutomation bool `json:"enableAutomation,omitempty"`
  Eta string `json:"eta,omitempty"`
  ExtraParcels []Parcel `json:"extraParcels,omitempty"`
  HasError bool `json:"hasError,omitempty"`
  // Height of parcel
  Height string `json:"height,omitempty"`
  HidePackageContentOnTag bool `json:"hidePackageContentOnTag,omitempty"`
  ID string `json:"id,omitempty"`
  InvoiceGenerated bool `json:"invoiceGenerated,omitempty"`
  InvoiceID string `json:"invoiceID,omitempty"`
  IsRecipientSmsActivated bool `json:"isRecipientSmsActivated,omitempty"`
  IsReturn bool `json:"isReturn,omitempty"`
  IsReturned bool `json:"isReturned,omitempty"`
  IsTrackingOnly bool `json:"isTrackingOnly,omitempty"`
  Items []Item `json:"items,omitempty"`
  LabelFileType string `json:"labelFileType,omitempty"`
  LabelURL string `json:"labelURL,omitempty"`
  LastErrorCode string `json:"lastErrorCode,omitempty"`
  LastErrorMessage string `json:"lastErrorMessage,omitempty"`
  // Length of parcel
  Length string `json:"length,omitempty"`
  // Weight unit of parcel
  MassUnit string `json:"massUnit,omitempty"`
  Metadata *JSONContent `json:"metadata,omitempty"`
  // Meta string to add additional info on your shipment/parcel
  MetadataText string `json:"metadataText,omitempty"`
  Offers *OfferList `json:"offers,omitempty"`
  OldDesi string `json:"oldDesi,omitempty"`
  OldWeight string `json:"oldWeight,omitempty"`
  Order *Order `json:"order,omitempty"`
  OrderID string `json:"orderID,omitempty"`
  OrganizationShipmentID int `json:"organizationShipmentID,omitempty"`
  Owner string `json:"owner,omitempty"`
  PackageAcceptedAt string `json:"packageAcceptedAt,omitempty"`
  // Instead of setting parcel size manually, you can set this to a predefined Parcel Template
  ParcelTemplateID string `json:"parcelTemplateID,omitempty"`
  ProductPaymentOnDelivery bool `json:"productPaymentOnDelivery,omitempty"`
  ProviderAccountID string `json:"providerAccountID,omitempty"`
  ProviderAccountIDs []string `json:"providerAccountIDs,omitempty"`
  ProviderBranchName string `json:"providerBranchName,omitempty"`
  ProviderCode string `json:"providerCode,omitempty"`
  ProviderCodes []string `json:"providerCodes,omitempty"`
  ProviderInvoiceNo string `json:"providerInvoiceNo,omitempty"`
  ProviderReceiptNo string `json:"providerReceiptNo,omitempty"`
  ProviderSerialNo string `json:"providerSerialNo,omitempty"`
  ProviderServiceCode string `json:"providerServiceCode,omitempty"`
  ProviderServiceCodes []string `json:"providerServiceCodes,omitempty"`
  ProviderTotalAmount string `json:"providerTotalAmount,omitempty"`
  QrCodeURL string `json:"qrCodeUrl,omitempty"`
  RecipientAddress *Address `json:"recipientAddress,omitempty"`
  RecipientAddressID string `json:"recipientAddressID,omitempty"`
  RefundInvoiceID string `json:"refundInvoiceID,omitempty"`
  ResponsiveLabelURL string `json:"responsiveLabelURL,omitempty"`
  ReturnAddressID string `json:"returnAddressID,omitempty"`
  SellerNote string `json:"sellerNote,omitempty"`
  SenderAddress *Address `json:"senderAddress,omitempty"`
  SenderAddressID string `json:"senderAddressID,omitempty"`
  ShipmentDate string `json:"shipmentDate,omitempty"`
  StatusCode string `json:"statusCode,omitempty"`
  Tags []string `json:"tags,omitempty"`
  TenantID string `json:"tenantId,omitempty"`
  Test bool `json:"test,omitempty"`
  TotalAmount string `json:"totalAmount,omitempty"`
  TotalAmountLocal string `json:"totalAmountLocal,omitempty"`
  // Tracking number
  TrackingNumber string `json:"trackingNumber,omitempty"`
  TrackingStatus *Tracking `json:"trackingStatus,omitempty"`
  TrackingURL string `json:"trackingUrl,omitempty"`
  UpdatedAt string `json:"updatedAt,omitempty"`
  // If true, auto calculates total parcel size using the size of items
  UseDimensionsOfItems bool `json:"useDimensionsOfItems,omitempty"`
  // If true, auto calculates total parcel weight using the weight of items
  UseWeightOfItems bool `json:"useWeightOfItems,omitempty"`
  // Weight of parcel
  Weight string `json:"weight,omitempty"`
  // Width of parcel
  Width string `json:"width,omitempty"`
}

// ShipmentResponse model
type ShipmentResponse struct {
  AdditionalMessage string `json:"additionalMessage,omitempty"`
  Code string `json:"code,omitempty"`
  Data *Shipment `json:"data,omitempty"`
  Message string `json:"message,omitempty"`
  Result bool `json:"result,omitempty"`
}

// Tracking model
type Tracking struct {
  CreatedAt string `json:"createdAt,omitempty"`
  Hash string `json:"hash,omitempty"`
  ID string `json:"id,omitempty"`
  LocationLat float64 `json:"locationLat,omitempty"`
  LocationLng float64 `json:"locationLng,omitempty"`
  LocationName string `json:"locationName,omitempty"`
  Owner string `json:"owner,omitempty"`
  StatusDate string `json:"statusDate,omitempty"`
  StatusDetails string `json:"statusDetails,omitempty"`
  Test bool `json:"test,omitempty"`
  TrackingStatusCode string `json:"trackingStatusCode,omitempty"`
  TrackingSubStatusCode string `json:"trackingSubStatusCode,omitempty"`
  UpdatedAt string `json:"updatedAt,omitempty"`
}

// WebhookUpdateTrackingRequest model
type WebhookUpdateTrackingRequest struct {
  Event string `json:"event"`
  Metadata string `json:"metadata"`
  Shipment Shipment `json:"data"`
}
