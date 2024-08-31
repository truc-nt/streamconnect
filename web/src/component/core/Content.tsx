"use client";

import { Layout } from "antd";
const Content = ({ children }: { children: React.ReactNode }) => {
  return <Layout.Content className="p-8">{children}</Layout.Content>;
};

export default Content;
