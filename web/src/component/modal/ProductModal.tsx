import React, { useState } from "react";
import Image from "next/image";
import { Avatar, Modal, Row, Col, Table, TableProps, ModalProps } from "antd";
import ProductInfo from "@/component/info/ProductInfo";
import ProductForm from "@/component/form/ProductForm";
import { IBaseExternalVariant } from "@/model/product";

import { useGetProductById } from "@/hook/product";
import { ECOMMERCE_LOGOS } from "@/constant/ecommerce";

interface IProductModalProps extends ModalProps {
  productId: number;
}

const ProductModal = ({ productId, ...props }: IProductModalProps) => {
  const { data: product } = useGetProductById(productId);
  const [chosenOption, setChosenOption] = useState<{ [key: string]: string }>(
    {},
  );
  const chosenVariant = product?.variants?.find((variant) =>
    Object.entries(variant.option).every(
      ([key, value]) => chosenOption[key] === value,
    ),
  );
  const columns: TableProps<
    IBaseExternalVariant & {
      shop_name: string;
    }
  >["columns"] = [
    {
      title: "Ảnh",
      dataIndex: "image_url",
      key: "image_url",
      width: 50,
      render: (_, { image_url }) => (
        <Image src={image_url} alt="" width={30} height={30} />
      ),
    },
    {
      title: "Tên shop",
      dataIndex: "shop_name",
      key: "shop_name",
      render: (_, { id_ecommerce, shop_name }) => (
        <span>
          <Avatar
            src={ECOMMERCE_LOGOS[id_ecommerce]}
            alt="Shopify Logo"
            size={40}
          />{" "}
          {shop_name}
        </span>
      ),
    },
    {
      title: "Tồn kho",
      dataIndex: "stock",
      key: "stock",
    },
  ];

  console.log(chosenVariant);

  return (
    <Modal
      centered
      title="Thông tin sản phẩm"
      open={true}
      footer={null}
      width="80%"
      {...props}
    >
      <Row>
        <Col span={16}>
          <ProductInfo
            product={{
              productId: product?.id_product!,
              name: product?.name!,
              image_url: product?.image_url!,
              variants: product?.variants! ?? [],
            }}
            handleChangeOption={(option) => setChosenOption(option)}
          />
          <ProductForm {...product!} />
        </Col>
        <Col span={8}>
          <Table
            columns={columns}
            dataSource={chosenVariant?.external_variants}
          />
        </Col>
      </Row>
    </Modal>
  );
};

export default ProductModal;
