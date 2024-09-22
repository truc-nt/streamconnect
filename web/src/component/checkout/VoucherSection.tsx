import React, { useState } from "react";

import { Card, Form, Input, Button } from "antd";

import VoucherModal from "./VoucherModal";

const VoucherSection = ({
  shopId,
  ecommerceId,
}: {
  shopId: number;
  ecommerceId: number;
}) => {
  const [openChoseVoucherModal, setOpenChoseVoucherModal] = useState(false);
  return (
    <>
      <Card
        styles={{
          body: {
            display: "flex",
            flexDirection: "column",
            height: "100%",
            gap: "0.5rem",
          },
        }}
      >
        <Card.Meta title="Mã giảm giá" />
        <Form layout="vertical">
          <Form.Item label="Mã giảm giá sàn">
            <Button onClick={() => setOpenChoseVoucherModal(true)}>
              Chọn mã giảm giá
            </Button>
          </Form.Item>
          <Form.Item label="Mã giảm giá hệ thống">
            <Input />
          </Form.Item>
        </Form>
      </Card>

      <VoucherModal
        shopId={shopId}
        ecommerceId={ecommerceId}
        open={openChoseVoucherModal}
        onCancel={() => setOpenChoseVoucherModal(false)}
      />
    </>
  );
};

export default VoucherSection;
