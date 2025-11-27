# Geliver Go SDK  
[![Go Reference](https://pkg.go.dev/badge/github.com/GeliverApp/geliver-go.svg)](https://pkg.go.dev/github.com/GeliverApp/geliver-go)

Geliver Go SDK — Geliver Kargo Pazaryeri (Shipping Marketplace) API için resmi Golang istemcisi.
Türkiye’nin e‑ticaret gönderim altyapısı için kolay kargo entegrasyonu sağlar.

• Dokümantasyon (TR/EN): https://docs.geliver.io

---

## İçindekiler

- Kurulum
- Hızlı Başlangıç
- Türkçe Akış (TR)
- Örnekler
- Coğrafi, Sağlayıcı, Şablonlar
- Notlar ve İpuçları (TR)

## Kurulum

Requires Go 1.21+.

```bash
go get github.com/GeliverApp/geliver-go@latest
```

İçe aktarma:

```go
import g "github.com/GeliverApp/geliver-go/pkg/geliver"
```

## Akış (TR)

1. Geliver Kargo API tokenı alın: https://app.geliver.io/apitokens
2. Gönderici adresi oluşturun (`CreateSenderAddress`).
3. Gönderi oluşturun; alıcıyı ID ile ya da adres nesnesi ile verin (`CreateShipmentTyped`).
4. Teklifleri bekleyip kabul edin (`AcceptOffer`).
5. Barkod, takip numarası ve etiket URL’leri Transaction içindeki Shipment’tan alın.
6. Test ortamında her `GET /shipments` isteği kargo durumunu bir adım ilerletir; prod’da webhookları kurun.
7. Etiketleri indirin (teklif kabulünden sonra Transaction içindeki URL'lerden indirin).
8. İade gerekiyorsa `CreateReturnShipment` fonksiyonunu kullanın.

## Hızlı Başlangıç

```go
package main

import (
    "context"
    "errors"
    "os"
    g "github.com/GeliverApp/geliver-go/pkg/geliver"
)

func example(ctx context.Context) error {
    c := g.NewClient("YOUR_TOKEN")

    // Gönderici adresi oluşturulur. Her gönderici adresi için tek seferlik yapılır. Oluşan gönderici adres ID'sini saklayıp tekrar kullanılır.
    sender, _ := c.CreateSenderAddress(ctx, g.CreateAddressRequest{
        Name: "ACME",
        Email: "ops@acme.test",
        Address1: "Street 1",
        CountryCode: "TR",
        CityName: "Istanbul",
        CityCode: "34",
        DistrictName: "Esenyurt",
        Zip: ptrs("34020"),
    })

    // Paket boyutları ve ağırlık (istek alanları string pointer olmalıdır)
    length, width, height, weight := "10.0", "10.0", "10.0", "1.0"
    distanceUnit := "cm"
    massUnit := "kg"

    // Alıcı adres bilgileri ile gönderi oluşturulur.
    s, _ := c.CreateShipmentWithRecipientAddress(ctx, g.CreateShipmentWithRecipientAddress{
        CreateShipmentRequestBase: g.CreateShipmentRequestBase{
            SenderAddressID: sender.ID,
            Length:         &length,
            Width:          &width,
            Height:         &height,
            DistanceUnit:   &distanceUnit,
            Weight:         &weight,
            MassUnit:       &massUnit,
            // Normal akış: Kapıda ödeme varsayılan olarak yoktur;
            // ProductPaymentOnDelivery alanını true yaparsanız kapıda ödeme olur.
            ProductPaymentOnDelivery: ptrb(false),
            // Test modu sadece deneme amaçlıdır; canlı ortamda bu alanı set etmeyin.
            Test: ptrb(true),
            Order:          &g.OrderRequest{ OrderNumber: "ABC12333322", SourceIdentifier: &[]string{"https://magazaadresiniz.com"}[0], TotalAmount: &[]string{"150"}[0], TotalAmountCurrency: &[]string{"TL"}[0] },
        },
        RecipientAddress: g.Address{
            Name: "John Doe",
            Address1: "Dest St 2",
            CountryCode: "TR",
            CityName: "Istanbul",
            CityCode: "34",
            DistrictName: "Esenyurt",
        },
    })

    // Teklifler create yanıtındaki offers alanında gelir
    offers := s.Offers
    if offers.Cheapest == nil {
        return errors.New("offers not ready yet; GET /shipments ile tekrar deneyin")
    }

    tx, err := c.AcceptOffer(ctx, offers.Cheapest.ID)
    if err != nil {
        return err
    }

    // Etiketler teklif kabulünden (Transaction) sonra üretilir; kabulün ardından URL'lerden indirebilirsiniz de. URL'lere her shipment nesnesinin içinden ulaşılır.
    _ = tx
    return nil
}
```

> Alternatif: Uygun olduğunda sunucuda kayıtlı `recipientAddressID` alanını kullanabilirsiniz.

---

## İade Gönderisi Oluşturun

```go
provider := "SURAT_STANDART"
retReq := g.ReturnShipmentRequest{
    WillAccept:         true,
    ProviderServiceCode: &provider,
    Count:              1,
}
ret, _ := c.CreateReturnShipment(ctx, s.ID, retReq)
_ = ret
```

Not:

- `ProviderServiceCode` alanı opsiyoneldir. Varsayılan olarak orijinal gönderinin sağlayıcısı kullanılır; dilerseniz bu alanı vererek değiştirebilirsiniz.
- `SenderAddress` alanı opsiyoneldir. Varsayılan olarak orijinal gönderinin alıcı adresi kullanılır; gerekirse bu alanı ayarlayabilirsiniz.

## Alıcı ID'si ile oluşturma (recipientAddressID)

```go
// Alıcı adresini sunucuda oluşturun ve IDyi alın
recipient, _ := c.CreateRecipientAddress(ctx, g.CreateAddressRequest{
    Name: "John Doe", Email: "john@example.com",
    Address1: "Dest St 2", CountryCode: "TR", CityName: "Istanbul", CityCode: "34",
    DistrictName: "Kadikoy",
})

// Ardından recipientAddressID ile gönderi oluşturun (typed istek)
length2, width2, height2, weight2 := "10.0", "10.0", "10.0", "1.0"
distanceUnit2 := "cm"
massUnit2 := "kg"
req := g.CreateShipmentWithRecipientID{
    CreateShipmentRequestBase: g.CreateShipmentRequestBase{
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

// Aşağıdaki akışlar "Ek Akışlar" bölümünde detaylandırılmıştır: listele, getir, güncelle, iptal, klonla.

## Örnekler

- Full flow: `examples/fullflow` (go/examples/fullflow/main.go)
- Operasyonel akışlar (Listele/Getir/Güncelle/Klonla/İptal): `examples/ops` (go/examples/ops/main.go)
- Tek aşamada gönderi (Create Transaction): `examples/onestep` (go/examples/onestep/main.go)
- Kapıda ödeme: `examples/pod` (go/examples/pod/main.go)
- Kendi anlaşmanızla etiket satın alma: `examples/ownagreement` (go/examples/ownagreement/main.go)
- Minimal webhook handler: `examples/webhook_server`
- Modeller: `pkg/geliver/models.go` (OpenAPI’den üretilmiştir)

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

## Kargo şablonları:

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

## Notlar ve İpuçları (TR)

- İstek tarafında `length`, `width`, `height`, `weight` alanları `*string` (ör. `"10.0"`) olarak gönderilmelidir; `json:",string"` kullanmayın.
- API bazı ondalık değerleri response tarafında string olarak döndürebilir; gerekirse modellerde uygun tiplere map edilir.
- Test ortamında her `GET /shipments` çağrısı durumu bir adım ilerletir; prod için webhookları kurun.
- Test modunu yalnızca denemeler için `Test: ptrb(true)` ile açın; canlı gönderilerde bu alanı set etmeyin.
- Takip numarası ile takip URL'si bazı kargo firmalarında teklif kabulünün hemen ardından oluşmayabilir. Paketi kargo şubesine teslim ettiğinizde veya kargo sizden teslim aldığında bu alanlar tamamlanır. Webhooklar ile değerleri otomatik çekebilir ya da teslimden sonra `shipment` GET isteği yaparak güncel bilgileri alabilirsiniz.

- Şehir/ilçe verileri için API’yi kullanın:

```go
_, _ = c.ListCities(ctx, "TR")
_, _ = c.ListDistricts(ctx, "TR", "34")
```

[![Geliver](https://geliver.io/geliverlogo.png)](https://geliver.io/)

Etiketler: go, golang, sdk, api-client, geliver, kargo, kargo-pazaryeri, shipping, e-commerce, turkey

---

## Ek Akışlar

- Tek aşamada gönderi oluşturma: https://docs.geliver.io/docs/shipments_and_transaction/create_shipment/

  İsteğe bağlı olarak `ProviderServiceCode` alanını vererek sadece o hizmet için teklif üretirsiniz; ardından gelen `offers.cheapest` ile etiketi satın alın.

```go
// Tek aşama (Create Transaction): gönderi bilgileriniz ile /transactions çağrılır, teklif seçimi/accept offer yapılmaz.
prov := "SURAT_STANDART"
reqTx := g.CreateShipmentWithRecipientAddress{
    CreateShipmentRequestBase: g.CreateShipmentRequestBase{
        SenderAddressID: sender.ID,
        Length: &length, Width: &width, Height: &height, DistanceUnit: &distanceUnit,
        Weight: &weight, MassUnit: &massUnit,
        ProviderServiceCode: &prov,
    },
    RecipientAddress: g.Address{ /* ... */ },
}
trx, _ := c.CreateTransactionWithRecipientAddress(ctx, reqTx)
_ = trx
```
Örnek: go/examples/onestep/main.go

- Kapıda ödeme (Payment on Delivery): https://docs.geliver.io/docs/shipments_and_transaction/create_shipment/create_shipment_payment_on_delivery

```go
// Kapıda ödeme: order alanında toplam ve para birimi zorunludur.
pod := true
total := "150"; currency := "TL"
order := g.OrderRequest{ OrderNumber: "ABC12333322", TotalAmount: &total, TotalAmountCurrency: &currency }
reqPod := g.CreateShipmentWithRecipientAddress{
    CreateShipmentRequestBase: g.CreateShipmentRequestBase{
        SenderAddressID: sender.ID,
        Length: &length, Width: &width, Height: &height, DistanceUnit: &distanceUnit,
        Weight: &weight, MassUnit: &massUnit,
        ProductPaymentOnDelivery: &pod,
        Order: &order,
    },
    RecipientAddress: g.Address{ /* ... */ },
}
_, _ = c.CreateTransactionWithRecipientAddress(ctx, reqPod)
```
Örnek: go/examples/pod/main.go

- Kendi anlaşmanızla etiket satın alma (provider ile): https://docs.geliver.io/docs/shipments_and_transaction/create_shipment/create_shipment_provider

```go
// Kendi anlaşmanızla satın alma: providerAccountID alanını da gönderin.
prov2 := "ARAS_EXPRESS"
accID := "YOUR_PROVIDER_ACCOUNT_ID"
reqProv := g.CreateShipmentWithRecipientAddress{
    CreateShipmentRequestBase: g.CreateShipmentRequestBase{
        SenderAddressID: sender.ID,
        Length: &length, Width: &width, Height: &height, DistanceUnit: &distanceUnit,
        Weight: &weight, MassUnit: &massUnit,
        ProviderServiceCode: &prov2,
        ProviderAccountID: &accID,
    },
    RecipientAddress: g.Address{ /* ... */ },
}
trx2, _ := c.CreateTransactionWithRecipientAddress(ctx, reqProv)
_ = trx2
```
Örnek: go/examples/ownagreement/main.go

### Gönderi Listeleme, Getir, Güncelle, İptal, Klonla

- Listeleme (docs): https://docs.geliver.io/docs/shipments_and_transaction/list_shipments
- Gönderi getir (docs): https://docs.geliver.io/docs/shipments_and_transaction/list_shipments
- Paket güncelle (docs): https://docs.geliver.io/docs/shipments_and_transaction/update_package_shipment
- Gönderi iptal (docs): https://docs.geliver.io/docs/shipments_and_transaction/cancel_shipment
- Gönderi klonla (docs): https://docs.geliver.io/docs/shipments_and_transaction/clone_shipment

```go
// Listeleme (sayfalandırma)
params := &g.ListParams{ Limit: &[]int{20}[0], Page: &[]int{1}[0] }
list, err := c.ListShipments(ctx, params)
if err != nil { return err }
for _, s := range list.Data {
    _ = s // s.ID, s.StatusCode, s.TrackingNumber, vs.
}

// Getir
sh, _ := c.GetShipment(ctx, "SHIPMENT_ID")
// Önemli alanlar: takip durumu ve kodları
// Detaylı kodlar: https://docs.geliver.io/docs/shipments_and_transaction/tracking_status_codes
if sh.TrackingStatus != nil {
    fmt.Println(sh.TrackingStatus.TrackingStatusCode, sh.TrackingStatus.TrackingSubStatusCode)
}

// Paket güncelle (eni, boyu, yüksekliği ve ağırlığı string olarak gönderin)
nLen, nWid, nHei, nWei := "12.0", "12.0", "10.0", "1.2"
upd := g.UpdatePackageRequest{ Length: &nLen, Width: &nWid, Height: &nHei, Weight: &nWei, DistanceUnit: &[]string{"cm"}[0], MassUnit: &[]string{"kg"}[0] }
sh2, _ := c.UpdatePackageTyped(ctx, sh.ID, upd)
_ = sh2

// İptal
_, _ = c.CancelShipment(ctx, sh.ID)

// Klonla
clone, _ := c.CloneShipment(ctx, sh.ID)
_ = clone
```
Örnek: go/examples/ops/main.go

---

## Hatalar ve İstisnalar

- İstemci iki durumda hata döndürür: (1) HTTP 4xx/5xx, (2) JSON envelope `result == false`.
- Hata türü `*geliver.APIError` olup sunucu alanlarını taşır: `Code`, `Message`, `AdditionalMessage`, `Status`, `Body`.

```go
s, err := c.CreateShipmentWithRecipientAddress(ctx, req)
if err != nil {
    if ae, ok := err.(*geliver.APIError); ok {
        fmt.Println("code:", ae.Code)
        fmt.Println("message:", ae.Message)
        fmt.Println("additional:", ae.AdditionalMessage)
        fmt.Println("status:", ae.Status)
    }
    return err
}
```
