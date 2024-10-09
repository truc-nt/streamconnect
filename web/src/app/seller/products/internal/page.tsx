"use client";
import { useEffect, useState } from "react";
import { Table, Flex, Switch, TableProps, Tag } from "antd";
import { useGetProductsByShopId } from "@/hook/product";
import { IBaseProduct } from "@/model/product";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import Image from "next/image";

import ProductModal from "@/component/modal/ProductModal";
import { useAppSelector } from "@/store/store";
import useLoading from "@/hook/loading";
import { updateProduct } from "@/api/product";

const Page = () => {
  const { userId } = useAppSelector((state) => state.authReducer);
  const { data: products } = useGetProductsByShopId(userId!);
  const [productId, setProductId] = useState<number | null>(null);
  const [isOpenModal, setIsOpenModal] = useState(false);

  const executeUpdateProduct = useLoading(
    updateProduct,
    "Cập nhật thông tin sản phẩm thành công",
    "Cập nhật thông tin sản phẩm thất bại",
  );

  const columns: TableProps<IBaseProduct>["columns"] = [
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
      render: (_, { id_product, status }) => (
        <Flex gap="small">
          <Switch
            defaultChecked={status === "active"}
            onChange={async (status) => {
              await executeUpdateProduct(id_product, {
                status: status ? "active" : "inactive",
              });
            }}
          />
        </Flex>
      ),
    },
    {
      dataIndex: "action",
      key: "action",
      render: (_, { id_product }) => {
        return (
          <Flex gap="small">
            <EditOutlined
              onClick={() => {
                setIsOpenModal(true);
                setProductId(id_product);
              }}
            />
          </Flex>
        );
      },
    },
  ];

  return (
    <>
      <Table
        columns={columns}
        dataSource={products || []}
        rowKey={(row) => row.id_product}
      />
      {productId && (
        <ProductModal
          productId={productId!}
          open={isOpenModal}
          onCancel={() => setIsOpenModal(false)}
        />
      )}
    </>
  );
};

export default Page;
