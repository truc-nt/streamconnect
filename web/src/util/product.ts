const getProductOptionFromVariantOptions = (
  variantOptions: { [key: string]: string }[],
) => {
  const productOption: { [key: string]: string[] } = {};
  variantOptions?.forEach((variantOption) => {
    Object.keys(variantOption).forEach((key) => {
      if (!productOption[key]) {
        productOption[key] = [];
      }
      if (!productOption[key].includes(variantOption[key])) {
        productOption[key].push(variantOption[key]);
      }
    });
  });
  return productOption;
};

export const getAvailableProductOptions = (
  chosenOption: { [key: string]: string },
  variantOptions: { [key: string]: string }[],
) => {
  const availableVariantOptionsWithChosenOption = variantOptions.filter(
    (variantOption) =>
      Object.keys(chosenOption).every(
        (key) => chosenOption[key] === variantOption[key],
      ),
  );

  return getProductOptionFromVariantOptions(
    availableVariantOptionsWithChosenOption,
  );
};

export { getProductOptionFromVariantOptions };
