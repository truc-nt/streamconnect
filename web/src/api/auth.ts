import { axiosJava } from "@/api/axios";

export const login = async (username: string, password: string) => {
  const response = await axiosJava.post(`auth/login`, { username, password });
  return response.data;
};

export interface RegisterRequest {
  username: string;
  password: string;
  email: string;
}

export const register = async (data: RegisterRequest) => {
  return axiosJava.post<RegisterRequest>(`auth/register`, data);
};
