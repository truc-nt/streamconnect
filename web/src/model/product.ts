interface IBaseProduct {
  name: string;
  description: string;
  status: string;
}
export interface IBaseVariant extends IBaseProduct {
  id_variant: number;
  status: string;
  option: { [key: string]: string };
  image_url: string;
}

export interface IBaseExternalVariant extends IBaseVariant {
  id_external_variant: number;
  id_ext_shop: number;
  ext_product_id_mapping: string;
  ext_id_mapping: string;
  sku: string;
  status: string;
  option: { [key: string]: string };
  price: number;
  id_ecommerce: number;
}
