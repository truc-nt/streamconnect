import axios from "./axios";

export interface ICart {
  id_shop: number;
  shop_name: string;
  cart_livestream_external_variant: {
    name: string;
    option: { [key: string]: string };
    id_livestream_external_variant: number;
    id_ecommerce: number;
    price: number;
    quantity: number;
  }[];
}

export const getCart = async (cartId: number) => {
  return axios.get<ICart[]>(`cart/${cartId}`);
};
