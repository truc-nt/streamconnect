import { Flex, Row, Col, Typography, Button } from "antd";
import Image from "next/image";
import { useState, useEffect } from "react";
import { ECOMMERCE_PLATFORMS } from "@/constant/ecommerce";
import { getProductOptionFromVariantOptions } from "@/util/product";

export interface IExternalVariant {
  id_external_variant: number;
  id_ecommerce: number;
  price: number;
  stock: number;
}

interface IVariant {
  id_variant: number;
  option: { [key: string]: string };
  external_variants: IExternalVariant[];
}

export interface IProductInformation {
  productId: number;
  name: string;
  image_url: string;
  variants: IVariant[];
}

interface IProductInformationProps {
  product: IProductInformation;
  handleChangeOption?: (option: { [key: string]: string }) => void;
}

const ProductInformation = ({
  product,
  handleChangeOption,
}: IProductInformationProps) => {
  const [chosenOption, setChosenOption] = useState<{ [key: string]: string }>(
    {},
  );

  const option = getProductOptionFromVariantOptions(
    product.variants.map((variant) => variant.option),
  );

  useEffect(() => {
    setChosenOption({});
  }, [product.name]);

  return (
    <Row gutter={[16, 16]}>
      <Col span="12">
        <div className="h-[300px] relative">
          <Image
            src={product.image_url ?? ""}
            alt={product.name}
            layout="fill"
            objectFit="contain"
          />
        </div>
      </Col>
      <Col span="12">
        <Flex gap="middle" vertical>
          <Typography.Title level={3} style={{ margin: 0 }}>
            {product.name}
          </Typography.Title>
          {Object.entries(option || {}).map(([key, values]) => (
            <div key={key}>
              <Typography.Title level={5}>{key}</Typography.Title>
              <Flex gap="small">
                {values.map((value) => (
                  <Button
                    key={value}
                    size="small"
                    type={chosenOption[key] === value ? "primary" : "default"}
                    onClick={() => {
                      if (typeof handleChangeOption === "function")
                        handleChangeOption({ ...chosenOption, [key]: value });
                      setChosenOption({ ...chosenOption, [key]: value });
                    }}
                  >
                    {value}
                  </Button>
                ))}
              </Flex>
            </div>
          ))}
        </Flex>
      </Col>
    </Row>
  );
};

export default ProductInformation;
