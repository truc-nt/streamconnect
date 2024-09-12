import { IBaseExternalVariant } from "@/model/product";

export interface IBaseCart {
  id_shop: number;
  shop_name: string;
}

export interface IBaseCartItem extends IBaseExternalVariant {
  id_cart_item: number;
  quantity: number;
  max_quantity: number;
}
