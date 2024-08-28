import { useGetExternalVariants } from "@/hook/external_product";
import { useGetProductsByShopId } from "@/hook/product";
import {
  Modal,
  Row,
  Col,
  Divider,
  Flex,
  Typography,
  Table,
  Select,
  Button,
} from "antd";
import { getProductOptionFromVariantOptions } from "@/util/product";
import Image from "next/image";
import { createProductWithVariants } from "@/api/product";
import useLoading from "@/hook/loading";

const MappingExternalProduct = ({
  externalProductIdMapping,
  setExternalProductIdMapping,
}: {
  externalProductIdMapping: string | null;
  setExternalProductIdMapping: (value: string | null) => void;
}) => {
  const { data: externalVariants } = useGetExternalVariants(
    externalProductIdMapping!,
  );
  const { data: products } = useGetProductsByShopId(1);
  const productOption = getProductOptionFromVariantOptions(
    externalVariants?.data.map((externalVariant) => externalVariant.option) ??
      [],
  );

  const handleCreateProductWithVariants = useLoading(
    () =>
      createProductWithVariants(1, [
        { external_product_id_mapping: externalProductIdMapping ?? "" },
      ]),
    "Tạo sản phẩm mới thành công",
    "Tạo sản phẩm mới thất bại",
  );

  const columns = [
    {
      dataIndex: "image_url",
      key: "image_url",
      render: (_, { image_url }) => (
        <Image src={image_url} alt="" width={30} height={30} />
      ),
    },
    ...Object.keys(productOption).map((key) => ({
      title: key,
      dataIndex: key,
    })),
  ];

  const dataSource = externalVariants?.data.map((externalVariant) => {
    return {
      ...externalVariant.option,
      image_url: externalVariant.image_url,
    };
  });

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
      <Row align="middle" justify="center">
        <Col span={11}>
          <Table
            columns={columns}
            dataSource={dataSource}
            size="small"
            pagination={false}
          />
        </Col>
        <Divider type="vertical" />
        <Col span={11}>
          <Flex gap="small">
            <Select
              showSearch
              placeholder="Chọn sản phẩm"
              optionFilterProp="label"
              options={
                products?.data.map((product) => ({
                  label: product.name,
                  value: product.id_product,
                })) ?? []
              }
            />
            <Button>Liên kết</Button>
            <Button type="primary" onClick={handleCreateProductWithVariants}>
              Tạo mới
            </Button>
          </Flex>
          <Table />
        </Col>
      </Row>
    </Modal>
  );
};

export default MappingExternalProduct;
