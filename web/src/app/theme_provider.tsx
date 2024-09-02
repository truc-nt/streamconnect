"use client";
import React from "react";
import { ConfigProvider, theme, App } from "antd";

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
    <App>{children}</App>
  </ConfigProvider>
);

export default ThemeProvider;
