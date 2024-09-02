import axios from "./axios";

interface ICreateOrderRequest {
  id_user: number;
  id_cart_items: number[];
  address: string;
}

export const createOrder = async (createOrderRequest: ICreateOrderRequest) => {
  return axios.post(`/orders/create_with_cart_items`, createOrderRequest);
};
