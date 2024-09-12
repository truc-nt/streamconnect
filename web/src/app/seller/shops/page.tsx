"use client";
import React from "react";
import { Space, Table, Tag, Avatar } from "antd";
import type { TableProps } from "antd";
import { IExternalShop } from "@/api/external_shop";
import { useGetExternalShops } from "@/hook/external_shop";
import { SyncOutlined, DeleteOutlined } from "@ant-design/icons";
import { syncExternalVariants } from "@/api/external_product";
import useLoading from "@/hook/loading";
import { useAppSelector } from "@/store/store";

const Page = () => {
  const { userId } = useAppSelector((state) => state.authReducer);
  const handleSyncExternalVariants = useLoading(
    () => syncExternalVariants(userId!),
    "Đồng bộ sản phẩm thành công",
    "Đồng bộ sản phẩm thất bại",
  );

  const columns: TableProps<IExternalShop>["columns"] = [
    {
      title: "Tên shop",
      dataIndex: "name",
      //key: "name",
      render: (_, { name }) => {
        return (
          <Space>
            <Avatar
              src="/asset/img/logo_shopify.jpg"
              alt="Shopify Logo"
              className="p-1"
            />
            <span>{name}</span>
          </Space>
        );
      },
    },
    {
      title: "Trạng thái",
      //dataIndex: 'status',
      key: "status",
      render: (_, { status }) => {
        switch (status) {
          case "active":
            return <Tag color="success">Đang hoạt động</Tag>;
          case "inactive":
            return <Tag color="error">Ngưng hoạt động</Tag>;
        }
      },
    },
    {
      title: "Ngày kết nối",
      //dataIndex: 'created_at',
      key: "created_at",
      render: (_, { created_at }) => {
        return new Date(created_at).toLocaleString();
      },
    },
    {
      title: "Ngày kết nối",
      //dataIndex: 'updated_at',
      key: "updated_at",
      render: (_, { updated_at }) => {
        return new Date(updated_at).toLocaleString();
      },
    },
    {
      title: "",
      key: "action",
      render: () => (
        <Space>
          <SyncOutlined onClick={handleSyncExternalVariants} />
          <DeleteOutlined />
        </Space>
      ),
    },
  ];

  const { data: externalShops } = useGetExternalShops(userId!);
  return (
    <>
      <Table
        columns={columns}
        dataSource={externalShops}
        rowKey={(row) => row.id_external_shop}
      />
    </>
  );
};

export default Page;
