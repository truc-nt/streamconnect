import { getCart } from "@/api/cart";

import useSWR from "swr";

export const useGetCart = (cartId: number) => {
  return useSWR(
    `/api/cart/${cartId}`,
    () => getCart(cartId),
    {
      revalidateOnFocus: false,
    },
  );
};
