import axios from "./axios";

export interface ICart {
  id_shop: number;
  shop_name: string;
  cart_items: ICartItem[];
}

export interface ICartItem {
  id_cart_item: number;
  name: string;
  option: { [key: string]: string };
  id_livestream_external_variant: number;
  id_ecommerce: number;
  price: number;
  quantity: number;
  max_quantity: number;
  image_url: string;
}

export const getCart = async (cartId: number) => {
  const response = await axios.get<ICart[]>(`carts/${cartId}`);
  return response.data;
};

export interface ICartItemLivestreamExternalVariant {
  id_livestream_external_variant: number;
  quantity: number;
}
export const addToCart = async (
  cartId: number,
  item: ICartItemLivestreamExternalVariant[],
) => {
  return axios.post(`carts/${cartId}`, item);
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
  console.log(cartItemIds);
  const response = await axios.post<ICartItemsGroupByEcommerce[]>(
    `cart_items/get_cart_items_by_ids`,
    cartItemIds.map((id) => ({ id_cart_item: id })),
  );
  return response.data;
};
