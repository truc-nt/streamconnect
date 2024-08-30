import {axiosJava} from "@/api/axios";

export const login = async (username: string, password: string) => {
  return axiosJava.post(`auth/login`, {username, password});
};