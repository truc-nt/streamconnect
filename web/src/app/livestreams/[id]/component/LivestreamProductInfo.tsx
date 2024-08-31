"use client";
import {
  Modal,
  Row,
  Col,
  Flex,
  Typography,
  Button,
  Space,
  InputNumber,
  theme,
  Divider,
} from "antd";
import { useGetLivestreamProduct } from "@/hook/livestream_product";
import { useState } from "react";
import { ECOMMERCE_PLATFORMS } from "@/constant/ecommerce";
import QuantityInput from "@/component/core/QuantityInput";

const LivestreamProductInfo = ({
  livestreamProductId,
  setLivestreamProductId,
}: {
  livestreamProductId: number | null;
  setLivestreamProductId: (value: number | null) => void;
}) => {
  const { token } = theme.useToken();
  const { data, error } = useGetLivestreamProduct(livestreamProductId);

  const [chosenOption, setChosenOption] = useState<Record<string, string>>({});
  const [chosenEcommerce, setChosenEcommerce] = useState<number | null>(null);
  const livestreamProduct = data?.data;

  const livestreamVariant = livestreamProduct?.livestream_variants.find(
    (variant) =>
      Object.entries(variant.option).every(
        ([key, value]) => chosenOption[key] === value,
      ),
  );
  const livestreamExternalVariant =
    livestreamVariant?.livestream_external_variants.find(
      (variant) => variant.id_ecommerce === chosenEcommerce,
    );

  return (
    <Modal
      centered
      open={livestreamProductId !== null}
      footer={null}
      width="60%"
      onCancel={() => {
        setLivestreamProductId(null);
      }}
    >
      <Row>
        <Col span={12}>col</Col>
        <Col span={12}>
          <Flex vertical gap="large">
            <Typography.Title level={3} style={{ margin: 0 }}>
              {livestreamProduct?.name}
            </Typography.Title>
            {Object.entries(livestreamProduct?.option || {}).map(
              ([key, values]) => (
                <div key={key}>
                  <Typography.Title level={5}>{key}</Typography.Title>
                  <Flex gap="small">
                    {values.map((value) => (
                      <Button
                        key={value}
                        size="small"
                        type={
                          chosenOption[key] === value ? "primary" : "default"
                        }
                        onClick={() => setChosenOption({ [key]: value })}
                      >
                        {value}
                      </Button>
                    ))}
                  </Flex>
                </div>
              ),
            )}
            <div>
              {livestreamVariant?.livestream_external_variants && (
                <Typography.Title level={5}>Sàn thương mại</Typography.Title>
              )}
              {livestreamVariant?.livestream_external_variants.map(
                (variant) => (
                  <Flex
                    key={variant.id_livestream_external_variant}
                    gap="small"
                  >
                    <Button
                      key={variant.id_ecommerce}
                      size="small"
                      onClick={() => setChosenEcommerce(variant.id_ecommerce)}
                      type={
                        chosenEcommerce === variant.id_ecommerce
                          ? "primary"
                          : "default"
                      }
                    >
                      {ECOMMERCE_PLATFORMS[variant.id_ecommerce]}
                    </Button>
                  </Flex>
                ),
              )}
            </div>
            {
              <Typography.Title
                level={2}
                style={{ color: token.colorPrimaryText, margin: 0 }}
              >
                {livestreamExternalVariant?.price}
              </Typography.Title>
            }
            <Space size="small">
              <QuantityInput onDecrease={() => {}} onIncrease={() => {}} />
              <Button type="primary" size="large">
                Thêm vào giỏ
              </Button>
            </Space>
          </Flex>
        </Col>
      </Row>
      <Divider />
      <Divider />
      <Typography.Text style={{ fontWeight: "bold" }}>
        Giới thiệu sản phẩm
      </Typography.Text>
      <p>viuebvibeiovbiwovbiowebvioewbviwbeivobweriovbiweobviow</p>
    </Modal>
  );
};

export default LivestreamProductInfo;
