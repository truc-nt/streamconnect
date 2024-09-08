import { Flex } from "antd";
import AddLivestreamVariant from "./AddLivestreamExternalVariantModal";
import { Button, Table, Tag, Space, TableColumnType } from "antd";
import { useState } from "react";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ECOMMERCE_PLATFORMS } from "@/constant/ecommerce";

interface IChosenLivestreamVariant {
  productId: number;
  variantId: number;
  name: string;
  imageUrl: string;
  option: { [key: string]: string };
  externalVariants: {
    externalVariantId: number;
    ecommerceId: number;
    price: number;
    quantity: number;
  }[];
}

interface IChosenLivestreamVariantProps {
  shopId: number;
  initialChosenLivestreamVariants?: IChosenLivestreamVariant[];
}
const ChosenLivestreamVariant = ({
  shopId,
  initialChosenLivestreamVariants,
}: IChosenLivestreamVariantProps) => {
  const [openAddModal, setOpenAddModal] = useState(false);
  const [chosenLivestreamVariants, setChosenLivestreamVariants] = useState(
    initialChosenLivestreamVariants ?? [],
  );
  return (
    <Flex vertical gap="large">
      <Flex>
        <Button type="primary" onClick={() => setOpenAddModal(true)}>
          Thêm sản phẩm
        </Button>
      </Flex>
      <ChosenLivestreamVariantTable data={chosenLivestreamVariants} />
      {openAddModal && (
        <AddLivestreamVariant
          shopId={shopId}
          chosenLivestreamVariants={chosenLivestreamVariants}
          setChosenLivestreamVariants={setChosenLivestreamVariants}
          onCancel={() => setOpenAddModal(false)}
        />
      )}
      <Flex justify="end">
        <Button
          type="primary"
          //disabled={}
          //onClick={}
        >
          Tiếp theo
        </Button>
      </Flex>
    </Flex>
  );
};

const ChosenLivestreamVariantTable = ({
  data,
}: {
  data: IChosenLivestreamVariant[];
}) => {
  const columns: TableColumnType<IChosenLivestreamVariant>[] = [
    {
      title: "Tên sản phẩm",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "Định dạng",
      dataIndex: "option",
      key: "option",
      render: (_, { option }) =>
        Object.entries(option).map(([key, value]) => (
          <Tag key={key}>
            {key}: {value}
          </Tag>
        )),
    },
    {
      title: "Giá",
      dataIndex: "price",
      key: "price",
      render: (_, { externalVariants }) =>
        Object.values(externalVariants).map((externalVariant) => (
          <Tag key={externalVariant.externalVariantId}>
            {ECOMMERCE_PLATFORMS[externalVariant.ecommerceId]}:{" "}
            {externalVariant.price}
          </Tag>
        )),
    },
    {
      title: "Số lượng",
      dataIndex: "quantity",
      key: "quantity",
      render: (_, { externalVariants }) =>
        Object.values(externalVariants).map((externalVariant) => (
          <Tag key={externalVariant.externalVariantId}>
            {ECOMMERCE_PLATFORMS[externalVariant.ecommerceId]}:{" "}
            {externalVariant.quantity}
          </Tag>
        )),
    },
    {
      dataIndex: "action",
      key: "action",
      render: () => (
        <Space>
          <EditOutlined />
          <DeleteOutlined />
        </Space>
      ),
    },
  ];

  return (
    <Table
      columns={columns}
      dataSource={data}
      rowKey={(row) => row.variantId}
    />
  );
};

export default ChosenLivestreamVariant;
