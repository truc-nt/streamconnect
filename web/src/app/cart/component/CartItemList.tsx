import { List } from "antd";
import CartItemsGroupByShop from "./CartItemsGroupByShop";
import { ICart, ICartItem } from "@/api/cart";

const CartItemList = ({ cart }: { cart: ICart[] }) => {
  return (
    <List
      grid={{ gutter: 16, column: 1 }}
      dataSource={cart}
      renderItem={(item) => (
        <List.Item>
          <CartItemsGroupByShop {...item} />
        </List.Item>
      )}
    />
  );
};

export default CartItemList;
