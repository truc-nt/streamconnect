import React, { useState } from "react";
import { InputNumber, Table, Space, TableProps, Avatar } from "antd";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ICart, ICartItem } from "@/api/cart";
import Tag from "@/component/core/Tag";
import Image from "next/image";
import type { Key } from "react";
import { setSelectedCartItems } from "@/store/cart_item_ids_selection";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { updateQuantity } from "@/api/cart";
import useLoading from "@/hook/loading";
import { ECOMMERCE_LOGOS } from "@/constant/ecommerce";

const CartItemsGroupByShop = ({ shop_name, cart_items }: ICart) => {
  const [cartItems, setCartItems] = useState<ICartItem[]>(cart_items);
  const handleUpdateQuantity = useLoading(
    updateQuantity,
    "Thay đổi số lượng thành công",
    "Thay đổi số lượng thất bại",
  );

  const { cartItems: selectedCartItems } = useAppSelector(
    (state) => state.cartItemIdsSelection,
  );
  const dispatch = useAppDispatch();

  const columns: TableProps<ICartItem>["columns"] = [
    {
      title: () => <span>{shop_name}</span>,
      dataIndex: "name",
      key: "name",
      render: (_, { name, image_url }) => (
        <Space size="middle" align="center">
          <Image src={image_url} alt={name} width={50} height={50} />
          <span>{name}</span>
        </Space>
      ),
    },
    {
      dataIndex: "id_ecommerce",
      key: "id_ecommerce",
      render: (id_ecommerce) => (
        <Avatar
          src={ECOMMERCE_LOGOS[id_ecommerce]}
          alt="Shopify Logo"
          size={40}
        />
      ),
    },
    {
      dataIndex: "option",
      key: "option",
      render: (option) => (
        <Space.Compact block>
          {Object.entries(option).map(([key, value]) => (
            <Tag key={key} label={`${key}: ${value}`} />
          ))}
        </Space.Compact>
      ),
    },
    {
      dataIndex: "price",
      key: "price",
      render: (price) => <span>{price}</span>,
    },
    {
      dataIndex: "quantity",
      key: "quantity",
      render: (_, { id_cart_item, quantity, max_quantity }) => (
        <InputNumber
          min={1}
          max={max_quantity}
          value={quantity}
          onChange={async (value) => {
            const newQuantity = value ?? 1;
            try {
              await handleUpdateQuantity(id_cart_item, newQuantity);
              setCartItems((prevItems) =>
                prevItems.map((item) =>
                  item.id_cart_item === id_cart_item
                    ? { ...item, quantity: newQuantity }
                    : item,
                ),
              );

              dispatch(
                setSelectedCartItems(
                  selectedCartItems.map((selectedCartItem) =>
                    selectedCartItem.cartItemId === id_cart_item
                      ? { ...selectedCartItem, quantity: newQuantity }
                      : selectedCartItem,
                  ),
                ),
              );
            } catch (error) {
              console.error("Update quantity failed:", error);
            }
          }}
        />
      ),
    },
    {
      key: "total_price",
      render: (_, { quantity, price }) => <span>{price * quantity}</span>,
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

  const rowSelection = {
    selectedRowKeys: selectedCartItems.map((item) => item.cartItemId),
    onChange: (keys: Key[]) => {
      const newCartItemIds: {
        cartItemId: Key;
        price: number;
        quantity: number;
      }[] = [];
      keys.forEach((key) => {
        const cartItem = cart_items.find((item) => item.id_cart_item === key);
        if (cartItem) {
          newCartItemIds.push({
            cartItemId: cartItem.id_cart_item,
            price: cartItem.price,
            quantity: cartItem.quantity,
          });
        }
      });

      selectedCartItems.forEach((cartItem) => {
        if (!keys.includes(cartItem.cartItemId)) {
          const existingIndex = newCartItemIds.findIndex(
            (item) => item.cartItemId === cartItem.cartItemId,
          );
          if (existingIndex > -1) {
            newCartItemIds.splice(existingIndex, 1);
          }
        }
      });

      dispatch(setSelectedCartItems(newCartItemIds));
    },
  };

  return (
    <Table
      columns={columns}
      dataSource={cartItems}
      rowSelection={rowSelection}
      rowKey={(row) => row.id_cart_item}
      pagination={false}
    />
  );
};

export default CartItemsGroupByShop;
