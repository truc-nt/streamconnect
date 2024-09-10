import { axiosJava } from "@/api/axios";

export const login = async (username: string, password: string) => {
  return axiosJava.post(`auth/login`, { username, password });
};

export interface RegisterRequest {
  username: string;
  password: string;
  email: string;
}

export const register = async (data: RegisterRequest) => {
  return axiosJava.post<RegisterRequest>(`auth/register`, data);
};
