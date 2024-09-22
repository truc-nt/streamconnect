import { getBuyOrders, getOrder } from "@/api/order";

import useSWR from "swr";

export const useGetBuyOrders = () => {
  return useSWR(`/api/orders/buy`, () => getBuyOrders(), {
    revalidateOnFocus: false,
  });
};

export const useGetOrder = (orderId: number) => {
  return useSWR(`/api/orders/${orderId}`, () => getOrder(orderId), {
    revalidateOnFocus: false,
  });
};
