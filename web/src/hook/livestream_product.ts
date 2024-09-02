import { getLivestreamProduct } from "@/api/livestream_product";
import useSWR from "swr";

export const useGetLivestreamProduct = (livestreamProductId: number) => {
  return useSWR(
    `/api/livestream_products/${livestreamProductId}`,
    () => getLivestreamProduct(livestreamProductId),
    {
      revalidateOnFocus: false,
    },
  );
};
