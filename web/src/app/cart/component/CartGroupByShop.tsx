import { Button, Input, Table, Space } from "antd";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import QuantityInput from "@/component/core/QuantityInput";
import { ICart, ICartItem } from "@/api/cart";
import { TableProps } from "antd";
import Tag from "@/component/core/Tag";

const CartGroupByShop = ({
  id_shop,
  shop_name,
  cart_livestream_external_variant,
}: ICart) => {
  console.log("hello", id_shop);
  const columns: TableProps<ICartItem>["columns"] = [
    {
      title: () => <span>{shop_name}</span>,
      dataIndex: "name",
      key: "name",
    },
    {
      dataIndex: "option",
      key: "option",
      render: (option) => {
        return (
          <Space.Compact block>
            {Object.entries(option).map(([key, value]) => (
              <Tag key={key} label={`${key}: ${value}`} />
            ))}
          </Space.Compact>
        );
      },
    },
    {
      dataIndex: "price",
      key: "price",
    },
    {
      dataIndex: "quantity",
      key: "quantity",
      render: (_, { quantity }) => (
        <Space.Compact block>
          <QuantityInput
            quantity={quantity}
            onDecrease={() => {}}
            onIncrease={() => {}}
          />
        </Space.Compact>
      ),
    },
    {
      key: "total_price",
      render: (_, { quantity, price }) => {
        return <span>{price * quantity}</span>;
      },
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
      //showHeader={false}
      columns={columns}
      dataSource={cart_livestream_external_variant}
      rowSelection={{}}
    />
  );
};

export default CartGroupByShop;
