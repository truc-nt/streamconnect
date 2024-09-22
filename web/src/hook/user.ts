import { getUserInfo, getDefaultAddress, getAddresses } from "@/api/user";

import useSWR from "swr";

export const useGetUserInfo = () => {
  return useSWR(`/api/users`, () => getUserInfo(), {
    revalidateOnFocus: false,
  });
};

export const useGetDefaultAddress = () => {
  return useSWR(`/addresses/default_address`, () => getDefaultAddress(), {
    revalidateOnFocus: false,
  });
};

export const useGetAddresses = () => {
  return useSWR(`/addresses`, () => getAddresses(), {
    revalidateOnFocus: false,
  });
};
