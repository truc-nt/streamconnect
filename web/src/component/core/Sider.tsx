"use client";
import { Layout } from "antd";
import { useState } from "react";
import Menu from "@/component/core/Menu";
import Image from "next/image";
import Link from "next/link";

const Sider = () => {
  const [collapsed, setCollapsed] = useState(false);

  return (
    <Layout.Sider
      theme="light"
      collapsible
      collapsed={collapsed}
      onCollapse={(value) => setCollapsed(value)}
      className="overflow-auto h-screen fixed inset-0 scrollbar-thin scrollbar-color-transparent"
    >
      <div className="relative w-full h-16 p-2">
        <div className="relative w-full h-full">
          <Link href="/">
            <Image
              src={
                collapsed
                  ? "/asset/img/logo.svg"
                  : "/asset/img/logo_with_name.svg"
              }
              alt="Logo"
              layout="fill"
              objectFit="contain"
              className="object-contain"
            />
          </Link>
        </div>
      </div>
      <Menu />
    </Layout.Sider>
  );
};
export default Sider;
