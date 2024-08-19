import { getExternalShops, getExternalProducts } from "@/api/external_shop";

import useSWR from "swr";

export const useGetExternalShops = (shopId: number) => {
  return useSWR(
    [`/api/shops/${shopId}/external_shops`, shopId],
    async ([_, shopId]) => await getExternalShops(shopId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};

export const useGetExternalProducts = (externalShopId: number) => {
  return useSWR(
    [`/api/external_shops/${externalShopId}/external_products`, externalShopId],
    async ([_, externalShopId]) => await getExternalProducts(externalShopId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
