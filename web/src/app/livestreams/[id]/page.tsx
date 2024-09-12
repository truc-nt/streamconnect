import MeetingProvider from "@/component/livestream/MeetingProvider";

const Page = ({ params }: { params: { id: string } }) => {
  return <MeetingProvider livestreamId={Number(params.id)} />;
};

export default Page;
