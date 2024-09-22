import { Menu as AntdMenu } from "antd";
import type { MenuProps } from "antd";
import {
  ContactsOutlined,
  EnvironmentOutlined,
  CompassFilled,
  FireFilled,
  TeamOutlined,
  CarryOutOutlined,
  SignalFilled,
  VideoCameraFilled,
  ShopFilled,
  ProductFilled,
} from "@ant-design/icons";
import { usePathname } from "next/navigation";
import Link from "next/link";

type MenuItem = Required<MenuProps>["items"][number];

const mainNavigator: MenuItem[] = [
  {
    key: "/",
    label: <Link href="/">Khám phá</Link>,
    icon: <CompassFilled />,
  },
  {
    key: "/trending",
    label: <Link href="/trending">Xu hướng</Link>,
    icon: <FireFilled />,
  },
  {
    key: "/following",
    label: <Link href="/following">Đang theo dõi</Link>,
    icon: <TeamOutlined />,
  },
];

const userNavigator: MenuItem[] = [
  {
    key: "/user",
    label: <Link href="/user/profile">Thông tin tài khoản</Link>,
    icon: <ContactsOutlined />,
  },
  {
    key: "/user/orders",
    label: <Link href="/user/orders">Đơn hàng</Link>,
    icon: <CarryOutOutlined />,
  },
  {
    key: "/user/addresses",
    label: <Link href="/user/addresses">Địa chỉ</Link>,
    icon: <EnvironmentOutlined />,
  },
];

const sellerNavigator: MenuItem[] = [
  {
    key: "/seller",
    label: <Link href="/seller/dashboard">Kênh bán hàng</Link>,
    icon: <SignalFilled />,
  },
  {
    key: "/seller/orders",
    label: <Link href="/seller/orders">Đơn hàng</Link>,
    icon: <CarryOutOutlined />,
  },
  {
    key: "/seller/livestreams",
    label: <Link href="/seller/livestreams">Livestream</Link>,
    icon: <VideoCameraFilled />,
  },
  {
    key: "/seller/shops",
    label: <Link href="/seller/shops">Cửa hàng</Link>,
    icon: <ShopFilled />,
  },
  {
    key: "products",
    label: "Sản phẩm",
    icon: <ProductFilled />,
    children: [
      {
        key: "/seller/products/internal",
        label: <Link href="/seller/products/internal">Hệ thống</Link>,
      },
      {
        key: "/seller/products/external",
        label: <Link href="/seller/products/external">Liên kết</Link>,
      },
    ],
  },
];

const Menu = () => {
  const pathname = usePathname();
  let items = mainNavigator;
  if (pathname.startsWith("/user")) {
    items = userNavigator;
  } else if (pathname.startsWith("/seller")) {
    items = sellerNavigator;
  }

  return (
    <AntdMenu
      mode="inline"
      selectedKeys={[pathname]}
      items={items}
      style={{
        borderRight: 0,
      }}
    />
  );
};

export default Menu;
