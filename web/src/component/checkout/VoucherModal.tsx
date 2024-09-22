import { List, Modal, Button, Radio, ModalProps } from "antd";

import VoucherItem from "@/component/list_item/VoucherItem";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { useGetShopVouchers } from "@/hook/voucher";
import { setGroupByShop } from "@/store/checkout";

interface IVoucherModalProps extends ModalProps {
  shopId: number;
  ecommerceId: number;
}

const VoucherModal = (props: IVoucherModalProps) => {
  const dispatch = useAppDispatch();
  const { data: shopVouchers } = useGetShopVouchers(props.shopId);
  const { groupByShop } = useAppSelector((state) => state.checkoutReducer);
  const subTotal = groupByShop.reduce((total, shop) => {
    const shopSubTotal = shop.groupByEcommerce.reduce((subtotal, ecommerce) => {
      return subtotal + ecommerce.subTotal;
    }, 0);
    return total + shopSubTotal;
  }, 0);

  return (
    <Modal title="Thông tin voucher" footer={null} {...props}>
      <Radio.Group
        value={
          groupByShop
            .find((shop) => shop.shopId === props.shopId)
            ?.groupByEcommerce.find(
              (ecommerce) => ecommerce.ecommerceId === props.ecommerceId,
            )?.voucherIds?.[0]
        }
        onChange={(e) => {
          const selectedVoucherId = e.target.value;

          const selectedVoucher = shopVouchers?.find(
            (voucher) => voucher.id_voucher === selectedVoucherId,
          );

          const internalDiscountTotal =
            selectedVoucher?.type === "percentage"
              ? Math.min(
                  selectedVoucher?.max_discount ?? 0,
                  selectedVoucher?.discount * subTotal,
                )
              : (selectedVoucher?.discount ?? 0);

          const newGroupByShop = groupByShop.map((shop) => {
            if (shop.shopId === props.shopId) {
              return {
                ...shop,
                groupByEcommerce: shop.groupByEcommerce.map((ecommerce) => {
                  if (ecommerce.ecommerceId === props.ecommerceId) {
                    return {
                      ...ecommerce,
                      voucherIds: [selectedVoucherId],
                      internalDiscountTotal: internalDiscountTotal,
                    };
                  }
                  return ecommerce;
                }),
              };
            }
            return shop;
          });

          dispatch(setGroupByShop(newGroupByShop));
        }}
      >
        <List
          grid={{ gutter: [2, 2], column: 1 }}
          dataSource={shopVouchers}
          renderItem={(item) => (
            <List.Item key={item?.id_voucher} style={{ padding: 0 }}>
              <VoucherItem
                {...item}
                button={<Radio value={item?.id_voucher} />}
              />
            </List.Item>
          )}
        />
      </Radio.Group>
    </Modal>
  );
};

export default VoucherModal;
