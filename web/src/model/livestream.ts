export interface IBaseLivestream {
  id_livestream: number;
  fk_shop: number;
  title: string;
  description: string;
  status: string;
  meeting_id: string;
  hls_url: string;
  start_time: string;
}
