import { getLivestreamProducts } from "@/api/livestream";
import useSWR from "swr";

export const useGetLivestreamProducts = (livestreamId: number) => {
  return useSWR(
    [`/api/livestreams/${livestreamId}/products`, livestreamId],
    async ([_, livestreamId]) => await getLivestreamProducts(livestreamId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
