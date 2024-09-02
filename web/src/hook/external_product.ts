import {
  getExternalProducts,
  getExternalVariants,
  getVariantsByExternalProductIdMapping,
} from "@/api/external_product";

import useSWR from "swr";

export const useGetExternalProducts = () => {
  return useSWR(
    `/api/external_products/`,
    async () => await getExternalProducts(),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
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
