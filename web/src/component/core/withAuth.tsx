// isAuth.tsx

"use client";
import { useLayoutEffect } from "react";
import { useRouter } from "next/navigation";
import { useAppSelector } from "@/store/store";

const isAuth = (Component: React.ComponentType<any>) => {
  return function IsAuth(props: any) {
    const router = useRouter();
    const { userId } = useAppSelector((state) => state.authReducer);

    useLayoutEffect(() => {
      const token = localStorage.getItem("token");
      if (!token) {
        router.push("/");
      }
    }, [userId]);

    if (!userId) {
      return null;
    }

    return <Component {...props} />;
  };
};

export default isAuth;
