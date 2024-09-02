import {
  useGetExternalVariants,
  useGetVariantsByExternalProductIdMapping,
} from "@/hook/external_product";
import { useGetProductsByShopId } from "@/hook/product";
import {
  Modal,
  Row,
  Col,
  Divider,
  Flex,
  Table,
  Select,
  Button,
  theme,
  TableProps,
  Tag,
  InputNumber,
} from "antd";
import Image from "next/image";
import {
  createProductWithVariants,
  getVariantsByProductId,
} from "@/api/product";
import useLoading from "@/hook/loading";
import {
  IVariant,
  IExternalVariant,
  connectVariants,
} from "@/api/external_product";
import { useState, useEffect } from "react";

interface ISelectedVariant extends IVariant {
  externalVariantId: number;
}

const MappingExternalProduct = ({
  externalProductIdMapping,
  setExternalProductIdMapping,
}: {
  externalProductIdMapping: string | null;
  setExternalProductIdMapping: (value: string | null) => void;
}) => {
  const { token } = theme.useToken();
  const { data: externalVariants } = useGetExternalVariants(
    externalProductIdMapping!,
  );
  const { data: products } = useGetProductsByShopId(1);
  const { data: variants } = useGetVariantsByExternalProductIdMapping(
    externalProductIdMapping!,
  );

  const [selectedProductId, setSelectedProductId] = useState<number | null>(
    null,
  );
  const [selectedVariants, setSelectedVariants] = useState<ISelectedVariant[]>(
    [],
  );
  console.log("selectedVariants", selectedVariants);

  const handleCreateProductWithVariants = useLoading(
    () =>
      createProductWithVariants(1, [
        { external_product_id_mapping: externalProductIdMapping ?? "" },
      ]),
    "Tạo sản phẩm mới thành công",
    "Tạo sản phẩm mới thất bại",
  );

  const handleConnectVariants = useLoading(
    connectVariants,
    "Liên kết sản phẩm thành công",
    "Liên kết sản phẩm thất bại",
  );

  const externalVariantsColumns: TableProps<IExternalVariant>["columns"] = [
    {
      title: "STT",
      dataIndex: "index",
      key: "index",
      width: 20,
      render: (_, __, index) => index + 1,
    },
    {
      title: "Ảnh",
      dataIndex: "image_url",
      key: "image_url",
      width: 50,
      render: (_, { image_url }) => (
        <Image src={image_url} alt="" width={30} height={30} />
      ),
    },
    {
      title: "Định dạng",
      dataIndex: "option",
      key: "option",
      render: (option) => (
        <Flex gap="small">
          {Object.entries(option).map(([key, value]) => (
            <Tag key={key}>{`${key}: ${value}`} </Tag>
          ))}
        </Flex>
      ),
    },
  ];

  useEffect(() => {
    const selectedProductId = products?.find(
      (product) => product.id_product === variants?.[0]?.fk_product,
    )?.id_product;
    if (selectedProductId) setSelectedProductId(selectedProductId);
  }, [products, variants]);

  useEffect(() => {
    setSelectedVariants(
      variants?.map((variant) => {
        const externalVariantId = externalVariants?.find(
          (externalVariant) =>
            externalVariant.fk_variant === variant.id_variant,
        )?.id_external_variant;
        return {
          ...variant,
          externalVariantId: externalVariantId ?? 0,
        };
      }) ?? [],
    );
  }, [variants]);

  const variantsColumns: TableProps<ISelectedVariant>["columns"] = [
    {
      title: "Ảnh",
      dataIndex: "image_url",
      key: "image_url",
      width: 50,
      render: (_, { image_url }) => (
        <Image src={image_url} alt="" width={30} height={30} />
      ),
    },
    {
      title: "Định dạng",
      dataIndex: "option",
      key: "option",
      render: (option) => (
        <Flex gap="small">
          {Object.entries(option).map(([key, value]) => (
            <Tag key={key}>{`${key}: ${value}`} </Tag>
          ))}
        </Flex>
      ),
    },
    {
      title: "Liên kết",
      render: ({ id_variant, externalVariantId }) => (
        <InputNumber
          min={1}
          max={externalVariants?.length}
          defaultValue={
            (externalVariants?.findIndex(
              (externalVariant) =>
                externalVariant.id_external_variant === externalVariantId,
            ) ?? 0) + 1
          }
          onChange={(value) => {
            setSelectedVariants((prevVariants) => {
              if (value === null) return prevVariants;
              const externalVariantId =
                externalVariants?.[value - 1].id_external_variant;
              return prevVariants.map((variant) =>
                variant.id_variant === id_variant
                  ? { ...variant, externalVariantId: externalVariantId ?? 1 }
                  : variant,
              );
            });
          }}
        />
      ),
    },
  ];

  return (
    <Modal
      centered
      open={externalProductIdMapping !== null}
      footer={null}
      width="60%"
      onCancel={() => {
        setExternalProductIdMapping(null);
      }}
    >
      <Row gutter={[8, 8]} justify="center">
        <Col span={11} />
        <Col span={11}>
          <Flex gap="small">
            <Select
              showSearch
              placeholder="Chọn sản phẩm"
              optionFilterProp="label"
              value={selectedProductId}
              options={
                products?.map((product) => ({
                  label: product.name,
                  value: product.id_product,
                })) ?? []
              }
              onChange={async (value) => {
                setSelectedProductId(value);
                const variants = await getVariantsByProductId(value);
                setSelectedVariants(
                  variants?.data.map((variant) => ({
                    ...variant,
                    externalVariantId: 0,
                  })) ?? [],
                );
              }}
            />
            <Button
              onClick={() =>
                handleConnectVariants(
                  selectedVariants.map((selectedVariant) => ({
                    id_variant: selectedVariant.id_variant,
                    id_external_variant: selectedVariant.externalVariantId,
                  })),
                )
              }
            >
              Liên kết
            </Button>
            <Button type="primary" onClick={handleCreateProductWithVariants}>
              Tạo mới
            </Button>
          </Flex>
        </Col>
        <Col span={11}>
          <Table
            columns={externalVariantsColumns}
            dataSource={externalVariants}
            size="small"
            pagination={false}
            rowKey={(record) => record.id_external_variant}
          />
        </Col>
        <Col span={11}>
          {selectedVariants?.length ? (
            <Table
              columns={variantsColumns}
              dataSource={selectedVariants}
              size="small"
              pagination={false}
              rowKey={(record) => record.id_variant}
            />
          ) : (
            <div
              style={{
                display: "flex",
                height: "100%",
                alignItems: "center",
                justifyContent: "center",
                color: token.colorTextTertiary,
                backgroundColor: token.colorFillAlter,
                borderRadius: token.borderRadiusLG,
                border: `1px dashed ${token.colorBorder}`,
              }}
            >
              Hãy chọn sản phẩm để liên kết
            </div>
          )}
        </Col>
      </Row>
    </Modal>
  );
};

export default MappingExternalProduct;
