import { Card, Typography, Divider, Flex, Button, theme } from "antd";
import {
  EditOutlined,
  DeleteOutlined,
  CheckCircleOutlined,
} from "@ant-design/icons";
import { IUserAddress } from "@/api/user";

const AddressInfo = ({
  name,
  phone,
  city,
  address,
  is_default,
}: IUserAddress) => {
  console.log("AddressInfo", name, phone, city, address);
  const { token } = theme.useToken();
  return (
    <Card className="w-full">
      <Flex justify="space-between">
        <Flex vertical>
          <Typography.Text>
            <Flex gap="middle">
              {name.toUpperCase()}{" "}
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
    </Card>
  );
};

export default AddressInfo;
