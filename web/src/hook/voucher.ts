import { getShopVouchers } from "@/api/voucher";

import useSWR from "swr";

export const useGetShopVouchers = (shopId: number) => {
  return useSWR(`/api/shops/vouchers`, () => getShopVouchers(shopId), {
    revalidateOnFocus: false,
  });
};
