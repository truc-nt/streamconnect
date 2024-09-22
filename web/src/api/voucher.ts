import axios from "./axios";
import { IBaseVoucher } from "@/model/order";

interface IVoucherCreateRequest extends IBaseVoucher {}
export const createVoucher = async (request: IVoucherCreateRequest) => {
  return axios.post(`vouchers/shop/`, request);
};

export interface IShopVoucherGetResponse extends IBaseVoucher {
  is_saved: boolean;
}

export const getShopVouchers = async (shopId: number) => {
  const response = await axios.get<IShopVoucherGetResponse[]>(
    `/shops/${shopId}/vouchers`,
  );
  return response.data;
};

export const addUserVoucher = async (voucherId: number) => {
  return await axios.put(`vouchers/user/${voucherId}`);
};
