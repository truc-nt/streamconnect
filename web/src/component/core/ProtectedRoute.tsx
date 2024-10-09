"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAppSelector } from "@/store/store";

const ProtectedRoute = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const { userId } = useAppSelector((state) => state.authReducer);
  const token = localStorage.getItem("token");

  useEffect(() => {
    if (!token) {
      router.push("/");
    }
  }, [userId]);

  return <>{children}</>;
};

export default ProtectedRoute;
