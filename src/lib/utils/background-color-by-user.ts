export const bgColorByUserId = (userId: number): string => {
  if (userId === 300752) {
    return "bg-orange-600";
  }
  if (userId === 815162) {
    return "bg-purple-600";
  }
  return "";
};

export const bgColorByUserDisplayName = (displayName: string): string => {
  if (displayName === "Jacob Collins") {
    return "bg-orange-600";
  }
  if (displayName === "Jack Glenzinski") {
    return "bg-purple-600";
  }
  return "";
};
