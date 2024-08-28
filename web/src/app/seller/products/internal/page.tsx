"use client";
import { useEffect, useState } from "react";
import { Table, Flex, TableProps, Tag } from "antd";
import { useGetProductsByShopId } from "@/hook/product";
import { IProduct } from "@/api/product";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import Image from "next/image";

const Page = () => {
  const { data: products } = useGetProductsByShopId(1);

  const columns: TableProps<IProduct>["columns"] = [
    {
      title: "Tên",
      dataIndex: "name",
      key: "name",
      render: (_, { name, image_url }) => {
        return (
          <Flex gap="middle" align="center">
            <Image src={image_url} alt={name} width={50} height={50} />
            <span>{name}</span>
          </Flex>
        );
      },
    },
    {
      title: "Trạng thái",
      dataIndex: "status",
      key: "status",
      render: (_, { status }) => {
        if (status === "active") {
          return <Tag color="green">Đang hoạt động</Tag>;
        }
        return <Tag color="red">Ngừng hoạt động</Tag>;
      },
    },
    {
      dataIndex: "action",
      key: "action",
      render: (_, { external_product_id_mapping }) => {
        return (
          <Flex gap="small">
            <EditOutlined
            /*onClick={() =>
                setExternalProductIdMapping(external_product_id_mapping)
              }*/
            />
            <DeleteOutlined />
          </Flex>
        );
      },
    },
  ];

  return (
    <>
      <Table
        columns={columns}
        dataSource={products?.data || []}
        rowKey={(row) => row.id_product}
        rowSelection={{}}
      />
    </>
  );
};

export default Page;
