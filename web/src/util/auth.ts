import * as jwt from "jsonwebtoken";

export const decodeJwt = (token: string) => {
  const decodeToken = jwt.decode(token);
  return decodeToken?.sub ? JSON.parse(decodeToken?.sub) : null;
};
