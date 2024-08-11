"use client";

import CategorySlider from "@/components/CategorySlider";
import LivestreamPreview from "@/components/LivestreamPreview";
import LivestreamsList from "@/components/LivestreamsList";
import { Typography } from "@mui/material";
import Alert from "@/components/core/Alert";
import { useAppSelector } from "@/store/store";

const AlertProvider = ({ children }: { children: React.ReactNode }) => {
  const { open, message, type } = useAppSelector((state: any) => state.alert);
  return (
    <>
      {open && <Alert message={message} type={type} />}
      {children}
    </>
  );
};

export default AlertProvider;
