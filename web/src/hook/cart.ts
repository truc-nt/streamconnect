import { getCart, getCartItemsByIds } from "@/api/cart";

import useSWR from "swr";

export const useGetCart = (cartId: number) => {
  return useSWR(`/api/carts/${cartId}`, () => getCart(cartId), {
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
