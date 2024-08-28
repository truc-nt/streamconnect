import { getExternalShops } from "@/api/external_shop";

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
