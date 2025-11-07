package geliver

type CreateAddressRequest struct {
	Name               string  `json:"name"`
	Email              string  `json:"email"`
	Phone              *string `json:"phone,omitempty"`
	Address1           string  `json:"address1"`
	Address2           *string `json:"address2,omitempty"`
	CountryCode        string  `json:"countryCode"`
	CityName           string  `json:"cityName"`
	CityCode           string  `json:"cityCode"`
	DistrictName       string  `json:"districtName"`
	DistrictID         int     `json:"districtID"`
	Zip                string  `json:"zip"`
	ShortName          *string `json:"shortName,omitempty"`
	IsRecipientAddress *bool   `json:"isRecipientAddress,omitempty"`
}

type CreateShipmentRequestBase struct {
    SourceCode          string  `json:"sourceCode"`
    SenderAddressID     string  `json:"senderAddressID"`
    Length              *string `json:"length,omitempty"`
    Width               *string `json:"width,omitempty"`
    Height              *string `json:"height,omitempty"`
    DistanceUnit        *string `json:"distanceUnit,omitempty"`
    Weight              *string `json:"weight,omitempty"`
    MassUnit            *string `json:"massUnit,omitempty"`
    ProviderServiceCode *string `json:"providerServiceCode,omitempty"`
    Test                *bool   `json:"test,omitempty"`
}

type CreateShipmentWithRecipientID struct {
	CreateShipmentRequestBase
	RecipientAddressID string `json:"recipientAddressID"`
}

type CreateShipmentWithRecipientAddress struct {
	CreateShipmentRequestBase
	RecipientAddress Address `json:"recipientAddress"`
}

type UpdatePackageRequest struct {
    Height       *string `json:"height,omitempty"`
    Width        *string `json:"width,omitempty"`
    Length       *string `json:"length,omitempty"`
    DistanceUnit *string `json:"distanceUnit,omitempty"`
    Weight       *string `json:"weight,omitempty"`
    MassUnit     *string `json:"massUnit,omitempty"`
}

// ReturnShipmentRequest represents a return creation body.
// Note: SenderAddress is optional; backend defaults it to the original shipment's recipient address
// unless overridden here.
type ReturnShipmentRequest struct {
	IsReturn            bool     `json:"isReturn"`
	WillAccept          bool     `json:"willAccept"`
	ProviderServiceCode *string  `json:"providerServiceCode,omitempty"`
	Count               int      `json:"count"`
	SenderAddress       *Address `json:"senderAddress,omitempty"`
}

type CreateParcelTemplateRequest struct {
	Name         string `json:"name"`
	DistanceUnit string `json:"distanceUnit"`
	MassUnit     string `json:"massUnit"`
	Height       string `json:"height"`
	Length       string `json:"length"`
	Weight       string `json:"weight"`
	Width        string `json:"width"`
}
