import axios from "./axios";
import { IBaseCartItem } from "@/model/cart";

export interface ICart {
  id_shop: number;
  shop_name: string;
  cart_items: IBaseCartItem[];
}

export const getCart = async () => {
  const response = await axios.get<ICart[]>(`carts/`);
  return response.data;
};

export interface ICartItemLivestreamExternalVariant {
  id_livestream_external_variant: number;
  quantity: number;
}
export const addToCart = async (item: ICartItemLivestreamExternalVariant[]) => {
  return axios.post(`carts/`, item);
};

export const updateQuantity = async (cartItemId: number, quantity: number) => {
  return axios.patch(`cart_items/${cartItemId}`, {
    quantity,
  });
};

export interface ICartItemsGroupByEcommerce {
  id_ecommerce: number;
  cart_items_group_by_shop: ICart[];
}

export const getCartItemsByIds = async (cartItemIds: number[]) => {
  const response = await axios.post<ICartItemsGroupByEcommerce[]>(
    `cart_items/get_cart_items_by_ids`,
    cartItemIds.map((id) => ({ id_cart_item: id })),
  );
  return response.data;
};
