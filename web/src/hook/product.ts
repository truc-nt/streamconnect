import { getProductsByShopId } from "@/api/product";

import useSWR from "swr";

export const useGetProductsByShopId = (shopId: number) => {
  return useSWR(
    [`/api/shops/${shopId}/products`, shopId],
    async ([_, shopId]) => await getProductsByShopId(shopId),
    {
      revalidateOnFocus: false,
    },
  );
};
