export const bgColorByUserId = (userId: number): string => {
  if (userId === 251134) {
    return "bg-red-700";
  }
  if (userId === 300752) {
    return "bg-orange-700";
  }
  if (userId === 815162) {
    return "bg-purple-700";
  }
  return "";
};

export const bgColorByUserDisplayName = (displayName: string): string => {
  if (displayName === "Kyle Klendworth") {
    return "bg-red-700";
  }
  if (displayName === "Jacob Collins") {
    return "bg-orange-700";
  }
  if (displayName === "Jack Glenzinski") {
    return "bg-purple-700";
  }
  return "";
};
