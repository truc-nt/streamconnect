import {axiosJava} from "@/api/axios";

export const login = async (username: string, password: string) => {
  return axiosJava.post(`auth/login`, {username, password});
};

export interface SignUpRequest {
    username: string;
    password: string;
    email: string;
    fullName: string;
}

export const signUp = async (data: SignUpRequest) => {
    return axiosJava.post(`auth/register`, data);
};
