"use client";
import LivestreamCreate from "./component/LivestreamCreate";
import { useAppSelector } from "@/store/store";

const Page = () => {
  const { userId } = useAppSelector((state) => state.authReducer);
  if (!userId) return null;
  return <LivestreamCreate shopId={userId} />;
};

export default Page;
