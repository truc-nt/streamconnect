import React, { useEffect, useState } from "react";
import { InputNumber, Table, Space, TableProps, Avatar } from "antd";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ICart } from "@/api/cart";
import { IBaseCartItem } from "@/model/cart";
import Tag from "@/component/core/Tag";
import Image from "next/image";
import type { Key } from "react";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { updateQuantity } from "@/api/cart";
import useLoading from "@/hook/loading";
import { ECOMMERCE_LOGOS } from "@/constant/ecommerce";
import { setGroupByShop } from "@/store/checkout";

const CartItemsGroupByShop = ({ id_shop, shop_name, cart_items }: ICart) => {
  const [cartItems, setCartItems] = useState<IBaseCartItem[]>(cart_items);
  const handleUpdateQuantity = useLoading(
    updateQuantity,
    "Thay đổi số lượng thành công",
    "Thay đổi số lượng thất bại",
  );

  const { groupByShop } = useAppSelector((state) => state.checkoutReducer);
  const dispatch = useAppDispatch();

  useEffect(() => {
    const updatedGroupByEcommerce =
      groupByShop
        .find((cart) => cart.shopId === id_shop)
        ?.groupByEcommerce.map((ecommerce) => {
          const updatedCartItemIds = ecommerce.cartItemIds.filter((id) =>
            cartItems.some((item) => item.id_cart_item === id),
          );
          const subTotal = updatedCartItemIds.reduce((total, cartItemId) => {
            const item = cartItems.find(
              (item) => item.id_cart_item === cartItemId,
            );
            return total + (item ? item.price * item.quantity : 0);
          }, 0);

          return { ...ecommerce, cartItemIds: updatedCartItemIds, subTotal };
        }) || [];

    dispatch(
      setGroupByShop(
        groupByShop.map((cart) =>
          cart.shopId === id_shop
            ? { ...cart, groupByEcommerce: updatedGroupByEcommerce }
            : cart,
        ),
      ),
    );
  }, [cartItems]);

  const columns: TableProps<IBaseCartItem>["columns"] = [
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
    selectedRowKeys:
      groupByShop
        .find((cart) => cart.shopId === id_shop)
        ?.groupByEcommerce.flatMap((ecommerce) => ecommerce.cartItemIds) || [],

    onChange: (keys: Key[]) => {
      const newCartItemIds: {
        cartItemId: number;
        price: number;
        quantity: number;
      }[] = keys
        .map((key) => {
          const cartItem = cartItems.find((item) => item.id_cart_item === key);
          return cartItem
            ? {
                cartItemId: cartItem.id_cart_item,
                price: cartItem.price,
                quantity: cartItem.quantity,
              }
            : null;
        })
        .filter(Boolean) as any;

      const existingCartGroup = groupByShop.find(
        (cart) => cart.shopId === id_shop,
      );

      const newEcommerceEntries = newCartItemIds.reduce(
        (acc, item) => {
          const cartItem = cartItems.find(
            (cartItem) => cartItem.id_cart_item === item.cartItemId,
          );
          const ecommerceId = cartItem?.id_ecommerce;

          if (ecommerceId !== undefined) {
            if (!acc[ecommerceId]) {
              acc[ecommerceId] = {
                cartItemIds: [],
                ecommerceId,
                subTotal: 0,
              };
            }
            acc[ecommerceId].cartItemIds.push(item.cartItemId);
            acc[ecommerceId].subTotal += item.price * item.quantity;
          }
          return acc;
        },
        {} as Record<
          number,
          { cartItemIds: number[]; ecommerceId: number; subTotal: number }
        >,
      );

      const updatedGroupByEcommerce = Object.values(newEcommerceEntries);
      const updatedGroupByShop = existingCartGroup
        ? groupByShop.map((cart) =>
            cart.shopId === id_shop
              ? {
                  ...cart,
                  groupByEcommerce: [
                    ...cart.groupByEcommerce.filter(
                      (ecommerce) =>
                        !updatedGroupByEcommerce.some(
                          (newEntry) =>
                            newEntry.ecommerceId === ecommerce.ecommerceId,
                        ),
                    ),
                    ...updatedGroupByEcommerce,
                  ],
                }
              : cart,
          )
        : [
            ...groupByShop,
            { shopId: id_shop, groupByEcommerce: updatedGroupByEcommerce },
          ];

      dispatch(setGroupByShop(updatedGroupByShop));
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
