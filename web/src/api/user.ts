import axios from "./axios";

export interface IUserAddress {
  id_user_address: number;
  name: string;
  phone: string;
  address: string;
  city: string;
  is_default: boolean;
}

export const getDefaultAddress = async (userId: number) => {
  const response = await axios.get<IUserAddress>(`users/${userId}/address`);
  return response.data;
};
