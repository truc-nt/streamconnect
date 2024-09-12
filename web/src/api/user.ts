import axios from "./axios";

import { IBaseUserAddress } from "@/model/order";

export interface User {
  username: string;
  email: string;
}

export const getUserInfo = async () => {
  const response = await axios.get<User>(`users/`);
  return response.data;
};

export const getDefaultAddress = async () => {
  const response = await axios.get<IBaseUserAddress>(
    `addresses/default_address`,
  );
  return response.data;
};

export const getAddresses = async () => {
  const response = await axios.get<IBaseUserAddress[]>(`addresses/`);
  return response.data;
};
