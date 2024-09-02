"use client";
import React from "react";
import { Spin } from "antd";
import { useAppSelector } from "@/store/store";

const LoadingProvider = ({ children }: { children: React.ReactNode }) => {
  const { open } = useAppSelector((state) => state.loadingReducer);
  return (
    <Spin tip="Loading..." size="large" spinning={open}>
      {children}
    </Spin>
  );
};

export default LoadingProvider;
