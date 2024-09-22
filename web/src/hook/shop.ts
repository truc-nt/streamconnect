import { getShop, getProducts } from "@/api/shop";

import useSWR from "swr";

export const useGetShop = (shopId: number) => {
  return useSWR(`/api/shops/${shopId}`, () => getShop(shopId), {
    //shouldRetryOnError: false,
    revalidateOnFocus: false,
  });
};

export const useGetProducts = (shopId: number) => {
  return useSWR(
    [`/api/shops/${shopId}/products`, shopId],
    async ([_, shopId]) => await getProducts(shopId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
