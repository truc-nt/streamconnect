import {
  getExternalProducts,
  getExternalVariants,
} from "@/api/external_product";

import useSWR from "swr";

export const useGetExternalProducts = () => {
  return useSWR(
    `/api/external_variants`,
    async () => await getExternalProducts(),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};

export const useGetExternalVariants = (externalProductIdMapping: string) => {
  return useSWR(
    [
      `/api/external_variants/${externalProductIdMapping}`,
      externalProductIdMapping,
    ],
    async ([_, externalProductIdMapping]) =>
      await getExternalVariants(externalProductIdMapping),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
