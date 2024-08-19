"use client";
import React from "react";
import { ConfigProvider, theme } from "antd";

const ThemeProvider = ({ children }: { children: React.ReactNode }) => (
  <ConfigProvider
    theme={{
      token: {
        //colorPrimary: '#EA6B62',
        //colorBgContainer: '#E33E33',
        colorPrimary: "#E33E33",
      },
    }}
    //componentSize="large"
  >
    {children}
  </ConfigProvider>
);

export default ThemeProvider;
