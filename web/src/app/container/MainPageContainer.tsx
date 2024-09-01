// "use client";
// import CategorySlider from "@/components/CategorySlider";
// import LivestreamPreview from "@/components/LivestreamPreview";
// import LivestreamsList from "@/components/LivestreamsList";
// import { Typography } from "@mui/material";
// import Alert from "@/components/core/Alert";
// import { useAppSelector } from "@/store/store";
// import {useEffect, useState} from "react";
// import { fetchMeeting } from "@/api/livestream";
// import {sdkToken} from "@/api/axios";
// import {MeetingProvider, useMeeting} from "@videosdk.live/react-sdk";
//
// // function Container({children}) {
// //   const [joined, setJoined] = useState("");
// //   //Get the method which will be used to join the meeting.
// //   console.log(useMeeting());
// //   const { join } = useMeeting();
// //   const mMeeting = useMeeting({
// //     //callback for when a meeting is joined successfully
// //     onMeetingJoined: () => {
// //       setJoined("JOINED");
// //     },
// //     //callback for when there is an error in a meeting
// //     onError: (error) => {
// //       alert(error.message);
// //     },
// //   });
// //   const joinMeeting = () => {
// //     setJoined("JOINING");
// //     join();
// //   };
// //   // useEffect(() => {
// //   //   joinMeeting();
// //   // }, []);
// //   return (
// //       <div>
// //         {children}
// //       </div>
// //   );
// // }
//
// function MainPageContainer() {
//   // const [meetingIds, setMeetingIds] = useState<string[]>([]);
//   // useEffect(() => {
//   //   const fetchData = async () => {
//   //     const meetings = await fetchMeeting("upcoming", true);
//   //     setMeetingIds(meetings.map((meeting: any) => meeting.id));
//   //   };
//   //   fetchData().catch(err => console.log(err));
//   // }, []);
//   const meetingId = "xi23-ne78-bljs";
//   return (
//       <>
//         <CategorySlider />
//         <MeetingProvider
//             config={{
//               meetingId: meetingId,
//               micEnabled: true,
//               webcamEnabled: true,
//               name: "TestUser",
//               mode: "VIEWER",
//               multiStream: false
//             }}
//             token={sdkToken}
//             reinitialiseMeetingOnConfigChange={true}
//             joinWithoutUserInteraction={true}
//         >
//           <LivestreamPreview meetingId={meetingId}/>
//         </MeetingProvider>
//         <Typography
//             variant="h6"
//             sx={{ fontSize: "20px", fontWeight: "bold", color: "white" }}
//         >
//           Livestreams Được Đề Xuất
//         </Typography>
//         <LivestreamsList />
//       </>
//   );
// };
//
// export default MainPageContainer;
