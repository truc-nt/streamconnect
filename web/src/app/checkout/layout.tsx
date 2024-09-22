import dynamic from "next/dynamic";
const ProtectedRoute = dynamic(
  () => import("@/component/core/ProtectedRoute"),
  {
    ssr: false,
  },
);

const Layout = ({ children }: { children: React.ReactNode }) => {
  return <ProtectedRoute>{children}</ProtectedRoute>;
};

export default Layout;
