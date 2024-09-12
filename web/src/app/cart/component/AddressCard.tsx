import { Card, Typography, Divider, Flex } from "antd";
import { useGetDefaultAddress } from "@/hook/user";
import { useAppSelector, useAppDispatch } from "@/store/store";
import { useEffect } from "react";

const AddressCard = () => {
  const dispatch = useAppDispatch();
  const { userId } = useAppSelector((state) => state.authReducer);
  const { data } = useGetDefaultAddress();

  /*useEffect(() => {
    if (data) {
      dispatch(setUserAddressId(data.id_user_address));
    }
  }, []);*/

  return (
    <Card
      bordered={false}
      styles={{
        body: {
          display: "flex",
          flexDirection: "column",
          height: "100%",
          gap: "0.5rem",
        },
      }}
    >
      <Card.Meta
        title={
          <Flex justify="space-between">
            <Typography.Text style={{ fontSize: "16px" }}>
              Địa chỉ
            </Typography.Text>
            <Typography.Text>Thay đổi</Typography.Text>
          </Flex>
        }
      />
      <Flex justify="space-between">
        <Typography.Text>{data?.name}</Typography.Text>
        <Typography.Text>{data?.phone}</Typography.Text>
      </Flex>
      <Typography.Paragraph
        ellipsis={{ rows: 2, expandable: false }}
        style={{ margin: 0 }}
      >
        {data?.address}
      </Typography.Paragraph>
    </Card>
  );
};

export default AddressCard;
