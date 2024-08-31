"use client";
import { Space, Flex, Table, Button, InputNumber, Input } from "antd";
import { PlusOutlined, MinusOutlined } from "@ant-design/icons";

const Page = () => {
  const columns = [
    {
      title: () => <span>Shop Name</span>,
      dataIndex: "name",
      key: "name",
    },
    {
      title: "",
      dataIndex: "price",
      key: "price",
    },
    {
      title: "",
      dataIndex: "quantity",
      key: "quantity",
      render: () => (
        <Space.Compact block>
          <Button icon={<MinusOutlined />} size="small" />
          <Input size="small" defaultValue={1} style={{ width: "30px" }} />
          <Button icon={<PlusOutlined />} size="small" />
        </Space.Compact>
      ),
    },
    {
      title: "",
      dataIndex: "total",
      key: "total",
    },
    {
      title: "",
      dataIndex: "action",
      key: "action",
    },
  ];
  const data = [
    {
      key: "1",
      name: "John Brown",
      price: 32,
      quantity: 2,
      action: "Edit",
    },
  ];
  return (
    <Space.Compact direction="vertical" style={{ display: "flex" }}>
      <div></div>
      <Flex vertical>
        <Table
          //showHeader={false}
          columns={columns}
          dataSource={data}
          rowSelection={{}}
        />
      </Flex>
    </Space.Compact>
  );
};

export default Page;
