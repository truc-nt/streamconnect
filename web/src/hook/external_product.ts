import {
  getExternalProducts,
  getExternalVariants,
  getVariantsByExternalProductIdMapping,
} from "@/api/external_product";

import useSWR from "swr";

export const useGetExternalProducts = (shopId: number) => {
  return useSWR(
    `/api/${shopId}/external_products/`,
    async () => await getExternalProducts(shopId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: true,
    },
  );
};

export const useGetExternalVariants = (externalProductIdMapping: string) => {
  return useSWR(
    `/api/external_products/${externalProductIdMapping}`,
    () => getExternalVariants(externalProductIdMapping),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};

export const useGetVariantsByExternalProductIdMapping = (
  externalProductIdMapping: string,
) => {
  return useSWR(
    `/api/external_products/${externalProductIdMapping}/variants`,
    () => getVariantsByExternalProductIdMapping(externalProductIdMapping),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
