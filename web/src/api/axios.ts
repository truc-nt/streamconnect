import axios from "axios";
export const sdkToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGlrZXkiOiJjN2MwOTgwMy05OWUzLTRmMGUtOTg3Ny0zYjU1MTdiNThkY2IiLCJwZXJtaXNzaW9ucyI6WyJhbGxvd19qb2luIl0sImlhdCI6MTcyMzMyOTgyOSwiZXhwIjoxNzI1OTIxODI5fQ.5x11sT5M7jzIM9EslqanSiMpnLeLTImr-zlzDKUuntc";

export default axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_GO,
  headers: { "Content-Type": "application/json" },
});

export const axiosJava = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_JAVA,
  headers: { "Content-Type": "application/json" }
});

export const axiosPrivate = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_URL_GO,
  headers: { "Content-Type": "application/json" },
  withCredentials: true,
});
