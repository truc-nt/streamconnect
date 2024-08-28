"use client";
import { useEffect, useState } from "react";
import { Table, Flex, TableProps, Tag } from "antd";
import { useGetExternalProducts } from "@/hook/external_product";
import { IExternalProduct } from "@/api/external_product";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import Image from "next/image";
import MappingExternalProduct from "./component/MappingExternalProduct";

const Page = () => {
  const { data: externalProducts } = useGetExternalProducts();
  const [externalProductIdMapping, setExternalProductIdMapping] = useState<
    string | null
  >(null);

  const columns: TableProps<IExternalProduct>["columns"] = [
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
      title: "Tên cửa hàng",
      dataIndex: "shop_name",
      key: "shop_name",
    },
    {
      title: "Tên hệ thống",
      dataIndex: "product_name",
      key: "product_name",
      render: (_, { product_name }) => {
        if (!product_name) return <span>Chưa được liên kết</span>;
        return <span>{product_name}</span>;
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
      title: "Giá",
      dataIndex: "price",
      key: "price",
      render: (_, { external_variants }) => {
        const minPrice = external_variants.reduce((min, external_variant) => {
          return external_variant.price < min ? external_variant.price : min;
        }, Infinity);
        const maxPrice = external_variants.reduce((max, external_variant) => {
          return external_variant.price > max ? external_variant.price : max;
        }, -Infinity);
        return (
          <span>
            {minPrice} - {maxPrice}
          </span>
        );
      },
    },
    {
      dataIndex: "action",
      key: "action",
      render: (_, { external_product_id_mapping }) => {
        return (
          <Flex gap="small">
            <EditOutlined
              onClick={() =>
                setExternalProductIdMapping(external_product_id_mapping)
              }
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
        dataSource={externalProducts?.data || []}
        rowKey={(row) => row.external_product_id_mapping}
        rowSelection={{}}
      />
      {externalProductIdMapping !== null && (
        <MappingExternalProduct
          externalProductIdMapping={externalProductIdMapping}
          setExternalProductIdMapping={setExternalProductIdMapping}
        />
      )}
    </>
  );
};

export default Page;
