import { getCart, getCartItemsByIds } from "@/api/cart";

import useSWR from "swr";

export const useGetCart = () => {
  return useSWR(`/api/carts/`, () => getCart(), {
    revalidateOnFocus: false,
  });
};

export const useGetCartItemsByIds = (cartItemIds: number[]) => {
  return useSWR(
    `/api/cart_items/get_cart_items_by_ids`,
    () => getCartItemsByIds(cartItemIds),
    {
      revalidateOnFocus: false,
    },
  );
};
