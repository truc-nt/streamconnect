import axios from "./axios";
import { IBaseOrder, IBaseExternalOrder, IBaseOrderItem } from "@/model/order";

interface ICreateOrderRequest {
  id_address: number;
  external_orders: {
    cart_item_ids: number[];
    shipping_fee: number;
    shipping_fee_discount: number;
    external_discount: number;
    voucher_ids: number[];
  }[];
}

export const createOrder = async (createOrderRequest: ICreateOrderRequest) => {
  return axios.post(`/orders/create_with_cart_items`, createOrderRequest);
};

export interface IBuyOrdersGetRequest extends IBaseOrder {
  external_orders: IExternalOrder[];
}
interface IExternalOrder extends IBaseExternalOrder {
  order_items: IBaseOrderItem[];
}

export const getBuyOrders = async () => {
  const response = await axios.get<IBuyOrdersGetRequest[]>(`orders/buy`);
  return response.data;
};

export interface IOrderGetRequest extends IExternalOrder {}
export const getOrder = async (orderId: number) => {
  const response = await axios.get<IOrderGetRequest[]>(`orders/${orderId}`);
  return response.data;
};
