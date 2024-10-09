import axios from "./axios";

import { IBaseUserAddress } from "@/model/order";
import { IBaseUser } from "@/model/user";

export const getUser = async () => {
  const response = await axios.get<IBaseUser>(`users/`);
  return response.data;
};

export interface IUpdateUserRequest {
  email: string;
  gender: string;
  birthdate: string;
}
export const updateUser = async (request: IUpdateUserRequest) => {
  const response = await axios.patch<IBaseUser>(`users/`, request);
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

interface ICreateAddressRequest {
  name: string;
  phone: string;
  address: string;
  city: string;
  is_default: boolean;
}
export const createAddress = async (request: ICreateAddressRequest) => {
  const response = await axios.post(`addresses/`, request);
  return response.data;
};
