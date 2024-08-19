import { Row, Col } from "antd";
import Sidebar from "./component/Sidebar";

const Page = ({ params }: { params: { id: number } }) => {
  return (
    <>
      <Row gutter={{ xs: 8 }} className="h-full flex items-stretch">
        <Col span={16}></Col>
        <Col span={8} className="relative">
          <Sidebar livestreamId={params.id} />
        </Col>
      </Row>
    </>
  );
};

export default Page;
