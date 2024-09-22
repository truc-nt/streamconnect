import axios from "axios";
import { decodeJwt } from "@/util/auth";

const axiosDefault = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_GO,
  headers: { "Content-Type": "application/json" },
  withCredentials: true,
});

axiosDefault.interceptors.request.use(function (config) {
  const token = localStorage.getItem("token");
  if (token) {
    const userInfo = decodeJwt(token);
    if (userInfo) {
      config.headers["user_id"] = userInfo.userId;
    }
  }
  config.headers.Authorization = token ? `Bearer ${token}` : "";
  return config;
});

export default axiosDefault;

const axiosJava = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_JAVA,
  headers: { "Content-Type": "application/json" },
});

axiosJava.interceptors.request.use(function (config) {
  const token = localStorage.getItem("token");
  config.headers.Authorization = token ? `Bearer ${token}` : "";
  return config;
});

export { axiosJava };

/*export const axiosPrivate = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_GO,
  headers: { "Content-Type": "application/json" },
  withCredentials: true,
});*/
