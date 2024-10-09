"use client";
import { Card, Typography, Divider, Flex, theme } from "antd";
import {
  CheckCircleOutlined,
  EditOutlined,
  DeleteOutlined,
} from "@ant-design/icons";

import { IBaseUserAddress } from "@/model/order";

const AddressInfo = ({
  name,
  phone,
  address,
  is_default,
}: IBaseUserAddress) => {
  const { token } = theme.useToken();
  return (
    <Flex justify="space-between">
      <Flex vertical>
        <Typography.Text>
          <Flex gap="middle">
            {name?.toUpperCase()}{" "}
            {is_default && (
              <div style={{ color: token.colorPrimaryText }}>
                <CheckCircleOutlined /> {"Địa chỉ mặc định"}
              </div>
            )}
          </Flex>
        </Typography.Text>
        <Typography.Text>Số điện thoại: {phone}</Typography.Text>
        <Typography.Paragraph
          ellipsis={{ rows: 2, expandable: false }}
          style={{ margin: 0 }}
        >
          Địa chỉ: {address}
        </Typography.Paragraph>
      </Flex>
      <Flex>
        <Flex gap="small">
          <EditOutlined />
          <DeleteOutlined />
        </Flex>
      </Flex>
    </Flex>
  );
};

export default AddressInfo;
