import axios from "./axios";

export interface ICart {
  id_shop: number;
  shop_name: string;
  cart_livestream_external_variant: ICartItem[];
}

export interface ICartItem {
  name: string;
  option: Record<string, string>;
  id_livestream_external_variant: number;
  id_ecommerce: number;
  price: number;
  quantity: number;
}

export const getCart = async (cartId: number) => {
  const response = await axios.get<ICart[]>(`cart/${cartId}`);
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
  return axios.post(`cart/${cartId}`, item);
};
