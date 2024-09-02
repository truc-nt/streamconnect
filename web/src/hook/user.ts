import { getDefaultAddress } from "@/api/user";

import useSWR from "swr";

export const useGetDefaultAddress = (userId: number) => {
  return useSWR(
    `/api/users/${userId}/address`,
    () => getDefaultAddress(userId),
    {
      revalidateOnFocus: false,
    },
  );
};
