import axios from "axios";
import * as jwt from "jsonwebtoken";

interface TokenData {
  userId: number;
  username: string;
}

const axiosDefault = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_GO,
  headers: { "Content-Type": "application/json" },
});

axiosDefault.interceptors.request.use(function (config) {
  const token = localStorage.getItem('token');
  if (!!token) {
    const data: TokenData = jwt.decode(token);
    if (!!data) {
      config.headers["user_id"] = data.userId;
    }
  }
  config.headers.Authorization =  token ? `Bearer ${token}` : '';
  return config;
});

export default axiosDefault;


const axiosJava = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_JAVA,
  headers: { "Content-Type": "application/json" },
});

axiosJava.interceptors.request.use(function (config) {
  const token = localStorage.getItem('token');
  config.headers.Authorization =  token ? `Bearer ${token}` : '';
  return config;
});

export { axiosJava };

export const axiosPrivate = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_GO,
  headers: { "Content-Type": "application/json" },
  withCredentials: true,
});
