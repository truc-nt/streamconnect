import { getLivestreamProduct } from "@/api/livestream_product";
import useSWR from "swr";

export const useGetLivestreamProduct = (livestreamProductId: number | null) => {
  if (!livestreamProductId) return { data: null, error: null };
  return useSWR(
    [`/api/livestreams/${livestreamProductId}`, livestreamProductId],
    async ([_, livestreamProductId]) =>
      await getLivestreamProduct(livestreamProductId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
