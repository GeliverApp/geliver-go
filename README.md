Geliver SDK (Go)

Geliver Go SDK — official Golang client for Geliver Kargo Pazaryeri (Shipping Marketplace) API.
Türkiye’nin e‑ticaret gönderim altyapısı için kolay kargo entegrasyonu sağlar.

Dokümantasyon (TR): Modeller ve endpoint detayları için https://docs.geliver.com
Documentation (EN): For detailed models and endpoints, see https://docs.geliver.com

Türkçe (TR)

- 0. Geliver Kargo API tokenı alın (https://app.geliver.io/apitokens adresinden)
- 1. Gönderici adresi oluşturun (CreateSenderAddress)
- 2. Gönderi oluşturun; alıcıyı ID ile ya da adres nesnesi ile verin (CreateShipmentTyped)
- 3. Teklifleri bekleyip kabul edin (AcceptOffer)
- 4. Barkod, takip numarası, etiket URL’leri Transaction içindeki Shipment’ten alın
- 5. Test gönderilerinde her GET /shipments isteği kargo durumunu bir adım ilerletir; prod'da webhookları kurun
- 6. Etiketleri indirin (DownloadShipmentLabel)
- 7. İade gönderisi gerekiyorsa CreateReturnShipment fonksiyonunu kullanın

Kurulum

- Use as a module: `module github.com/geliver-io/yourapp` and require `github.com/geliver-io/sdk-go` (local path until published)

Akış (Seçenek A: alıcı adresini inline gönderme)

```go
import (
  "context"
  g "github.com/geliver-io/sdk-go/pkg/geliver"
)

func example(ctx context.Context) error {
  c := g.NewClient("YOUR_TOKEN")
  sender, _ := c.CreateSenderAddress(ctx, g.CreateAddressRequest{ Name: "ACME", Email: "ops@acme.test", Address1: "Street 1", CountryCode: "TR", CityName: "Istanbul", CityCode: "34", DistrictName: "Esenyurt", DistrictID: 107605, Zip: "34020" })
  s, _ := c.CreateShipmentWithRecipientAddress(ctx, g.CreateShipmentWithRecipientAddress{ CreateShipmentRequestBase: g.CreateShipmentRequestBase{ SourceCode: "API", SenderAddressID: sender.ID }, RecipientAddress: g.Address{ Name: "John Doe", Address1: "Dest St 2", CountryCode: "TR", CityName: "Istanbul", CityCode: "34", DistrictName: "Esenyurt", DistrictID: 107605, Zip: "34020" } })
  _ = s
  return nil
}
```

Seçenek B: recipientAddressID kullanımı (sunucuda kayıtlı adres) ana README'de gösterilmiştir.

Diğer Örnekler (Go)

- Coğrafi Servisler (Geo)

```go
cities, _ := c.ListCities(ctx, "TR")
districts, _ := c.ListDistricts(ctx, "TR", "34")
```

- Sağlayıcı Hesapları

```go
acc, _ := c.CreateProviderAccount(ctx, geliver.ProviderAccountRequest{
    Username: "user", Password: "pass", Name: "My Account", ProviderCode: "SURAT",
    Version: 1, IsActive: true, IsPublic: false, Sharable: false, IsDynamicPrice: false,
})
list, _ := c.ListProviderAccounts(ctx)
_, _ = c.DeleteProviderAccount(ctx, acc.ID, ptrb(true))
```

- Kargo Şablonları

```go
tpl, _ := c.CreateParcelTemplate(ctx, geliver.CreateParcelTemplateRequest{
    Name: "Small Box", DistanceUnit: "cm", MassUnit: "kg", Height: "4", Length: "4", Weight: "1", Width: "4",
})
tpls, _ := c.ListParcelTemplates(ctx)
_, _ = c.DeleteParcelTemplate(ctx, tpl.ID)
```

Örnekler

- See `sdks/go/examples/fullflow`.
- Generated structs are in `pkg/geliver/models.go` (auto-generated from OpenAPI).

Manuel takip kontrolü (isteğe bağlı)

```go
s, _ := c.GetShipment(ctx, "shipment-id")
// trackingStatus is part of the JSON payload; access via generic map if needed
// or use your own typed wrapper based on models.go
```

Modeller

- Shipment, Transaction, TrackingStatus, Address, ParcelTemplate, ProviderAccount, Webhook, Offer, PriceQuote and more.
- Full list in `pkg/geliver/models.go`.

Enum Kullanımı (TR)

```go
// Enum türleri (ör. ShipmentDistanceUnit, ShipmentLabelFileType) string tabanlıdır
gs, _ := c.GetShipment(ctx, "shipment-id")
if gs.LabelFileType == geliver.ShipmentLabelFileTypePDF {
    fmt.Println("PDF etiket hazır")
}
```

Notlar ve İpuçları (TR)

- API, ondalıklı değerleri string olarak gönderebilir; Go tarafında `json:",string"` ile float/int'e dönüştürülür.
- Teklif beklerken 1 sn aralık idealdir; gereksiz sorgudan kaçının.
- Test gönderilerinde her GET /shipments isteği kargo durumunu bir adım ilerletir; prod'da webhookları kurun.
- İlçe seçimi: districtID (number) önerilir. districtName her zaman güvenilir olmayabilir.
- Şehir/İlçe seçimi: cityCode ve cityName birlikte veya ayrı verilebilir; eşleşme için cityCode daha güvenlidir. Şehir/ilçe verilerini API ile alabilirsiniz:

```go
_ , _ = c.ListCities(ctx, "TR")
_ , _ = c.ListDistricts(ctx, "TR", "34")
```

[![Geliver Kargo Pazaryeri](https://geliver.io/geliverlogo.png)](https://geliver.io/)
Geliver Kargo Pazaryeri: https://geliver.io/

Etiketler (Tags): go, golang, sdk, api-client, geliver, kargo, kargo-pazaryeri, shipping, e-commerce, turkey
