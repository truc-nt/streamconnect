import { IBaseExternalVariant } from "@/model/product";

export interface IBaseOrder {
  id_order: number;
  id_shop: number;
  shop_name: string;
}

export interface IBaseExternalOrder {
  id_external_order: number;
  id_ecommerce: number;
  external_order_id_mapping: string;
  shipping_fee: number;
  internal_discount: number;
  external_discount: number;
}
export interface IBaseOrderItem extends IBaseExternalVariant {
  id_order_item: number;
  quantity: number;
  unit_price: number;
  paid_price: number;
}

export interface IBaseOrderAdditionInfo extends IBaseUserAddress {
  id_shipping_method: number;
  id_payment_method: number;
}

export interface IBaseUserAddress {
  id_user_address: number;
  name: string;
  phone: string;
  address: string;
  city: string;
  is_default: boolean;
}

export interface IBaseVoucher {
  id_voucher: number;
  code: string;
  discount: number;
  max_discount: number;
  type: string;
  target: string;
  method: string;
  min_purchase: number;
  quantity: number;
  start_time: string;
  end_time: string;
}
