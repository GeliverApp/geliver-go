
# Geliver Go SDK

Geliver Go SDK — Geliver Kargo Pazaryeri (Shipping Marketplace) API için resmi Golang istemcisi.
Türkiye’nin e‑ticaret gönderim altyapısı için kolay kargo entegrasyonu sağlar.

• Dokümantasyon (TR/EN): https://docs.geliver.com

---

## İçindekiler

- Kurulum
- Hızlı Başlangıç
- Türkçe Akış (TR)
- Örnekler
- Coğrafi, Sağlayıcı, Şablonlar
- Notlar (TR)

---

## Kurulum

Requires Go 1.21+.

```bash
go get github.com/GeliverApp/geliver-go@latest
```

İçe aktarma:

```go
import g "github.com/GeliverApp/geliver-go/pkg/geliver"
```

---

## Hızlı Başlangıç

```go
package main

import (
    "context"
    "os"
    g "github.com/GeliverApp/geliver-go/pkg/geliver"
)

func example(ctx context.Context) error {
    c := g.NewClient("YOUR_TOKEN")

    sender, _ := c.CreateSenderAddress(ctx, g.CreateAddressRequest{
        Name: "ACME",
        Email: "ops@acme.test",
        Address1: "Street 1",
        CountryCode: "TR",
        CityName: "Istanbul",
        CityCode: "34",
        DistrictName: "Esenyurt",
        DistrictID: 107605,
        Zip: "34020",
    })

    // Paket boyutları ve ağırlık
    length, width, height, weight := 10.0, 10.0, 10.0, 1.0
    distanceUnit := "cm"
    massUnit := "kg"

    s, _ := c.CreateShipmentWithRecipientAddress(ctx, g.CreateShipmentWithRecipientAddress{
        CreateShipmentRequestBase: g.CreateShipmentRequestBase{
            SourceCode:     "API",
            SenderAddressID: sender.ID,
            Length:         &length,
            Width:          &width,
            Height:         &height,
            DistanceUnit:   &distanceUnit,
            Weight:         &weight,
            MassUnit:       &massUnit,
        },
        RecipientAddress: g.Address{
            Name: "John Doe",
            Address1: "Dest St 2",
            CountryCode: "TR",
            CityName: "Istanbul",
            CityCode: "34",
            DistrictName: "Esenyurt",
            DistrictID: 107605,
            Zip: "34020",
        },
    })

    // Etiketler bazı akışlarda create sonrasında hazır olabilir; varsa hemen indirin
    if s.LabelURL != "" {
        b, _ := c.DownloadShipmentLabel(ctx, s.ID)
        _ = os.WriteFile("label_pre.pdf", b, 0644)
    }
    if s.ResponsiveLabelURL != "" {
        html, _ := c.DownloadResponsiveURL(ctx, s.ResponsiveLabelURL)
        _ = os.WriteFile("label_pre.html", []byte(html), 0644)
    }

    _ = s
    return nil
}
```

> Alternatif: Uygun olduğunda sunucuda kayıtlı `recipientAddressID` alanını kullanabilirsiniz.

## Alıcıyı ID ile oluşturma (recipientAddressID)

```go
// Alıcı adresini sunucuda oluşturun ve ID alın
recipient, _ := c.CreateRecipientAddress(ctx, g.CreateAddressRequest{
    Name: "John Doe", Email: "john@example.com",
    Address1: "Dest St 2", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
    DistrictName: "Kadikoy", DistrictID: 100000, Zip: "34000",
})

// Ardından recipientAddressID ile gönderi oluşturun (typed istek)
length2, width2, height2, weight2 := 10.0, 10.0, 10.0, 1.0
distanceUnit2 := "cm"
massUnit2 := "kg"
req := g.CreateShipmentWithRecipientID{
    CreateShipmentRequestBase: g.CreateShipmentRequestBase{
        SourceCode: "API",
        SenderAddressID: sender.ID,
        Length:       &length2,
        Width:        &width2,
        Height:       &height2,
        DistanceUnit: &distanceUnit2,
        Weight:       &weight2,
        MassUnit:     &massUnit2,
        Test: ptrb(true),
    },
    RecipientAddressID: recipient.ID,
}
s2, _ := c.CreateShipmentWithRecipientID(ctx, req)
_ = s2
```

---

## Türkçe Akış (TR)

1. Geliver Kargo API tokenı alın: https://app.geliver.io/apitokens
2. Gönderici adresi oluşturun (`CreateSenderAddress`).
3. Gönderi oluşturun; alıcıyı ID ile ya da adres nesnesi ile verin (`CreateShipmentTyped`).
4. Teklifleri bekleyip kabul edin (`AcceptOffer`).
5. Barkod, takip numarası ve etiket URL’leri Transaction içindeki Shipment’tan alın.
6. Test ortamında her `GET /shipments` isteği kargo durumunu bir adım ilerletir; prod’da webhookları kurun.
7. Etiketleri indirin (`DownloadShipmentLabel`).
8. İade gerekiyorsa `CreateReturnShipment` fonksiyonunu kullanın.

---

## Örnekler

- Full flow: `examples/fullflow`
- Minimal webhook handler: `examples/webhook_server`
- Modeller: `pkg/geliver/models.go` (OpenAPI’den üretilmiştir)

---

## Coğrafi, Sağlayıcı, Şablonlar

Geo (şehir/ilçe):

```go
cities, _ := c.ListCities(ctx, "TR")
districts, _ := c.ListDistricts(ctx, "TR", "34")
```

Provider hesapları:

```go
acc, _ := c.CreateProviderAccount(ctx, g.ProviderAccountRequest{
    Username: "user",
    Password: "pass",
    Name:     "My Account",
    ProviderCode: "SURAT",
    Version: 1,
    IsActive: true,
    IsPublic: false,
    Sharable: false,
    IsDynamicPrice: false,
})
list, _ := c.ListProviderAccounts(ctx)
_, _ = c.DeleteProviderAccount(ctx, acc.ID, ptrb(true))
```

Kargo şablonları:

```go
tpl, _ := c.CreateParcelTemplate(ctx, g.CreateParcelTemplateRequest{
    Name:         "Small Box",
    DistanceUnit: "cm",
    MassUnit:     "kg",
    Height:       "4",
    Length:       "4",
    Weight:       "1",
    Width:        "4",
})
tpls, _ := c.ListParcelTemplates(ctx)
_, _ = c.DeleteParcelTemplate(ctx, tpl.ID)
```

---

## Notlar (TR)

- API bazı ondalık değerleri string olarak döndürebilir; Go tarafında `json:",string"` ile tip dönüşümü yapılır.
- Teklif beklerken ~1 sn aralık idealdir.
- Test ortamında her `GET /shipments` çağrısı durumu bir adım ilerletir; prod için webhookları kurun.
- İlçe seçimi için `districtID` tercih edilir; `districtName` her zaman güvenilir olmayabilir.
- Şehir/ilçe verileri için API’yi kullanın:

```go
_, _ = c.ListCities(ctx, "TR")
_, _ = c.ListDistricts(ctx, "TR", "34")
```

[![Geliver](https://geliver.io/geliverlogo.png)](https://geliver.io/)

Etiketler: go, golang, sdk, api-client, geliver, kargo, kargo-pazaryeri, shipping, e-commerce, turkey
