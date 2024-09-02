import { getProductsByShopId } from "@/api/product";

import useSWR from "swr";

export const useGetProductsByShopId = (shopId: number) => {
  return useSWR(
    `/api/shops/${shopId}/products`,
    () => getProductsByShopId(shopId),
    {
      revalidateOnFocus: false,
    },
  );
};
