# Changelog

Bu dosya SDK'daki önemli değişiklikleri listeler.

This file documents notable changes in the SDK.

## Sürüm / Version

- Türkçe: Bu değişiklikler `1.2.0` sürümü için hazırlandı.
- English: These changes are prepared for version `1.2.0`.

## Türkçe

### 1.2.0

- Tüm yeni gönderilerde gönderici ve alıcı için adres ID veya satır içi adresin tam olarak biri doğrulanır.
- `CreateShipmentTyped(...)` ile `IsReturn` ve gönderici/alıcı için bağımsız ID veya satır içi adres seçimi kullanılarak iade gönderisi oluşturulabilir.
- `examples/incomingshipment` eklendi.

### 1.1.0

#### Eklendi

- `CreateReturnTransaction(...)` ile iadeyi oluşturup etiketi hemen satın alma akışı eklendi.
- İki yeni iade örneği eklendi:
  - `examples/returnshipment`
  - `examples/returntransaction`

#### Değişti

- `CreateReturnShipment(...)` artık shipment-only iade akışıdır ve etiketi satın almaz.
- İade dokümanı iki akışı ayrı anlatır.
- README örnekleri, etiketin daha sonra `AcceptOffer(...)` ile satın alınabileceğini açıklar.

## English

### 1.2.0

- All new shipment flows validate that each side provides exactly one address ID or inline address.
- `CreateShipmentTyped(...)` supports independent returns with `IsReturn` and either an address ID or an inline address for each side.
- Added `examples/incomingshipment`.

### 1.1.0

#### Added

- Added `CreateReturnTransaction(...)` for creating a return shipment and purchasing the label immediately.
- Added return examples for:
  - `examples/returnshipment`
  - `examples/returntransaction`

#### Changed

- `CreateReturnShipment(...)` now represents the shipment-only return flow and does not purchase the label.
- Return documentation now explains the two return flows separately.
- README examples now document that label purchase can be performed later with `AcceptOffer(...)`.
