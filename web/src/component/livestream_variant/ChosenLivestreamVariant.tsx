import { Flex } from "antd";
import AddLivestreamVariant from "./AddLivestreamExternalVariantModal";
import { Button, Table, Tag, Space, TableColumnType } from "antd";
import { useState, useEffect } from "react";
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
  onChange?: (chosenLivestreamVariants: IChosenLivestreamVariant[]) => void;
  onSubmit?: (chosenLivestreamVariants: IChosenLivestreamVariant[]) => void;
}
const ChosenLivestreamVariant = ({
  shopId,
  initialChosenLivestreamVariants,
  onChange,
  onSubmit,
}: IChosenLivestreamVariantProps) => {
  const [openAddModal, setOpenAddModal] = useState(false);
  const [chosenLivestreamVariants, setChosenLivestreamVariants] = useState(
    initialChosenLivestreamVariants ?? [],
  );

  useEffect(() => {
    onChange?.(chosenLivestreamVariants);
  }, [chosenLivestreamVariants]);

  return (
    <Flex vertical gap="large">
      <Flex>
        <Button type="primary" onClick={() => setOpenAddModal(true)}>
          Thêm sản phẩm
        </Button>
      </Flex>
      <ChosenLivestreamVariantTable data={chosenLivestreamVariants} />
      {onSubmit && (
        <Button
          type="primary"
          onClick={() => onSubmit(chosenLivestreamVariants)}
        >
          Lưu
        </Button>
      )}
      {openAddModal && (
        <AddLivestreamVariant
          shopId={shopId}
          chosenLivestreamVariants={chosenLivestreamVariants}
          setChosenLivestreamVariants={setChosenLivestreamVariants}
          onCancel={() => setOpenAddModal(false)}
        />
      )}
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
  ];

  return (
    <Table
      columns={columns}
      dataSource={data}
      rowKey={(row) => row.variantId}
      pagination={false}
    />
  );
};

export default ChosenLivestreamVariant;
