import { IBaseProduct } from "@/model/product";
export interface IBaseLivestream {
  id_livestream: number;
  id_shop: number;
  title: string;
  description: string;
  status: string;
  meeting_id: string;
  hls_url: string;
  start_time: string;
  end_time: string;
}

export interface IBaseLivestreamProduct extends IBaseProduct {
  id_livestream_product: number;
  id_livestream: number;
  priority: number;
  is_livestreamed: boolean;
}
