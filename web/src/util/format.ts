export const formatPrice = (price: number | string | undefined) => {
  if (price === undefined) return "";
  console.log(parseInt(price).toLocaleString("vi-VN"));
  return price.toString().replace(/\B(?<!\.\d*)(?=(\d{3})+(?!\d))/g, ".");
};
